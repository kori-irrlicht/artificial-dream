package main

import (
	"time"

	"github.com/kori-irrlicht/artificial-dream/core"
	"github.com/kori-irrlicht/artificial-dream/input"
	"github.com/veandco/go-sdl2/sdl"
)

type game struct {
	running      bool
	sceneManager core.SceneManager
	controller   input.Controller
}

func (g *game) FrameTime() time.Duration {
	return 13 * time.Millisecond
}

func (g *game) Input() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			g.running = false
		default:

			g.controller.Update(event)

		}

	}
}

func (g *game) Now() time.Time {
	return time.Now()
}

func (g *game) Render() {}

func (g *game) Running() bool {
	return g.running
}

func (g *game) Update() {}
