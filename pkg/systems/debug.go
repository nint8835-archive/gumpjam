package systems

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/sedyh/mizu/pkg/engine"

	"github.com/nint8835/gumpjam/pkg/components"
)

type Debug struct {
	*components.Camera
	*components.Position
	*components.Velocity
}

func (d *Debug) Update(w engine.World) {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		d.CellY--
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		d.CellY++
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		d.CellX--
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		d.CellX++
	}
}

func (d *Debug) Draw(w engine.World, screen *ebiten.Image) {
	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf(
			"CellX: %d, CellY: %d - %f FPS\nVX: %f, VY: %f, on ground: %t",
			d.CellX,
			d.CellY,
			ebiten.ActualFPS(),
			d.Velocity.X,
			d.Velocity.Y,
			d.Velocity.OnGround,
		),
	)
}

var _ engine.SystemUpdater = (*Debug)(nil)
var _ engine.SystemDrawer = (*Debug)(nil)
