package main

func main() {

}

func gameLoop(g Game) {
	g.Render()
	g.Input()
	g.Update()

}

type Game interface {
	Update()
	Input()
	Render()
}
