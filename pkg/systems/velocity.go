package systems

import (
	"github.com/sedyh/mizu/pkg/engine"

	"github.com/nint8835/gumpjam/pkg/components"
)

type Velocity struct {
	*components.Position
	*components.Velocity
}

func (v *Velocity) Update(w engine.World) {
	v.Position.X += v.Velocity.X
	v.Position.Y += v.Velocity.Y
}

var _ engine.SystemUpdater = (*Velocity)(nil)
