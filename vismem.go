package main

import (
	"os"
	svglib "svg"
)

var svg = svglib.New(os.Stdout)

func main() {
	width := 512
	height := 512
	n := 1024
	rowsize := 32
	diameter := 16
	var value int
	var source string

	if len(os.Args) > 1 {
		source = os.Args[1]
	} else {
		source = "/dev/urandom"
	}

	f, _ := os.Open(source, os.O_RDONLY, 0)
	mem := make([]byte, n)
	f.Read(mem)
	f.Close()

	svg.Start(width, height)
	svg.Title("Visualize Files")
	svg.Rect(0, 0, width, height, "fill:white")
	dx := diameter / 2
	dy := diameter / 2
	svg.Gstyle("fill-opacity:1.0")
	for i := 0; i < n; i++ {
		value = int(mem[i])
		if i%rowsize == 0 && i != 0 {
			dx = diameter / 2
			dy += diameter
		}
		svg.Circle(dx, dy, diameter/2, svg.RGB(value, value, value))
		dx += diameter
	}
	svg.Gend()
	svg.End()
}
