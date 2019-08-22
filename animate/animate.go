package main

import (
	"os"

	svg "github.com/ajstarks/svgo"
)

func main() {
	width, height := 500, 500
	rsize := 100
	csize := rsize / 2
	duration := 5.0
	repeat := 5
	imw, imh := 100, 144
	canvas := svg.New(os.Stdout)
	canvas.Start(width, height)
	canvas.Circle(csize, csize, csize, `fill="red"`, `id="circle"`)
	canvas.Image((width/2)-(imw/2), 0, imw, imh, "gopher.jpg", `id="image"`)
	canvas.Square(width-rsize, 0, rsize, `fill="blue"`, `id="square"`)
	canvas.Animate("#circle", "cx", 0, width, duration, repeat)
	canvas.Animate("#circle", "cy", 0, height, duration, repeat)
	canvas.Animate("#square", "x", width, 0, duration, repeat)
	canvas.Animate("#square", "y", height, 0, duration, repeat)
	canvas.Animate("#image", "y", 0, height, duration, repeat)
	canvas.End()
}
