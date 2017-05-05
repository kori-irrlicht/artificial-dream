package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type game struct {
	isInput             int
	isUpdate            int
	isUpdateAfterInput  bool
	isRender            int
	isRenderAfterUpdate bool
	isRunning           chan bool
}

func (g *game) Input() {
	g.isInput++
	g.isUpdateAfterInput = false
}

func (g *game) Update() {
	g.isUpdate++
	g.isUpdateAfterInput = true
	g.isRenderAfterUpdate = false
}

func (g *game) Render() {
	g.isRender++
	g.isRenderAfterUpdate = true
}

func (g *game) Running() bool {
	return <-g.isRunning
}

func TestGameLoop(t *testing.T) {

	Convey("Main is called", t, func() {
		g := &game{}
		g.isRunning = make(chan bool, 4)
		Convey("Game is not running", func() {
			g.isRunning <- false
			gameLoop(g)
			Convey("Input is not run", func() {
				So(g.isInput, ShouldEqual, 0)
			})
			Convey("Update is not run", func() {
				So(g.isUpdate, ShouldEqual, 0)

			})
			Convey("Render is not run", func() {
				So(g.isRender, ShouldEqual, 0)
			})

		})

		Convey("Game is running", func() {
			g.isRunning <- true
			g.isRunning <- false
			gameLoop(g)
			Convey("Input is run", func() {
				So(g.isInput, ShouldEqual, 1)
			})
			Convey("Update is run", func() {
				So(g.isUpdate, ShouldEqual, 1)
				Convey("After Input", func() {
					So(g.isUpdateAfterInput, ShouldBeTrue)
				})

			})
			Convey("Render is run", func() {
				So(g.isRender, ShouldEqual, 1)
				Convey("After Update", func() {
					So(g.isRenderAfterUpdate, ShouldBeTrue)
				})
			})
		})

		Convey("Game is running 3 times", func() {
			g.isRunning <- true
			g.isRunning <- true
			g.isRunning <- true
			g.isRunning <- false
			gameLoop(g)
			Convey("Input is run 3 times", func() {
				So(g.isInput, ShouldEqual, 3)
			})
			Convey("Update is run 3 times", func() {
				So(g.isUpdate, ShouldEqual, 3)
			})
			Convey("Render is run 3 times", func() {
				So(g.isRender, ShouldEqual, 3)
			})
		})

	})
}
