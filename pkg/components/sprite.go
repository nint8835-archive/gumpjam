package components

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"

	"github.com/nint8835/gumpjam/pkg/resources/fonts"
	"github.com/nint8835/gumpjam/pkg/utils"
)

type SpriteLayer int

const (
	SpriteLayerBackground SpriteLayer = iota
	SpriteLayerForeground
)

type Sprite struct {
	Image image.Image
	Layer SpriteLayer

	FlipX bool
}

func NewPlaceholderSprite(width, height int, layer SpriteLayer, label string, colour color.Color) Sprite {
	image := ebiten.NewImage(width, height)
	image.Fill(colour)

	boundRect := text.BoundString(fonts.MingLiU[fonts.SizeSmall], label)

	text.Draw(
		image,
		label,
		fonts.MingLiU[12],
		(width/2)-(boundRect.Dx()/2),
		(height/2)+((boundRect.Dy()-boundRect.Max.Y)/2),
		utils.GetContrastingTextColour(colour),
	)

	return Sprite{
		Image: image,
		Layer: layer,
	}
}
