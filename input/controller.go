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
			for _, v := range gc.mapping[int(t.Button)] {
				gc.keyState[v] = true
			}
		}
		if t.State == sdl.RELEASED {
			for _, v := range gc.mapping[int(t.Button)] {
				gc.keyState[v] = false
			}
		}
	case *sdl.ControllerAxisEvent:

	}

}

func NewGameController() Controller {
	gc := &gameController{}
	gc.keyState = make(map[InputType]bool)
	gc.ResetMapping()
	return gc
}

func (gc *gameController) ResetMapping() {
	m := Mapping{
		sdl.CONTROLLER_BUTTON_DPAD_UP:    []InputType{InputUp},
		sdl.CONTROLLER_BUTTON_DPAD_DOWN:  []InputType{InputDown},
		sdl.CONTROLLER_BUTTON_DPAD_LEFT:  []InputType{InputLeft},
		sdl.CONTROLLER_BUTTON_DPAD_RIGHT: []InputType{InputRight},

		sdl.CONTROLLER_BUTTON_A:             []InputType{InputOK, InputInteract},
		sdl.CONTROLLER_BUTTON_B:             []InputType{InputBack},
		sdl.CONTROLLER_BUTTON_RIGHTSHOULDER: []InputType{InputAttack},

		sdl.CONTROLLER_BUTTON_START: []InputType{InputPause},
	}

	gc.mapping = m
}
