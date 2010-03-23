// richter: inspired by Gerhard Richter's 256 colors, 1974

package main

import (
	svglib "./svg"
	"rand"
	"time"
	"os"
)

var svg = svglib.New(os.Stdout)

var width = 700
var height = 400

func main() {
	rand.Seed(time.Nanoseconds() % 1e9)
	svg.Start(width, height)
	svg.Title("Richter")
	svg.Rect(0, 0, width, height, "fill:white")
	rw := 32
	rh := 18
	margin := 5
	for i, x := 0, 20; i < 16; i++ {
		x += (rw + margin)
		for j, y := 0, 20; j < 16; j++ {
			svg.Rect(x, y, rw, rh, svg.RGB(rand.Intn(255), rand.Intn(255), rand.Intn(255)))
			y += (rh + margin)
		}
	}
	svg.End()
}
