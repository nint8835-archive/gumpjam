package components

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	Image *ebiten.Image
	Layer int
}

func NewPlaceholderSprite(width, height, layer int, color color.Color) Sprite {
	image := ebiten.NewImage(width, height)
	image.Fill(color)

	return Sprite{
		Image: image,
		Layer: layer,
	}
}
