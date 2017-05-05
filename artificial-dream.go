package main

func main() {

}

func gameLoop(g Game) {
	g.Input()
	g.Update()
	g.Render()

}

type Game interface {
	Update()
	Input()
	Render()
}
