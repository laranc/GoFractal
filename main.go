package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"runtime"

	"github.com/laranc/monorepo/engine/graphics2d"
	"github.com/veandco/go-sdl2/sdl"
	"gonum.org/v1/gonum/floats"
)

const (
	title         = "GoFractals"
	width   int32 = 680
	height  int32 = 460
	maxIter uint8 = 100
)

var (
	renderer graphics2d.Renderer2D
	viewMin  = -0.5
	viewMax  = 0.5
	x        []float64
	y        []float64
	img      [width][height]uint8
)

func init() {
	runtime.LockOSThread()
	x = make([]float64, width)
	y = make([]float64, height)
}

func main() {
	var err error
	if renderer, err = graphics2d.MakeRenderer2D(title, width, height); err != nil {
		panic(err)
	}
	run()
	fmt.Println("Closing...")
}

func run() {
	running := true
	for running {
		for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			switch t := e.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				switch t.Keysym.Sym {
				case sdl.K_ESCAPE:
					running = false
				default:
					break

				}
			default:
				break
			}
		}
		keyboard()
		mandelbrotSet(viewMin, viewMax, viewMin, viewMax)
		renderer.RenderBegin()
		for i := range width {
			for j := range height {
				n := img[i][j]
				bright := bound(float64(n), 0, float64(maxIter), 0, 1)
				bright = bound(math.Sqrt(float64(n)), 0, 1, 0, 255)
				if n == maxIter {
					bright = 0
				}
				renderer.DrawRect(sdl.Rect{X: i, Y: j, W: 1, H: 1}, sdl.Color{R: bright, G: bright, B: bright, A: 255})
			}
		}
		renderer.RenderEnd()
	}
}

func keyboard() {
	keys := sdl.GetKeyboardState()
	if keys[sdl.SCANCODE_UP] != 0 {
		viewMin -= 0.01
		if viewMin < -2.5 {
			viewMin = -2.5
		}
	}
	if keys[sdl.SCANCODE_DOWN] != 0 {
		viewMin += 0.01
		if viewMin > 0 {
			viewMin = 0
		}

	}
	if keys[sdl.SCANCODE_LEFT] != 0 {
		viewMax -= 0.01
		if viewMax < 0 {
			viewMax = 0
		}
	}
	if keys[sdl.SCANCODE_RIGHT] != 0 {
		viewMax += 0.01
		if viewMax > 2.5 {
			viewMax = 2.5
		}
	}
}

func bound(n, inMin, inMax, outMin, outMax float64) uint8 {
	return uint8((n-inMin)*(outMax-outMin)/(inMax-inMin) + outMin)
}

func mandelbrot(c complex128) uint8 {
	z := complex(0, 0)
	for n := range maxIter {
		if cmplx.Abs(z) > 2 {
			return n
		}
		z = z*z + c
	}
	return maxIter
}

func mandelbrotSet(xmin, xmax, ymin, ymax float64) {
	floats.Span(x, xmin, xmax)
	floats.Span(y, ymin, ymax)
	for i := range width {
		for j := range height {
			c := complex(x[i], y[j])
			img[i][j] = mandelbrot(c)
		}
	}
}
