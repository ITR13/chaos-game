// chaos-game project main.go
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Display interface {
	Display()
	AddPoint(Point)
	Clear()
	Quit()
}

var (
	size           *int
	displayS       *string
	delay          *int
	exponential    *bool
	lerpDist       *float64
	N              *int64
	startingPoints *int
	dimensions     *int
	swap           *int
	swapchance     *float64
	infinite       *bool
)

func init() {
	size = flag.Int("display-size", 64, "sets the size of the display")

	displayS = flag.String("display", "ascii", "how to display the points")
	delay = flag.Int("delay", 2, "How many points will be added before "+
		"displaying")
	exponential = flag.Bool("linear", true, "If the points needed to be "+
		"added before displaying will not increase over time")

	lerpDist = flag.Float64("lerp", 0.5, "How far the point will "+
		"move towards the randomly chosen point")

	N = flag.Int64("n", -1, "How long to iterate")

	startingPoints = flag.Int("p", 3, "Amount of origin-points")
	dimensions = flag.Int("dim", 2, "Amount of dimensions")

	swap = flag.Int("swap", 1, "Amount of structures to swap between")
	swapchance = flag.Float64("swapchance", 0.25, "Chance to swap to the "+
		"next structure")

	infinite = flag.Bool("inf", false, "Will it go on forever?")

	flag.Parse()
}

func main() {
	rand.Seed(time.Now().Unix())

	var display Display
	switch *displayS {
	case "ascii":
		display = MakeAsciiDisplay(*size, true)
	case "stdout":
		display = MakeStdoutDisplay()
	case "sdl":
		fallthrough
	case "sdl2":
		display = MakeSDLDisplay(*size)
	default:
		panic("Got unknown value for -display\nLegal values are \"ascii\"" +
			" \"stdout\"")
	}

recreate:

	games := make([]*Game, *swap)
	for i := range games {
		game := MakeGame(*dimensions, *startingPoints, *lerpDist)
		for i := range game.origin {
			display.AddPoint(game.origin[i])
		}
		games[i] = game
	}
	display.AddPoint(games[0].active)
	swapGame := MakeSwitchingGame(*swapchance, games...)

	cdelay := *delay
	n := *N
	for n != 0 {
		if cdelay > 0 {
			for i := 0; i < cdelay && n != 0; i++ {
				display.AddPoint(swapGame.Iterate())
				n--
			}
			display.Display()
			if *exponential && cdelay*(*delay) > 0 {
				cdelay *= *delay
			}
		} else {
			for n > 0 {
				display.AddPoint(swapGame.Iterate())
				n--
			}
		}
	}
	display.Display()
	var s string
	fmt.Scanln(&s)
	s = strings.ToLower(s)

	if *infinite && s != "quit" && s != "close" {
		display.Clear()
		goto recreate
	}

	display.Quit()
}
