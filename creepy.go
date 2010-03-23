package main

import (
  "./svg"
)

var width = 500
var height = 500

func background(v int) { svg.Rect(0, 0, width, height, svg.RGB(v, v, v)) }

func smile(x, y, r int, style ...string) {

	if len(style) > 0 {
		svg.Gstyle(style[0])
	}

	svg.Roundrect(x-(r*2), y-(r*2), r*7, r*20, r*2, r*2, svg.RGB(200, 200, 200))
	svg.Circle(x, y, r, svg.RGB(127, 0, 0))
	svg.Circle(x, y, r/4, "fill:white")
	svg.Circle(x+(r*3), y, r)
	svg.Arc(x-r, y+(r*3), r/4, r/4, 0, true, false, x+(r*4), y+(r*3))

	if len(style) > 0 {
		svg.Gend()
	}

}

func creepy() {
	smile(200, 100, 10)
	svg.Gtransform("rotate(30)")
	smile(200, 100, 10)
	svg.Gend()

	svg.Gtransform("translate(50,0) scale(2,2)")
	smile(200, 100, 30, "opacity:0.3")
	svg.Gend()

}

func main() {
  svg.Start(width, height)
  background(255)
  creepy()  
  svg.End()
}

