package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type game struct {
	isInput            bool
	isUpdate           bool
	isInputAfterUpdate bool
	isRender           bool
}

func (g *game) Input() {
	g.isInput = true
	g.isInputAfterUpdate = false
}

func (g *game) Update() {
	g.isUpdate = true
	g.isInputAfterUpdate = true
}

func (g *game) Render() {
	g.isRender = true
}

func TestGameLoop(t *testing.T) {
	g := &game{}
	Convey("Main is called", t, func() {
		gameLoop(g)
		Convey("Input is run", func() {
			So(g.isInput, ShouldBeTrue)
		})
		Convey("Update is run", func() {
			So(g.isUpdate, ShouldBeTrue)
			Convey("After Input", func() {
				So(g.isInputAfterUpdate, ShouldBeTrue)
			})

		})
		Convey("Render is run", func() {
			So(g.isRender, ShouldBeTrue)
		})

	})
}
