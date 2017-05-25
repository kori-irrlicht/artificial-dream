package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	mapper "github.com/birkirb/loggers-mapper-logrus"
	"github.com/kori-irrlicht/artificial-dream/core"
	"github.com/kori-irrlicht/artificial-dream/input"
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

	sdl.GameControllerAddMapping("030000004c050000cc09000011810000,PS4 Controller,a:b0,b:b1,back:b8,dpdown:h0.4,dpleft:h0.8,dpright:h0.2,dpup:h0.1,guide:b10,leftshoulder:b4,leftstick:b11,lefttrigger:a2,leftx:a0,lefty:a1,rightshoulder:b5,rightstick:b12,righttrigger:a5,rightx:a3,righty:a4,start:b9,x:b3,y:b2,")

	sdl.GameControllerOpen(0)
	core.GameLoop(&game{
		running:    true,
		controller: input.NewGameController()})

	defer window.Destroy()

	sdl.Delay(1000)
	sdl.Quit()

}
