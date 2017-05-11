package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	mapper "github.com/birkirb/loggers-mapper-logrus"
	"github.com/kori-irrlicht/artificial-dream/core"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	l := logrus.New()
	l.Out = os.Stdout
	l.Level = logrus.InfoLevel

	core.SetLogger(mapper.NewLogger(l))
	sdl.Init(sdl.INIT_EVERYTHING)

	window, err := sdl.CreateWindow("Artificial Dream", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)

	if err != nil {
		panic(err)
	}

	core.GameLoop(&game{running: true})

	defer window.Destroy()

	sdl.Delay(1000)
	sdl.Quit()

}
