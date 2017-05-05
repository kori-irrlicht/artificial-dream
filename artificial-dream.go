package main

func main() {

}

func gameLoop(g Game) {
	if g.Running() {

		g.Input()
		g.Update()
		g.Render()
	}

}

type Game interface {
	Update()
	Input()
	Render()
	Running() bool
}
