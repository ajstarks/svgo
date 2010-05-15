package main

import (
	svglib "./svg"
	"os"
)

var (
	width  = 500
	height = 400
	svg    = svglib.New(os.Stdout)
)

func background(v int) { svg.Rect(0, 0, width, height, svg.RGB(v, v, v)) }

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

	svg.Gstyle("fill:none; stroke:none")
	svg.Roundrect(x, y, w, h*4, w/2, w/2, "fill:brown; fill-opacity:0.4")
	svg.Circle(x, y+h, w12, brf) // left ear
	svg.Circle(x, y+h, w12-10, nf)

	svg.Circle(x+w, y+h, w12, brf) // right ear
	svg.Circle(x+w, y+h, w12-10, nf)

	svg.Circle(x+w3, y+h23, w8, wf) // left eye
	svg.Circle(x+w3+10, y+h23, w10-10, blf)
	svg.Circle(x+w3+15, y+h23, 5, wf)

	svg.Circle(xw-w3, y+h23, w8, wf) // right eye
	svg.Circle(xw-w3+10, y+h23, w10-10, blf)
	svg.Circle(xw-(w3)+15, y+h23, 5, wf)

	svg.Roundrect(x+w2-w8, y+h+30, w8, w6, 5, 5, wf) // left tooth
	svg.Roundrect(x+w2, y+h+30, w8, w6, 5, 5, wf)    // right tooth

	svg.Ellipse(x+(w2), y+h+30, w6, w12, nf)   // snout
	svg.Ellipse(x+(w2), y+h+10, w10, w12, blf) // nose

	svg.Circle(x-20, y+h+120, w3, wf) // "bite"
	svg.Gend()
}

func main() {
	svg.Start(width, height)
	background(255)
	svg.Gtransform("translate(100, 100)")
	svg.Gtransform("rotate(-30)")
	gordon(48, 48, 240, 72)
	svg.Gend()
	svg.Gend()
	svg.Text(90, 142, "SVG", "font-family:Calibri; font-size:84; fill:brown")
	svg.End()
}
