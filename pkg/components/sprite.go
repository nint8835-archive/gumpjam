package components

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type SpriteLayer int

const (
	SpriteLayerBackground SpriteLayer = iota
	SpriteLayerForeground
)

type Sprite struct {
	Image *ebiten.Image
	Layer SpriteLayer
}

func NewPlaceholderSprite(width, height int, layer SpriteLayer, color color.Color) Sprite {
	image := ebiten.NewImage(width, height)
	image.Fill(color)

	return Sprite{
		Image: image,
		Layer: layer,
	}
}
