package main

import "time"

func main() {

}

func gameLoop(g Game) {
	dt := g.FrameTime()

	currentTime := g.Now()
	acc := 0 * time.Millisecond
	for g.Running() {
		newTime := g.Now()
		diff := newTime.Sub(currentTime)
		currentTime = newTime

		acc += diff
		g.Input()
		for acc >= dt {
			g.Update()
			acc -= dt
		}

		g.Render()
	}

}

type Game interface {
	Update()
	Input()
	Render()
	Running() bool
	Now() time.Time
	FrameTime() time.Duration
}
