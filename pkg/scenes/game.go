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
	)

	w.AddSystems(
		&systems.Render{},
	)

	var placeholders []any

	for x := 0; x < 20; x++ {
		for y := 0; y < 15; y++ {
			placeholderColor := placeholderColors[(x+y)%len(placeholderColors)]

			placeholders = append(placeholders, &entities.Placeholder{
				Sprite: *components.NewPlaceholderSprite(32, 32, placeholderColor),
				Position: components.Position{
					X: float64(x) / 20,
					Y: float64(y) / 15,
				},
			})
		}
	}

	w.AddEntities(
		placeholders...,
	)
}

var _ engine.Scene = (*Game)(nil)
