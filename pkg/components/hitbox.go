package components

type Hitbox struct {
	Width  float64
	Height float64

	AllowJumpThrough bool
	AllowFallThrough bool
}

func HasOverlap(aPos *Position, aHitbox *Hitbox, bPos *Position, bHitbox *Hitbox) bool {
	// TODO: handle cross-cell overlap?
	return aPos.X < bPos.X+bHitbox.Width &&
		aPos.X+aHitbox.Width > bPos.X &&
		aPos.Y < bPos.Y+bHitbox.Height &&
		aPos.Y+aHitbox.Height > bPos.Y
}
