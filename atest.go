package main

import (
	"./svg"
)

var width = 500
var height = 500

func main() {

	svg.Start(width, height)
	svg.Gstyle("stroke-width:1; font-size:20; text-anchor:middle")
	for a := 10; a < 50; a += 2 {
		svg.Arc(100, 400, 20, a, 20, true, true, 400, 400, "fill:none; stroke:gray")
	}
	svg.Text(100, 410, "Start", "fill:black")
	svg.Text(400, 410, "End", "fill:black")

	svg.Gend()

	svg.Grid(0, 0, width, height, 10, "stroke:gray; opacity:0.3")
	svg.End()

}
