package main

import "github.com/hajimehoshi/ebiten/v2"

type GoSergeGame struct {
	screenManager *ScreenManager
}

func (g *GoSergeGame) GetScreenManager() *ScreenManager {
	return g.screenManager
}

func (g *GoSergeGame) Update() error {
	if err := g.screenManager.Update(); err != nil {
		return err
	}
	return nil
}

func (g *GoSergeGame) Draw(screen *ebiten.Image) {
	g.screenManager.Draw(screen)
}

func (g *GoSergeGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}
