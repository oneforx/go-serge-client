package elements

import (
	"bytes"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	goecs "github.com/oneforx/go-ecs"
)

type InputElement struct {
	Id        goecs.Identifier
	value     bytes.Buffer
	keyStates map[ebiten.Key]bool
	isFocus   bool
	Position  struct {
		X int
		Y int
	}
}

func (e *InputElement) GetId() goecs.Identifier {
	return e.Id
}

func (e *InputElement) Update() error {
	mx, my := ebiten.CursorPosition()
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if mx >= e.Position.X && mx <= e.Position.X+100 && my >= e.Position.Y && my <= e.Position.Y+20 {
			e.isFocus = true
		} else {
			e.isFocus = false
		}
	}

	if e.isFocus {
		keysToCheck := []ebiten.Key{
			ebiten.KeyBackspace,
			ebiten.KeyEnter,
			ebiten.KeyShift,
		}

		for i := ebiten.KeyA; i <= ebiten.KeyZ; i++ {
			keysToCheck = append(keysToCheck, i)
		}

		for _, k := range keysToCheck {
			isPressed := ebiten.IsKeyPressed(k)
			if isPressed && !e.keyStates[k] {
				e.keyStates[k] = true
				if k == ebiten.KeyBackspace {
					// Delete the last character
					if e.value.Len() > 0 {
						e.value.Truncate(e.value.Len() - 1)
					}
				} else if k == ebiten.KeyEnter {
					// Clear the input text
					e.value.Reset()
				} else if k != ebiten.KeyShift && k.String() != "Unknown" {
					// Add the character to the input text
					e.value.WriteRune(rune(k.String()[0]))
				}
			} else if !isPressed {
				e.keyStates[k] = false
			}
		}
	}
	return nil
}

func (e *InputElement) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, e.value.String())
	ebitenutil.DebugPrintAt(screen, "|", e.Position.X, e.Position.Y)
}

func InputElementFactory() *InputElement {
	return &InputElement{
		Id: goecs.Identifier{
			Namespace: "goserge",
			Path:      "input_element",
		},
		keyStates: make(map[ebiten.Key]bool),
	}
}
