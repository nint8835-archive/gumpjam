package scenes

import (
	"github.com/sedyh/mizu/pkg/engine"
	"golang.org/x/image/colornames"

	"github.com/nint8835/gumpjam/pkg/components"
	"github.com/nint8835/gumpjam/pkg/entities"
	"github.com/nint8835/gumpjam/pkg/systems"
)

type Game struct{}

func (g *Game) Setup(w engine.World) {
	w.AddComponents(
		components.Sprite{},
		components.Position{},
		components.Camera{},
		components.Velocity{},
		components.Hitbox{},
	)

	w.AddSystems(
		&systems.Player{},
		&systems.ScreenEdgeTransition{},
		&systems.Velocity{},
		&systems.Render{},
		&systems.Debug{},
	)

	w.AddEntities(
		&entities.Player{
			Sprite:   components.NewPlaceholderSprite(32, 32, components.SpriteLayerForeground, "RAT", colornames.Red),
			Position: components.NewGridPosition(10, 10, 0, 0),
			Hitbox:   components.Hitbox{Width: 32, Height: 32},
		},
	)

	var walls []any
	for cellX := 0; cellX < 5; cellX++ {
		for cellY := 0; cellY < 5; cellY++ {
			for x := 0; x < 20; x++ {
				for y := 0; y < 15; y++ {
					// Create a border around the map
					if (x != 0 && x != 19) && (y != 0 && y != 14) {
						continue
					}
					// Create gaps in the center of each side if there's an adjacent cell
					if (x == 9 || x == 10) && ((y == 0 && cellY != 0) || (y == 14 && cellY != 4)) {
						continue
					}
					if (y == 6 || y == 7) && ((x == 0 && cellX != 0) || (x == 19 && cellX != 4)) {
						continue
					}

					walls = append(walls, &entities.Placeholder{
						Sprite:   components.NewPlaceholderSprite(32, 32, components.SpriteLayerForeground, "WALL", colornames.Gray),
						Position: components.NewGridPosition(x, y, cellX, cellY),
						Hitbox:   components.Hitbox{Width: 32, Height: 32},
					})
				}
			}
		}
	}

	w.AddEntities(walls...)
}

var _ engine.Scene = (*Game)(nil)
