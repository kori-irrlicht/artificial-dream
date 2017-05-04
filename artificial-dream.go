package main

func main() {

}

func gameLoop(g Game) {
	g.Input()

}

type Game interface {
	Input()
}
