package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type game struct {
	isInput  bool
	isUpdate bool
}

func (g *game) Input() {
	g.isInput = true
}

func (g *game) Update() {
	g.isUpdate = true
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

		})

	})
}
