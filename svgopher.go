package main

import (
	"svg"
	"os"
)

var (
	width  = 500
	height = 400
	canvas    = svg.New(os.Stdout)
)

func background(v int) { canvas.Rect(0, 0, width, height, canvas.RGB(v, v, v)) }

func gordon(x, y, w, h int) {

	w10 := w / 10
	w12 := w / 12
	w2 := w / 2
	w3 := w / 3
	w8 := w / 8
	w6 := w / 6
	xw := x + w
	h23 := (h * 2) / 3

	blf := "fill:black"
	wf := "fill:white"
	nf := "fill:brown"
	brf := "fill:brown; fill-opacity:0.2"

	canvas.Gstyle("fill:none; stroke:none")
	canvas.Roundrect(x, y, w, h*4, w/2, w/2, "fill:brown; fill-opacity:0.4")
	canvas.Circle(x, y+h, w12, brf) // left ear
	canvas.Circle(x, y+h, w12-10, nf)

	canvas.Circle(x+w, y+h, w12, brf) // right ear
	canvas.Circle(x+w, y+h, w12-10, nf)

	canvas.Circle(x+w3, y+h23, w8, wf) // left eye
	canvas.Circle(x+w3+10, y+h23, w10-10, blf)
	canvas.Circle(x+w3+15, y+h23, 5, wf)

	canvas.Circle(xw-w3, y+h23, w8, wf) // right eye
	canvas.Circle(xw-w3+10, y+h23, w10-10, blf)
	canvas.Circle(xw-(w3)+15, y+h23, 5, wf)

	canvas.Roundrect(x+w2-w8, y+h+30, w8, w6, 5, 5, wf) // left tooth
	canvas.Roundrect(x+w2, y+h+30, w8, w6, 5, 5, wf)    // right tooth

	canvas.Ellipse(x+(w2), y+h+30, w6, w12, nf)   // snout
	canvas.Ellipse(x+(w2), y+h+10, w10, w12, blf) // nose

	canvas.Circle(x-20, y+h+120, w3, wf) // "bite"
	canvas.Gend()
}

func main() {
	canvas.Start(width, height)
	canvas.Title("SVG Gopher")
	background(255)
	canvas.Gtransform("translate(100, 100)")
	canvas.Gtransform("rotate(-30)")
	gordon(48, 48, 240, 72)
	canvas.Gend()
	canvas.Gend()
	canvas.Link("svgdef.svg", "SVG Spec & Usage")
	canvas.Text(90, 142, "SVG", "font-family:Calibri; font-size:84; fill:brown")
	canvas.LinkEnd()
	canvas.End()
}
