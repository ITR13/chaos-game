package main

import (
	"math/rand"
)

type Point []float64

type SwitchingGame struct {
	games        []*Game
	currentGame  int
	switchChance float64
}

type Game struct {
	dimensions int
	origin     []Point
	active     Point
	lerpDist   float64
	points     []Point
}

func lerpPoint(a, b Point, d float64) Point {
	if len(a) != len(b) {
		panic("Tried to Lerp points of different dimensions")
	}
	point := make(Point, len(a))
	for i := range point {
		point[i] = a[i]*d + b[i]*(1-d)
	}
	return point
}

func (point Point) Copy() Point {
	p := make(Point, len(point))
	copy(p, point)
	return p
}

func (game *Game) RandomPoint() Point {
	point := make(Point, game.dimensions)
	for i := range point {
		point[i] = rand.Float64()
	}
	return point
}

func (game *Game) Point(v float64) Point {
	point := make(Point, game.dimensions)
	for i := range point {
		point[i] = v
	}
	return point
}

func (point Point) Min(other Point) {
	for i := range point {
		if point[i] > other[i] {
			point[i] = other[i]
		}
	}
}

func (point Point) Max(other Point) {
	for i := range point {
		if point[i] < other[i] {
			point[i] = other[i]
		}
	}
}

func (point Point) Sub(other Point) {
	for i := range other {
		point[i] -= other[i]
	}
}

func (point Point) Mul(v float64) {
	for i := range point {
		point[i] *= v
	}
}

func MakeGame(dimensions, points int, lerpDist float64) *Game {
	game := &Game{dimensions, make([]Point, points),
		Point{}, lerpDist, make([]Point, 0)}
	min, max := game.Point(1), game.Point(0)
	for i := 0; i < points; i++ {
		game.origin[i] = game.RandomPoint()
		min.Min(game.origin[i])
		max.Max(game.origin[i])
	}
	game.active = game.RandomPoint()
	min.Min(game.active)
	max.Max(game.active)
	d := float64(1)
	for i := range min {
		d2 := max[i] - min[i]
		if d2 > d {
			d = d2
		}
	}
	d = 1 / d
	for i := 0; i < points; i++ {
		game.origin[i].Sub(min)
		game.origin[i].Mul(d)
	}
	game.active.Sub(min)
	game.active.Mul(d)
	return game
}

func MakeSwitchingGame(chance float64, game ...*Game) *SwitchingGame {
	return &SwitchingGame{game, 0, chance}
}

func (game *Game) Iterate() Point {
	r := rand.Int() % len(game.origin)
	game.points = append(game.points, game.active)
	game.active = lerpPoint(game.active, game.origin[r], game.lerpDist)
	return game.active.Copy()
}

func (game *SwitchingGame) Iterate() Point {
	if rand.Float64() < game.switchChance {
		prevCurrentGame := game.currentGame
		game.currentGame = (game.currentGame + 1) % (len(game.games))
		game.games[game.currentGame].active = game.games[prevCurrentGame].active
	}
	return game.games[game.currentGame].Iterate()
}

func (game *Game) ClearStored() {
	game.points = make([]Point, 0)
}

func (game *SwitchingGame) ClearStored() {
	for i := range game.games {
		game.games[i].ClearStored()
	}
}
