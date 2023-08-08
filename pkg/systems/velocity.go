package systems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/pkg/engine"

	"github.com/nint8835/gumpjam/pkg/components"
)

type Velocity struct {
	*components.Position
	*components.Velocity
	*components.Hitbox
}

func (v *Velocity) Update(w engine.World) {
	futureX, futureY := v.Position.X+v.Velocity.X, v.Position.Y+v.Velocity.Y

	v.Velocity.OnGround = false

	w.View(&components.Position{}, &components.Hitbox{}).Each(func(e engine.Entity) {
		var otherEntPos *components.Position
		var otherEntHitbox *components.Hitbox
		e.Get(&otherEntPos, &otherEntHitbox)

		if otherEntPos == v.Position {
			return
		}

		if otherEntPos.CellX != v.Position.CellX || otherEntPos.CellY != v.Position.CellY {
			return
		}

		if v.Velocity.X != 0 && components.HasOverlap(
			&components.Position{X: futureX, Y: v.Position.Y},
			v.Hitbox,
			otherEntPos,
			otherEntHitbox,
		) {
			if v.Velocity.X > 0 {
				futureX = otherEntPos.X - v.Hitbox.Width
			} else {
				futureX = otherEntPos.X + otherEntHitbox.Width
			}
		}

		if v.Velocity.Y != 0 && components.HasOverlap(
			&components.Position{X: v.Position.X, Y: futureY},
			v.Hitbox,
			otherEntPos,
			otherEntHitbox,
		) {
			if v.Velocity.Y > 0 && !(otherEntHitbox.AllowFallThrough && ebiten.IsKeyPressed(ebiten.KeyS)) {
				futureY = otherEntPos.Y - v.Hitbox.Height
				v.Velocity.OnGround = true
			} else if v.Velocity.Y < 0 && !(otherEntHitbox.AllowJumpThrough && ebiten.IsKeyPressed(ebiten.KeyW)) {
				futureY = otherEntPos.Y + otherEntHitbox.Height
			}
		}
	})

	v.Velocity.X = futureX - v.Position.X
	v.Position.X = futureX
	v.Velocity.Y = futureY - v.Position.Y
	v.Position.Y = futureY
}

var _ engine.SystemUpdater = (*Velocity)(nil)
