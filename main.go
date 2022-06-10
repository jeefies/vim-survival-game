package main

import (
	"github.com/jeefies/vim-survival/scenes"
	"github.com/jeefies/vim-survival/log"
	vinput "github.com/jeefies/vim-survival/input"

	"github.com/hajimehoshi/ebiten/v2"
)

var logger = log.Logger

var (
	screenWidth, screenHeight = ebiten.ScreenSizeInFullscreen()
)

type Game struct {
	sceneManager * SceneManager
	input vinput.Input
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Update() error {
	// This is a copied way to manage  scenes, from ebiten/examples/blocks

	if g.sceneManager == nil {
		g.sceneManager = &SceneManager{}
		g.sceneManager.Register(&scenes.TitleScene{})
	}

	g.input.Update()
	// This is the old one, maybe it's made for handle error.
	if err := g.sceneManager.Update(&g.input); err != nil {
		// Handle Error Here!
		return err
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.sceneManager.Draw(screen)
}

func main() {
	// Why * 2 ? for pixel window?
	ebiten.SetWindowTitle("Vim survival")
	// This is a game, so we need it to be full screen ^_^
	ebiten.SetFullscreen(true)

	if err := ebiten.RunGame(&Game{}); err != nil {
		logger.Fatal(err)
	}
}
