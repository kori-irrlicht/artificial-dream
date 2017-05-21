package scenes

import "github.com/kori-irrlicht/artificial-dream/core"

var _ core.Scene = &Game{}

type Game struct {
}

func (m *Game) Input()  {}
func (m *Game) Update() {}
func (m *Game) Render() {}
func (m *Game) Name() string {
	return "Game"
}
