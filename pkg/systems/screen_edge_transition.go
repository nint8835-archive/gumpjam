package systems

import (
	"github.com/sedyh/mizu/pkg/engine"

	"github.com/nint8835/gumpjam/pkg/components"
)

type ScreenEdgeTransition struct {
	*components.Position
}

func (s *ScreenEdgeTransition) Update(w engine.World) {
	if s.X < 0 {
		s.CellX--
		s.X = 1 + s.X
	}
	if s.X > 1 {
		s.CellX++
		s.X = s.X - 1
	}
	if s.Y < 0 {
		s.CellY--
		s.Y = 1 + s.Y
	}
	if s.Y > 1 {
		s.CellY++
		s.Y = s.Y - 1
	}
}

var _ engine.SystemUpdater = (*ScreenEdgeTransition)(nil)
