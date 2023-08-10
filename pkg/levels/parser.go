package levels

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"strconv"

	"github.com/davecgh/go-spew/spew"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/sedyh/mizu/pkg/engine"
	"golang.org/x/image/colornames"

	"github.com/nint8835/gumpjam/pkg/components"
	"github.com/nint8835/gumpjam/pkg/entities"
	"github.com/nint8835/gumpjam/pkg/levels/ldtk_parser"
)

//go:embed Gumpjam.ldtk
var ldtkFile []byte

//go:embed sprites/Tileset.png
var tilesetBytes []byte
var tileset *ebiten.Image

func worldToGrid(x, y int64) (int, int) {
	return int(x) / 640, int(y) / 480
}

func getFieldInstance(entity ldtk_parser.EntityInstance, fieldName string) ldtk_parser.FieldInstance {
	for _, field := range entity.FieldInstances {
		if field.Identifier == fieldName {
			return field
		}
	}

	return ldtk_parser.FieldInstance{}
}

func parseColourValue(hexString string) color.Color {
	colorInt, _ := strconv.ParseInt(hexString[1:], 16, 64)
	return color.RGBA{
		R: uint8(colorInt >> 16),
		G: uint8(colorInt >> 8),
		B: uint8(colorInt),
		A: 255,
	}
}

type tileData struct {
	AllowFallThrough bool
	AllowJumpThrough bool
}

func getTileData(file ldtk_parser.LdtkJSON, tile ldtk_parser.TileInstance) tileData {
	for _, customData := range file.Defs.Tilesets[0].CustomData {
		if customData.TileID != tile.T {
			continue
		}

		var data tileData
		err := json.Unmarshal([]byte(customData.Data), &data)
		if err != nil {
			panic(fmt.Errorf("failed to unmarshal tile data: %w", err))
		}

		return data
	}

	return tileData{}
}

func Load(w engine.World) error {
	data, err := ldtk_parser.UnmarshalLdtkJSON(ldtkFile)
	if err != nil {
		return fmt.Errorf("failed to unmarshal ldtk file: %w", err)
	}

	spew.Dump(data)

	for _, level := range data.Levels {
		for _, layer := range level.LayerInstances {
			switch ldtk_parser.Type(layer.Type) {
			case ldtk_parser.Entities:
				if err := loadEntityLayer(w, layer, level); err != nil {
					return fmt.Errorf("failed to load entity layer: %w", err)
				}
			case ldtk_parser.IntGrid:
				if err := loadIntGridLayer(w, layer, level); err != nil {
					return fmt.Errorf("failed to load int grid layer: %w", err)
				}
			case ldtk_parser.Tiles:
				if err := loadTileLayer(w, data, layer, level); err != nil {
					return fmt.Errorf("failed to load tile layer: %w", err)
				}
			default:
				return fmt.Errorf("unknown layer type: %s", layer.Type)
			}
		}
	}

	return nil
}

func loadIntGridLayer(w engine.World, layer ldtk_parser.LayerInstance, level ldtk_parser.Level) error {
	for x := int64(0); x < layer.CWid; x++ {
		for y := int64(0); y < layer.CHei; y++ {
			tile := layer.IntGridCSV[x+y*layer.CWid]
			if tile == 0 {
				continue
			}

			cellX, cellY := worldToGrid(level.WorldX, level.WorldY)

			w.AddEntities(&entities.Placeholder{
				Position: components.NewGridPosition(int(x), int(y), cellX, cellY),
				Sprite:   components.NewPlaceholderSprite(int(layer.GridSize), int(layer.GridSize), components.SpriteLayerBackground, "WALL", colornames.Grey),
				Hitbox:   components.Hitbox{Width: float64(layer.GridSize), Height: float64(layer.GridSize)},
			})
		}
	}

	return nil
}

func loadTileLayer(w engine.World, file ldtk_parser.LdtkJSON, layer ldtk_parser.LayerInstance, level ldtk_parser.Level) error {
	for _, tile := range layer.GridTiles {
		cellX, cellY := worldToGrid(level.WorldX, level.WorldY)

		tileData := getTileData(file, tile)

		w.AddEntities(&entities.Tile{
			Position: components.NewGridPosition(int(tile.Px[0]/layer.GridSize), int(tile.Px[1]/layer.GridSize), cellX, cellY),
			Sprite: components.Sprite{
				Image: tileset.SubImage(image.Rect(int(tile.Src[0]), int(tile.Src[1]), int(tile.Src[0])+int(layer.GridSize), int(tile.Src[1])+int(layer.GridSize))),
				Layer: components.SpriteLayerBackground,
			},
			Hitbox: components.Hitbox{
				Width:            float64(layer.GridSize),
				Height:           float64(layer.GridSize),
				AllowFallThrough: tileData.AllowFallThrough,
				AllowJumpThrough: tileData.AllowJumpThrough,
			},
		})
	}

	return nil
}

func loadEntityLayer(w engine.World, layer ldtk_parser.LayerInstance, level ldtk_parser.Level) error {
	for _, entity := range layer.EntityInstances {
		cellX, cellY := worldToGrid(level.WorldX, level.WorldY)
		position := components.NewGridPosition(int(entity.Grid[0]), int(entity.Grid[1]), cellX, cellY)

		switch entity.Identifier {
		case "Player":
			w.AddEntities(&entities.Player{
				Position: position,
				Sprite:   components.NewPlaceholderSprite(int(entity.Width), int(entity.Height), components.SpriteLayerForeground, "RAT", colornames.Magenta),
				Gravity:  components.NewGravity(),
				Hitbox:   components.Hitbox{Width: float64(entity.Width), Height: float64(entity.Height)},
			})
		case "Placeholder":
			w.AddEntities(&entities.Placeholder{
				Position: position,
				Sprite:   components.NewPlaceholderSprite(int(entity.Width), int(entity.Height), components.SpriteLayerForeground, "TEMP", parseColourValue(getFieldInstance(entity, "Colour").Value.(string))),
				Hitbox: components.Hitbox{
					Width:            float64(entity.Width),
					Height:           float64(entity.Height),
					AllowJumpThrough: getFieldInstance(entity, "AllowJumpThrough").Value.(bool),
					AllowFallThrough: getFieldInstance(entity, "AllowFallThrough").Value.(bool),
				},
			})
		default:
			return fmt.Errorf("unknown entity type: %s", entity.Identifier)
		}
	}

	return nil
}

func init() {
	var err error
	tileset, _, err = ebitenutil.NewImageFromReader(bytes.NewReader(tilesetBytes))
	if err != nil {
		panic(fmt.Errorf("failed to load tileset image: %w", err))
	}
}
