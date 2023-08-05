package scenes

import (
	"image/color"

	"github.com/sedyh/mizu/pkg/engine"
	"golang.org/x/image/colornames"

	"github.com/nint8835/gumpjam/pkg/components"
	"github.com/nint8835/gumpjam/pkg/entities"
	"github.com/nint8835/gumpjam/pkg/systems"
)

var placeholderColors = []color.Color{
	colornames.Lavenderblush,
	colornames.Lightseagreen,
	colornames.Palevioletred,
	colornames.Powderblue,
	colornames.Chocolate,
}

type Game struct{}

func (g *Game) Setup(w engine.World) {
	w.AddComponents(
		components.Sprite{},
		components.Position{},
		components.Camera{},
	)

	w.AddSystems(
		&systems.Render{},
		&systems.Debug{},
	)

	var placeholders []any
	for cellX := 0; cellX < 10; cellX++ {
		for cellY := 0; cellY < 10; cellY++ {
			for x := 0; x < components.GridWidth; x++ {
				for y := 0; y < components.GridHeight; y++ {
					if (x+y)%(cellX+cellY+1) != 0 {
						continue
					}

					placeholders = append(placeholders, &entities.Placeholder{
						Sprite:   components.NewPlaceholderSprite(32, 32, 0, placeholderColors[(x+y)%len(placeholderColors)]),
						Position: components.NewGridPosition(x, y, cellX, cellY),
					})
				}
			}
		}
	}

	w.AddEntities(
		&entities.Player{
			Sprite:   components.NewPlaceholderSprite(32, 32, 1, colornames.Purple),
			Position: components.NewGridPosition(10, 10, 0, 0),
			Camera:   components.Camera{},
		},
	)
	w.AddEntities(placeholders...)
}

var _ engine.Scene = (*Game)(nil)
