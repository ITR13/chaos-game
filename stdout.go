package main

import (
	"fmt"
	"os"
)

type stdoutDisplay []Point

func MakeStdoutDisplay() *stdoutDisplay {
	disp := make(stdoutDisplay, 0)
	return &disp
}

func (d *stdoutDisplay) Display() {
	for i := range *d {
		fmt.Println((*d)[i])
	}
	*d = make(stdoutDisplay, 0)
}

func (d *stdoutDisplay) AddPoint(point Point) {
	*d = append(*d, point)
}
func (d *stdoutDisplay) Clear() {
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
}

func (d *stdoutDisplay) Quit() {
	os.Stdout.Close()
}
