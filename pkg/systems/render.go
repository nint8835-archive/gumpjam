package systems

import (
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/pkg/engine"

	"github.com/nint8835/gumpjam/pkg/components"
)

type Render struct {
	*components.Camera
	*components.Position
}

type spriteData struct {
	*components.Position
	*components.Sprite
}

func (r *Render) Draw(w engine.World, screen *ebiten.Image) {
	view := w.View(components.Position{}, components.Sprite{})

	var visibleSprites []spriteData

	view.Each(func(entity engine.Entity) {
		var position *components.Position
		var sprite *components.Sprite
		entity.Get(&position, &sprite)

		if position.CellX != r.CellX || position.CellY != r.CellY {
			return
		}

		visibleSprites = append(visibleSprites, spriteData{
			Position: position,
			Sprite:   sprite,
		})
	})

	sort.Slice(visibleSprites, func(i, j int) bool {
		return visibleSprites[i].Sprite.Layer < visibleSprites[j].Sprite.Layer
	})

	for _, sprite := range visibleSprites {
		options := &ebiten.DrawImageOptions{}
		if sprite.Sprite.FlipX {
			options.GeoM.Scale(-1, 1)
			options.GeoM.Translate(float64(sprite.Image.Bounds().Dx()), 0)
		}
		options.GeoM.Translate(sprite.Position.X, sprite.Position.Y)
		screen.DrawImage(sprite.Sprite.Image.(*ebiten.Image), options)
	}
}

var _ engine.SystemDrawer = (*Render)(nil)
