package input

import "github.com/veandco/go-sdl2/sdl"

// keyboardController is used, when the game is played with a keyboard
type keyboardController struct {
	keyState map[InputType]bool
	mapping  Mapping
}

// IsDown implements the Controller interface
func (k keyboardController) IsDown(it InputType) bool {
	return false
}

// Update implements the Controller interface
func (k *keyboardController) Update(event sdl.Event) {

}

// NewKeyboardController creates a new Controller with the given KeyboardMapping
func NewKeyboardController() Controller {
	c := &keyboardController{}
	c.keyState = make(map[InputType]bool)
	return c
}

func (k *keyboardController) ResetMapping() {

}
