package main

import (
	"fmt"
)

type AsciiDisplay struct {
	grid    [][]uint8
	twoChar bool
}

func MakeAsciiDisplay(size int, twoChar bool) *AsciiDisplay {
	grid := make([][]uint8, size)
	for i := range grid {
		grid[i] = make([]uint8, size)
	}
	return &AsciiDisplay{grid, twoChar}
}

func (d *AsciiDisplay) Display() {
	fmt.Println()
	for y := range d.grid {
		for x := range d.grid[y] {
			v := d.grid[y][x]
			if d.twoChar {
				if v == 0 {
					fmt.Print("**")
				} else {
					fmt.Printf("%02X", v)
				}
			} else if v == 0 {
				fmt.Print("*")
			} else if v < 16 {
				fmt.Printf("%01X", d.grid[y][x])
			} else {
				fmt.Print("F")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (d *AsciiDisplay) AddPoint(point Point) {
	if len(point) != 2 {
		panic("Ascii display currently only accepts 2 dimensions!")
	}
	s := float64(len(d.grid))
	x, y := int(point[0]*s), int(point[1]*s)
	if d.grid[y][x] != 0xff {
		d.grid[y][x]++
	}
}

func (d *AsciiDisplay) Clear() {
	for y := range d.grid {
		for x := range d.grid[y] {
			d.grid[y][x] = 0
		}
	}
}
func (d *AsciiDisplay) Quit() {}
