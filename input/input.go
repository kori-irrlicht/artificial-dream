package input

import "github.com/veandco/go-sdl2/sdl"

// InputType is an abstract layer between controller/keyboard input and the input
// the game will receive
type InputType int

// The InputTypes, which can be used
const (
	InputUp InputType = iota
	InputDown
	InputLeft
	InputRight

	InputOK
	InputBack

	InputPause

	InputInteract
	InputAttack
	InputUseItem
)

// Controller is an abstraction of the keyboard or a game controller (PS4, XBox1, Steam,...)
type Controller interface {
	IsDown(InputType) bool
	Update(sdl.Event)
	ResetMapping()
}

// Mapping maps a key to an InputType
type Mapping map[int][]InputType
