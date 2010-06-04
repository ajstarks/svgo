package main

import (
	svglib "svg"
	"rand"
	"time"
	"fmt"
	"os"
)

var svg = svglib.New(os.Stdout)

func main() {
	width := 200
	height := 200
	svg.Start(width, height)
	svg.Title("Random Lines")
	svg.Rect(0, 0, width, height, "fill:black")
	rand.Seed(time.Nanoseconds() % 1e9)
	svg.Gstyle("stroke-width:10")
	r := 0
	for i := 0; i < width; i++ {
		r = rand.Intn(255)
		svg.Line(i, 0, rand.Intn(width), height, fmt.Sprintf("stroke:rgb(%d,%d,%d); opacity:0.39", r, r, r))
	}
	svg.Gend()

	svg.Text(width/2, height/2, "Random Lines", "fill:white; font-size:20; font-family:Calibri; text-anchor:middle")
	svg.End()
}
