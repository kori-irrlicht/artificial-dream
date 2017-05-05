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
	isRunning           bool
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

func (g *game) Running() bool {
	return g.isRunning
}

func TestGameLoop(t *testing.T) {
	g := &game{}
	Convey("Main is called", t, func() {
		Convey("Game is not running", func() {
			g.isRunning = false
			gameLoop(g)
			Convey("Input is not run", func() {
				So(g.isInput, ShouldBeFalse)
			})
			Convey("Update is not run", func() {
				So(g.isUpdate, ShouldBeFalse)

			})
			Convey("Render is not run", func() {
				So(g.isRender, ShouldBeFalse)
			})

		})

		Convey("Game is running", func() {
			g.isRunning = true
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

	})
}
