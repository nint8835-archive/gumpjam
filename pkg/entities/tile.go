package entities

import "github.com/nint8835/gumpjam/pkg/components"

type Tile struct {
	components.Sprite
	components.Position
	components.Hitbox
}
