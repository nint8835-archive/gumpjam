package systems

import (
	"github.com/sedyh/mizu/pkg/engine"

	"github.com/nint8835/gumpjam/pkg/components"
)

type ScreenEdgeTransition struct {
	*components.Position
	*components.Hitbox
}

func (s *ScreenEdgeTransition) Update(w engine.World) {
	screenWidth, screenHeight := float64(w.Bounds().Dx()), float64(w.Bounds().Dy())

	if s.X+s.Hitbox.Width < 0 {
		s.CellX--
		s.X = screenWidth + s.X + s.Hitbox.Width - 0.01
	}
	if s.X > screenWidth {
		s.CellX++
		s.X = s.X - screenWidth - s.Hitbox.Width + 0.01
	}

	if s.Y+s.Hitbox.Height < 0 {
		s.CellY--
		s.Y = screenHeight + s.Y + s.Hitbox.Height - 0.01
	}
	if s.Y > screenHeight {
		s.CellY++
		s.Y = s.Y - screenHeight - s.Hitbox.Height + 0.01
	}
}

var _ engine.SystemUpdater = (*ScreenEdgeTransition)(nil)
