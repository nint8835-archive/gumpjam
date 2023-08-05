package components

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	Image *ebiten.Image
}

func NewPlaceholderSprite(width, height int, color color.Color) *Sprite {
	image := ebiten.NewImage(width, height)

	image.Fill(color)

	return &Sprite{
		Image: image,
	}
}
