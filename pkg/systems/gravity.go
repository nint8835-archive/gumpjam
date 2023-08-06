package systems

import (
	"math"

	"github.com/sedyh/mizu/pkg/engine"

	"github.com/nint8835/gumpjam/pkg/components"
)

type Gravity struct {
	*components.Velocity
	*components.Gravity
}

func (g *Gravity) Update(w engine.World) {
	g.Velocity.Y = math.Min(g.Velocity.Y+g.Gravity.Acceleration, g.TerminalVelocity)
}

var _ engine.SystemUpdater = (*Gravity)(nil)
