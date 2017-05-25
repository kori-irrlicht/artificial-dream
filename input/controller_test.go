package input

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/veandco/go-sdl2/sdl"
)

func TestGameController(t *testing.T) {

	Convey("A new GameController", t, func() {
		gc := NewGameController().(*gameController)
		So(gc, ShouldNotBeNil)
		gc.mapping[sdl.CONTROLLER_BUTTON_DPAD_UP] = InputUp
		Convey("DPAD UP is pressed", func() {
			gc.Update(&sdl.ControllerButtonEvent{
				Button: sdl.CONTROLLER_BUTTON_DPAD_UP,
				State:  sdl.PRESSED,
			})
			So(gc.keyState[InputUp], ShouldBeTrue)
			Convey("DPAD UP is released", func() {
				gc.Update(&sdl.ControllerButtonEvent{
					Button: sdl.CONTROLLER_BUTTON_DPAD_UP,
					State:  sdl.RELEASED,
				})
				So(gc.keyState[InputUp], ShouldBeFalse)
			})
		})
	})

}
