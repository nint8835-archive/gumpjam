package systems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/sedyh/mizu/pkg/engine"

	"github.com/nint8835/gumpjam/pkg/components"
)

type Player struct {
	*components.Camera
	*components.Velocity
}

func (p *Player) Update(w engine.World) {
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.X = -5
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.X = 5
	} else {
		p.X = 0
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) && p.Velocity.OnGround {
		p.Velocity.Y = -13
	}
}

var _ engine.SystemUpdater = (*Player)(nil)
