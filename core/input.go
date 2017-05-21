package core

type InputType int

const (
	INPUT_UP InputType = iota
	INPUT_DOWN
	INPUT_LEFT
	INPUT_RIGHT
)

type Controller interface {
	IsDown(InputType) bool
}
