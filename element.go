package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	goecs "github.com/oneforx/go-ecs"
)

type IElement interface {
	// GetId returns the id of the element
	GetId() goecs.Identifier
	Update() error
	Draw(screen *ebiten.Image)
}
