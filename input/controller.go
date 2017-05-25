package input

import (
	"github.com/veandco/go-sdl2/sdl"
)

type gameController struct {
	keyState map[InputType]bool
	mapping  Mapping
}

func (gc *gameController) IsDown(it InputType) bool {
	return gc.keyState[it]
}

func (gc *gameController) Update(ev sdl.Event) {

	switch t := ev.(type) {
	case *sdl.ControllerButtonEvent:
		if t.State == sdl.PRESSED {
			gc.keyState[gc.mapping[int(t.Button)]] = true
		}
		if t.State == sdl.RELEASED {
			gc.keyState[gc.mapping[int(t.Button)]] = false
		}
	}
}

func NewGameController() Controller {
	gc := &gameController{}
	gc.keyState = make(map[InputType]bool)
	gc.ResetMapping()
	return gc
}

func (gc *gameController) ResetMapping() {
	m := make(Mapping)

	m[sdl.CONTROLLER_BUTTON_DPAD_UP] = InputUp
	m[sdl.CONTROLLER_BUTTON_DPAD_DOWN] = InputDown
	m[sdl.CONTROLLER_BUTTON_DPAD_LEFT] = InputLeft
	m[sdl.CONTROLLER_BUTTON_DPAD_RIGHT] = InputRight

	gc.mapping = m
}
