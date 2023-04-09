package screens

import (
	"github.com/hajimehoshi/ebiten/v2"
	goecs "github.com/oneforx/go-ecs"
	"github.com/oneforx/go-serge-client/elements"
)

type MainScreen struct {
	Id           goecs.Identifier
	inputElement *elements.InputElement
}

func (s *MainScreen) GetId() goecs.Identifier {
	return s.Id
}

func (s *MainScreen) OnActive() {
	s.inputElement = elements.InputElementFactory()
}

func (s *MainScreen) Update() error {
	s.inputElement.Update()
	return nil
}

func (s *MainScreen) Draw(screen *ebiten.Image) {
	s.inputElement.Draw(screen)
}

func (s *MainScreen) OnClose() {

}
