package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	goecs "github.com/oneforx/go-ecs"
	"github.com/oneforx/go-serge-client/screens"
)

var goSergeGame *GoSergeGame = &GoSergeGame{
	screenManager: &ScreenManager{},
}

func main() {
	goSergeGame.GetScreenManager().AddScreen(&screens.MainScreen{Id: goecs.Identifier{Namespace: "goserge", Path: "main_screen"}})
	goSergeGame.GetScreenManager().SetScreen(goecs.Identifier{Namespace: "goserge", Path: "main_screen"})
	if err := ebiten.RunGame(goSergeGame); err != nil {
		panic(err)
	}
}
