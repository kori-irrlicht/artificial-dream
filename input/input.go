package core

// InputType is an abstract layer between controller/keyboard input and the input
// the game will receive
type InputType int

// KeyboardMapping maps a key to an InputType
type KeyboardMapping map[int]InputType

// The InputTypes, which can be used
const (
	InputUp InputType = iota
	InputDown
	InputLeft
	InputRight
)

// Controller is an abstraction of the keyboard or a game controller (PS4, XBox1, Steam,...)
type Controller interface {
	IsDown(InputType) bool
	Update()
}

// keyboardController is used, when the game is played with a keyboard
type keyboardController struct {
	keyState map[InputType]bool
	mapping  KeyboardMapping
}

// IsDown implements the Controller interface
func (k keyboardController) IsDown(it InputType) bool {
	return false
}

// Update implements the Controller interface
func (k keyboardController) Update() {

}

// NewKeyboardController creates a new Controller with the given KeyboardMapping
func NewKeyboardController(km KeyboardMapping) Controller {
	c := keyboardController{}
	c.keyState = make(map[InputType]bool)
	c.mapping = km
	return c
}
