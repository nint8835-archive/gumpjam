package sprites

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed rat.png
var ratBytes []byte
var Rat *ebiten.Image

//go:embed success.png
var successBytes []byte
var Success *ebiten.Image

//go:embed tileset.png
var tilesetBytes []byte
var Tileset *ebiten.Image

func GetTile(xOffset, yOffset int64) *ebiten.Image {
	return Tileset.SubImage(image.Rect(int(xOffset), int(yOffset), int(xOffset+32), int(yOffset+32))).(*ebiten.Image)
}

func init() {
	var err error
	Rat, _, err = ebitenutil.NewImageFromReader(bytes.NewReader(ratBytes))
	if err != nil {
		panic(fmt.Errorf("failed to load rat image: %w", err))
	}

	Success, _, err = ebitenutil.NewImageFromReader(bytes.NewReader(successBytes))
	if err != nil {
		panic(fmt.Errorf("failed to load success image: %w", err))
	}

	Tileset, _, err = ebitenutil.NewImageFromReader(bytes.NewReader(tilesetBytes))
	if err != nil {
		panic(fmt.Errorf("failed to load tileset image: %w", err))
	}
}
