package main

import (
	svglib "./svg"
	"os"
)

var svg = svglib.New(os.Stdout)

var width = 320
var height = 480

func funnel(bg int, fg int, grid int, dim int) {
	h := dim / 2
	svg.Rect(0, 0, width, height, svg.RGB(bg, bg, bg))
	for size := grid; size < width; size += grid {
		svg.Ellipse(h, size, size/2, size/2, svg.RGBA(fg, fg, fg, 0.2))
	}
}

func main() {
	svg.Start(width, height)
	svg.Title("Funnel")
	funnel(0, 255, 25, width)
	svg.End()
}
