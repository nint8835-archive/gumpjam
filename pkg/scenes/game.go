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
		&systems.Player{},
		&systems.ScreenEdgeTransition{},
	)

	w.AddEntities(
		&entities.Player{
			Sprite:   components.NewPlaceholderSprite(32, 32, components.SpriteLayerForeground, "RAT", colornames.Red),
			Position: components.NewGridPosition(10, 10, 0, 0),
			Camera:   components.Camera{},
		},
	)
}

var _ engine.Scene = (*Game)(nil)
