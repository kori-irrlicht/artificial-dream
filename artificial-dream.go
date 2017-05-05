package main

func main() {

}

func gameLoop(g Game) {
	g.Update()
	g.Input()

}

type Game interface {
	Update()
	Input()
}
