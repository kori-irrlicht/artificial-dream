package input

import "github.com/veandco/go-sdl2/sdl"

type gameController struct {
}

func (gc *gameController) IsDown(it InputType) bool {
	return false
}

func (gc *gameController) Update(ev sdl.Event) {

}

func NewGameController() Controller {
	gc := &gameController{}
	return gc
}
