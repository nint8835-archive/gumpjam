package systems

import (
	"github.com/sedyh/mizu/pkg/engine"

	"github.com/nint8835/gumpjam/pkg/components"
)

type ScreenEdgeTransition struct {
	*components.Position
}

func (s *ScreenEdgeTransition) Update(w engine.World) {
	screenWidth, screenHeight := float64(w.Bounds().Dx()), float64(w.Bounds().Dy())

	if s.X < 0 {
		s.CellX--
		s.X = screenWidth + s.X
	}
	if s.X > screenWidth {
		s.CellX++
		s.X = s.X - screenWidth
	}
	if s.Y < 0 {
		s.CellY--
		s.Y = screenHeight + s.Y
	}
	if s.Y > screenHeight {
		s.CellY++
		s.Y = s.Y - screenHeight
	}
}

var _ engine.SystemUpdater = (*ScreenEdgeTransition)(nil)
