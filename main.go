package main

import (
	"runtime"

	"github.com/laranc/monorepo/engine/graphics2d"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	title  = "Template"
	width  = 680
	height = 460
)

var (
	renderer graphics2d.Renderer2D
)

func init() {
	runtime.LockOSThread()
}

func main() {
	var err error

	renderer, err = graphics2d.MakeRenderer2D(title, width, height)
	if err != nil {
		panic(err)
	}
	run()
}

func run() {
	running := true
	for running {
		for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			switch e.(type) {
			case *sdl.QuitEvent:
				running = false
			default:
				break
			}
		}
		renderer.RenderBegin()
		renderer.RenderEnd()
	}
}
