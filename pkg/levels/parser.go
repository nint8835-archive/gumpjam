package levels

import (
	_ "embed"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/sedyh/mizu/pkg/engine"
	"golang.org/x/image/colornames"

	"github.com/nint8835/gumpjam/pkg/components"
	"github.com/nint8835/gumpjam/pkg/entities"
	"github.com/nint8835/gumpjam/pkg/levels/ldtk_parser"
)

//go:embed Gumpjam.ldtk
var ldtkFile []byte

func worldToGrid(x, y int64) (int, int) {
	fmt.Println(x, y)
	return int(x) / 640, int(y) / 480
}

func Load(w engine.World) error {
	data, err := ldtk_parser.UnmarshalLdtkJSON(ldtkFile)
	if err != nil {
		return fmt.Errorf("failed to unmarshal ldtk file: %w", err)
	}

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
			default:
				return fmt.Errorf("unknown layer type: %s", layer.Type)
			}
		}
	}

	spew.Dump(data)

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
			fmt.Println(cellX, cellY)

			w.AddEntities(&entities.Placeholder{
				Position: components.NewGridPosition(int(x), int(y), cellX, cellY),
				Sprite:   components.NewPlaceholderSprite(int(layer.GridSize), int(layer.GridSize), components.SpriteLayerBackground, "WALL", colornames.Grey),
				Hitbox:   components.Hitbox{Width: float64(layer.GridSize), Height: float64(layer.GridSize)},
			})
		}
	}

	return nil
}

func loadEntityLayer(w engine.World, layer ldtk_parser.LayerInstance, level ldtk_parser.Level) error {
	for _, entity := range layer.EntityInstances {
		cellX, cellY := worldToGrid(level.WorldX, level.WorldY)

		switch entity.Identifier {
		case "Player":
			w.AddEntities(&entities.Player{
				Position: components.NewGridPosition(int(entity.Grid[0]), int(entity.Grid[1]), cellX, cellY),
				Sprite:   components.NewPlaceholderSprite(int(entity.Width), int(entity.Height), components.SpriteLayerForeground, "RAT", colornames.Magenta),
				Gravity:  components.NewGravity(),
				Hitbox:   components.Hitbox{Width: float64(entity.Width), Height: float64(entity.Height)},
			})
		default:
			return fmt.Errorf("unknown entity type: %s", entity.Identifier)
		}
	}

	return nil
}
