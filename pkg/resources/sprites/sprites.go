package sprites

import (
	"bytes"
	_ "embed"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed rat.png
var ratBytes []byte
var Rat *ebiten.Image

func init() {
	var err error
	Rat, _, err = ebitenutil.NewImageFromReader(bytes.NewReader(ratBytes))
	if err != nil {
		panic(fmt.Errorf("failed to load rat image: %w", err))
	}
}
