package main

import (
	"os"

	svg "github.com/ajstarks/svgo"
)

func main() {
	width, height := 600, 600
	rsize := 20
	csize := rsize / 2
	duration := 5.0
	repeat := 10
	canvas := svg.New(os.Stdout)
	canvas.Start(width, height)
	canvas.Image(0, 0, 100, 140, "gopher.jpg", `id="gopher"`)
	canvas.Arc(0, 250, 10, 10, 0, false, true, 500, 250, `id="top"`, `fill="none"`, `stroke="red"`)
	canvas.Arc(0, 250, 10, 10, 0, true, false, 500, 250, `id="bot"`, `fill="none"`, `stroke="blue"`)
	canvas.Circle(0, 0, csize, `fill="red"`, `id="red-dot"`)
	canvas.Circle(0, 0, csize, `fill="blue"`, `id="blue-dot"`)
	canvas.AnimateMotion("#red-dot", "#top", duration, repeat)
	canvas.AnimateMotion("#blue-dot", "#bot", duration, repeat)
	canvas.AnimateMotion("#gopher", "#top", duration, repeat)
	canvas.End()
}
