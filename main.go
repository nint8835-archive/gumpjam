package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/pkg/engine"

	"github.com/nint8835/gumpjam/pkg/scenes"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Gumpjam")

	instance := engine.NewGame(&scenes.Game{})

	if err := ebiten.RunGame(instance); err != nil {
		log.Fatal(err)
	}
}
