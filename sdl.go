package main

import (
	"fmt"
	"runtime"

	"github.com/veandco/go-sdl2/sdl"
)

type SDLDisplay struct {
	points     []Point
	pointCount int
	quit       bool
}

func MakeSDLDisplay(size int) *SDLDisplay {
	display := &SDLDisplay{}
	go display.run(size)
	return display
}

func (d *SDLDisplay) run(size int) {
	runtime.LockOSThread()

	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}

	window, err := sdl.CreateWindow("Chaos Game", sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED, size, size, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	fmt.Println("Created window")

	renderer, err := sdl.CreateRenderer(window, -1,
		sdl.RENDERER_SOFTWARE)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()
	fmt.Println("Created renderer")

	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()
	fmt.Println("Cleared screen")
	renderer.Present()
	fmt.Println("Presented")
	renderer.SetDrawColor(255, 255, 255, 255)
	s := float64(size)
	drawn := 0
	for !d.quit {
		for sdl.PollEvent() != nil {
		}
		if drawn > d.pointCount {
			drawn = 0
			renderer.SetDrawColor(0, 0, 0, 255)
			renderer.Clear()
			renderer.Present()
			renderer.SetDrawColor(255, 255, 255, 255)
		}

		for i := drawn; i < d.pointCount; i++ {
			drawn++
			p := d.points[i]
			r := size / 64
			if r > 3 {
				r = 3
			} else if r < 1 {
				r = 1
			}
			x, y := 0, 0
			switch len(p) {
			default:
				r = int(float64(r)*p[2]*3) + 1
				fallthrough
			case 2:
				x, y = int(s*p[0]), int(s*p[1])
			case 1:
				x, y = int(s*p[0]), size/2
			}
			DrawCircle(x, y, r, renderer)
		}
		renderer.Present()
	}
}

//taken from: https://rosettacode.org/wiki/Bitmap/Midpoint_circle_algorithm#Go
func DrawCircle(x, y, r int, renderer *sdl.Renderer) {
	if r < 0 {
		return
	}
	// Bresenham algorithm
	x1, y1, err := -r, 0, 2-2*r
	for {
		renderer.DrawLine(x+x1, y+y1, x-x1, y+y1)
		renderer.DrawLine(x-x1, y-y1, x+x1, y-y1)
		renderer.DrawLine(x+y1, y-x1, x-y1, y-x1)
		renderer.DrawLine(x-y1, y+x1, x+y1, y+x1)
		r = err
		if r > x1 {
			x1++
			err += x1*2 + 1
		}
		if r <= y1 {
			y1++
			err += y1*2 + 1
		}
		if x1 >= 0 {
			break
		}
	}
}

func (d *SDLDisplay) AddPoint(p Point) {
	d.points = append(d.points, p)
}

func (d *SDLDisplay) Display() {
	d.pointCount = len(d.points)
}

func (d *SDLDisplay) Clear() {
	d.pointCount = 0
	d.points = make([]Point, 0)
}

func (d *SDLDisplay) Quit() {
	d.quit = true
}
