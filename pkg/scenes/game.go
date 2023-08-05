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
	)

	w.AddSystems(
		&systems.Render{},
		&systems.Debug{},
		&systems.Player{},
		&systems.ScreenEdgeTransition{},
		&systems.Velocity{},
	)

	w.AddEntities(
		&entities.Player{
			Sprite:   components.NewPlaceholderSprite(32, 32, components.SpriteLayerForeground, "RAT", colornames.Red),
			Position: components.NewGridPosition(10, 10, 0, 0),
		},
	)
}

var _ engine.Scene = (*Game)(nil)
