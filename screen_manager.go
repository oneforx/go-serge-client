package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	goecs "github.com/oneforx/go-ecs"
)

type ScreenManager struct {
	screens       []*IScreen
	currentScreen goecs.Identifier
}

func (sm *ScreenManager) Update() error {
	for _, screen := range sm.screens {
		screenLocalised := *screen
		if screenLocalised.GetId().Equals(sm.currentScreen) {
			if err := screenLocalised.Update(); err != nil {
				return err
			}
		}
	}
	return nil
}

func (sm *ScreenManager) Draw(screen *ebiten.Image) {
	for _, s := range sm.screens {
		sLocalised := *s
		if sLocalised.GetId().Equals(sm.currentScreen) {
			sLocalised.Draw(screen)
		}
	}
}

func (sm *ScreenManager) AddScreen(screen IScreen) {
	sm.screens = append(sm.screens, &screen)
}

func (sm *ScreenManager) GetScreen(id goecs.Identifier) *IScreen {
	for _, screen := range sm.screens {
		screenLocalised := *screen
		if screenLocalised.GetId().Equals(id) {
			return screen
		}
	}
	return nil
}

func (sm *ScreenManager) SetScreen(id goecs.Identifier) {
	// If the screen is already active, do nothing
	if sm.currentScreen.Equals(id) {
		return
	}

	previousScreen := sm.GetScreen(sm.currentScreen)
	if previousScreen != nil {
		previousScreenLocalised := *previousScreen
		previousScreenLocalised.OnClose()
	}

	sm.currentScreen = id
	currentScreenLocalised := *sm.GetScreen(id)
	currentScreenLocalised.OnActive()
}
