package main

import (
	svglib "./svg"
	"os"
	"fmt"
)

var (
	width  = 500
	height = 500
	svg    = svglib.New(os.Stdout)
)

const androidcolor = "rgb(164,198,57)"

func background(v int) { svg.Rect(0, 0, width, height, svg.RGB(v, v, v)) }

func android(x, y int, fill string, opacity float) {
  var linestyle = []string{`stroke="`+fill+`"`, `stroke-linecap="round"`, `stroke-width="5"`}
	globalstyle := fmt.Sprintf("fill:%s;opacity:%.2f", fill, opacity)
	svg.Gstyle(globalstyle)
	svg.Arc(x+30, y+70, 35, 35, 0, false, true, x+130, y+70) // head
	svg.Line(x+60, y+25, x+50, y+10, linestyle[0], linestyle[1], linestyle[2])   // left antenna
	svg.Line(x+100, y+25, x+110, y+10, linestyle[0], linestyle[1], linestyle[2]) // right antenna
	svg.Circle(x+60, y+45, 5, "fill:white")                  // left eye
	svg.Circle(x+100, y+45, 5, `fill="white"`)               // right eye
	svg.Roundrect(x+30, y+75, 100, 90, 10, 10)               // body
	svg.Rect(x+30, y+75, 100, 80)
	svg.Roundrect(x+5, y+80, 20, 70, 10, 10)   // left arm
	svg.Roundrect(x+135, y+80, 20, 70, 10, 10) // right arm
	svg.Roundrect(x+50, y+150, 20, 50, 10, 10) // left leg
	svg.Roundrect(x+90, y+150, 20, 50, 10, 10) // right leg
	svg.Gend()
}

func main() {
	svg.Start(width, height)
	background(255)

	android(100, 100, androidcolor, 1.0)
	svg.Gtransform("scale(3.0,3.0)")
	android(50, 50, "gray", 0.5)
	svg.Gend()

	svg.Gtransform("scale(0.5,0.5)")
	android(100, 100, "red", 1.0)
	svg.Gend()
	svg.End()
}
