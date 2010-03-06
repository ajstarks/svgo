package main

import (
	"./svg"
	"fmt"
	"os"
)

func main() {
	width := 768
	height := 320
	image := os.Args[1]
	svg.Start(width, height)
	opacity := 1.0
	for i := 0; i < width-128; i += 100 {
		svg.Image(i, 0, 128, 128, image, fmt.Sprintf("opacity:%.2f", opacity))
		opacity -= 0.10
	}
	svg.Grid(0, 0, width, height, 16, "stroke:gray; opacity:0.2")

	svg.End()
}
