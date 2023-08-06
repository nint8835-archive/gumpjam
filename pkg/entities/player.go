package entities

import "github.com/nint8835/gumpjam/pkg/components"

type Player struct {
	components.Sprite
	components.Position
	components.Camera
	components.Velocity
	components.Hitbox
}
