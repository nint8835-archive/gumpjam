package systems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/pkg/engine"

	"github.com/nint8835/gumpjam/pkg/components"
)

type Render struct {
	*components.Sprite
	*components.Position
}

func (r *Render) Draw(_ engine.World, screen *ebiten.Image) {
	screenWidth, screenHeight := float64(screen.Bounds().Dx()), float64(screen.Bounds().Dy())
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(r.Position.X*screenWidth, r.Position.Y*screenHeight)

	screen.DrawImage(r.Sprite.Image, options)
}

var _ engine.SystemDrawer = (*Render)(nil)
