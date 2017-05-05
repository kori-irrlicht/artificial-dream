package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type game struct {
	isInput             bool
	isUpdate            bool
	isUpdateAfterInput  bool
	isRender            bool
	isRenderAfterUpdate bool
}

func (g *game) Input() {
	g.isInput = true
	g.isUpdateAfterInput = false
}

func (g *game) Update() {
	g.isUpdate = true
	g.isUpdateAfterInput = true
	g.isRenderAfterUpdate = false
}

func (g *game) Render() {
	g.isRender = true
	g.isRenderAfterUpdate = true
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
				So(g.isUpdateAfterInput, ShouldBeTrue)
			})

		})
		Convey("Render is run", func() {
			So(g.isRender, ShouldBeTrue)
			Convey("After Update", func() {
				So(g.isRenderAfterUpdate, ShouldBeTrue)
			})
		})

	})
}
