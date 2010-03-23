package main

import (
	"./svg"
)

var width = 600
var height = 600

func box3d(x, y, w, h, d int, style string) {

	var px1 = []int{x, x + d, x + w + d, x + w}
	var py1 = []int{y, y - d, y - d, y}

	var px2 = []int{x + w + d, x + w + d, x + w, x + w}
	var py2 = []int{y - d, (y + h) - d, y + h, y}

	svg.Gstyle(style)
	svg.Rect(x, y, w, h, "")
	svg.Polyline(px1, py1, "")
	svg.Polyline(px2, py2, "")
	svg.Gend()
}

func drum(x, y, w, h int, style string) {
	svg.Ellipse(x, y, w, w/4, style)
	svg.Line(x-w, y-10, x-w, y+h, style)
	svg.Line(x+w, y-10, x+w, y+h, style)
	svg.Arc(x-w, y+h, w, w/4, 0, true, false, x+w, y+h, style)
}

func servermap(x, y, w, h int) {

	v := 240
	//mid := ((x+w) - x)/2

	svg.Rect(x, y, w, h, svg.RGB(v, v, v))
	v -= 15
	svg.Rect(x+(w/20), y+(w/20), w-(w/10), 2*(h/3), svg.RGB(v, v, v))
	v -= 15
	svg.Rect(x+(w/10), y+(w/10), w-(w/5), h/3, svg.RGB(v, v, v))
}


func main() {

	svg.Start(width, height)
	/*
	  svg.Rect(100, 150, 200, 150, "fill:rgb(200,200,255)")
	  svg.Rect(110, 190, 180, 100, "fill:rgb(150,150,255)")
	  svg.Rect(120, 230, 160, 50, "fill:rgb(100,100,255)")
	  svg.Text(100 + (200/2), 150+20, "One", "text-anchor:middle")
	  svg.Text(110 + (180/2), 190+20, "two", "text-anchor:middle")
	  svg.Text(120 + (160/2), 230+20, "three", "text-anchor:middle")
	*/

	//servermap(0,0 , width, height  )

	//servermap(100, 100, 300, 300)

	servermap(10, 100, 200, 400)
	servermap(250, 100, 300, 150)
	servermap(260, 260, 240, 240)
	//servermap(0,0,width,height)


	/*
	  box3d(200, 200, 160, 80, 20, "stroke:red; stroke-width:2; fill:none")
	  box3d(100, 100, 80, 40, 10, "stroke:blue; fill:none")
	  box3d(200, 400, 320, 160, 10, "stroke:black; fill:none")
	  box3d(300, 50, 160, 80, 40, "stroke:green; fill:none")
	  box3d(50, 400, 50, 50, 25, "fill:red; stroke:white")

	  drum(200, 300, 100, 200, "fill:none; stroke-width:10; stroke:blue")

	  svg.Grid(0, 0, width, height, 10, "stroke:black;opacity:0.1")
	*/
	svg.End()

}
