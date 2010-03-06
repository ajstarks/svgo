// svgobj: demo the SVG package

package main

import (
	"./svg"
)

const textsize = 15
const coordsize = 4
const objstyle = "fill:none; stroke-width:2; stroke:rgb(127,0,0); opacity:0.75"
const legendstyle = "fill:gray"
const titlestyle = "fill:black"
const linestyle = "stroke:gray; stroke-width:1"
const gtextstyle = "font-family:Calibri; text-anchor:middle; font-size:14"
const coordstring = "x, y"

var grayfill = svg.RGB(200, 200, 200)

func showcoordstr(x int, y int, s string) {
	svg.Circle(x, y, coordsize, grayfill)
	svg.Text(x, y-textsize, s, legendstyle)
}

func showcoord(x int, y int) {
	svg.Circle(x, y, coordsize, grayfill)
	svg.Text(x, y-textsize, coordstring, legendstyle)
}

func showlegend(x int, y int, size int, legend string) {
	svg.Text(x, y+size+textsize, legend, titlestyle)
}

func showcircle(x int, y int, r int, legend string) {
	showcoord(x, y)
	svg.Circle(x, y, r, objstyle)
	svg.Line(x, y, x+r, y, linestyle)
	svg.Text(x+(r/2), y+textsize, "r", legendstyle)
	showlegend(x, y, r, legend)
}

func showellipse(x int, y int, w int, h int, legend string) {
	showcoord(x, y)
	svg.Ellipse(x, y, w, h, objstyle)
	svg.Line(x, y, x+w, y, linestyle)
	svg.Text(x+(w/2), y+textsize, "w", legendstyle)
	svg.Line(x, y, x, y+h, linestyle)
	svg.Text(x-textsize, y+(h/2), "h", legendstyle)
	showlegend(x, y, h, legend)
}

func showrect(x int, y int, w int, h int, legend string) {
	showcoord(x, y)
	svg.Rect(x, y, w, h, objstyle)
	svg.Text(x-textsize, y+(h/2), "h", legendstyle)
	svg.Text(x+(w/2), y-textsize, "w", legendstyle)
	showlegend(x+(w/2), y, h, legend)
}

func showsquare(x int, y int, w int, legend string) {
	showcoord(x, y)
	svg.Square(x, y, w, objstyle)
	svg.Text(x+(w/2), y-textsize, "w", legendstyle)
	showlegend(x+(w/2), y, w, legend)
}

func showimage(x int, y int, w int, h int, s string, legend string) {
	showrect(x, y, w, h, legend)
	svg.Image(x, y, w, h, s, "")
}

func showline(x1 int, y1 int, x2 int, y2 int, legend string) {
	showcoordstr(x1, y1, "x1, y1")
	showcoordstr(x2, y2, "x2, y2")
	svg.Line(x1, y1, x2, y2, objstyle)
	showlegend(x1+(x2-x1)/2, y1, (y2 - y1), legend)
}

func showarc(sx int, sy int, ax int, ay int, r int, large bool, sweep bool, ex int, ey int, legend string) {
	showcoordstr(sx, sy, "sx, sy")
	showcoordstr(ex, ey, "ex, ey")
	svg.Arc(sx, sy, ax, ay, r, large, sweep, ex, ey, objstyle)
	showlegend(sx+(ex-sx)/2, sy, ax-textsize, legend)
}

func showbez(sx int, sy int, cx int, cy int, px int, py int, ex int, ey int, legend string) {
	showcoordstr(sx, sy, "sx, sy")
	showcoordstr(cx, cy, "cx, cy")
	showcoordstr(px, py, "px, py")
	showcoordstr(ex, ey, "ex, ey")
	svg.Bezier(sx, sy, cx, cy, px, py, ex, ey, objstyle)
	showlegend(px, py, 10, legend)
}

func showqbez(sx int, sy int, cx int, cy int, ex int, ey int, tx int, ty int, legend string) {
	showcoordstr(sx, sy, "sx, sy")
	showcoordstr(cx, cy, "cx, cy")
	showcoordstr(ex, ey, "ex, ey")
	showcoordstr(tx, ty, "tx, ty")
	svg.Qbezier(sx, sy, cx, cy, ex, ey, tx, ty, objstyle)
	showlegend(ex, ey, 10, legend)
}

func showroundrect(x int, y int, w int, h int, rx int, ry int, legend string) {
  xr := x+rx
  yr := y+ry
  showcoord(x, y)
  svg.Roundrect(x, y, w, h, rx, ry, objstyle)
  svg.Text(x-textsize, y+(h/2), "h", legendstyle)
	svg.Text(x+(w/2), y-textsize, "w", legendstyle)
	svg.Line(xr,y,   xr,yr, linestyle)
	svg.Line(x,yr,   xr,yr, linestyle)
	svg.Text(xr+textsize, yr-(ry/2), "ry", legendstyle)
	svg.Text(x+(rx/2), yr+textsize, "rx", legendstyle)
  showlegend(x+(w/2), y, h, legend)
}

func showpolygon(x []int, y []int, legend string) {
	for i := 0; i < len(x); i++ {
		showcoord(x[i], y[i])
	}
	svg.Polygon(x, y, objstyle)
	showlegend(x[4], y[4], 10, legend)
}

func showpolyline(x []int, y []int, legend string) {
	for i := 0; i < len(x); i++ {
		showcoord(x[i], y[i])
	}
	svg.Polyline(x, y, objstyle)
	showlegend(x[1], y[1], 20, legend)
}


var pgx = []int{60, 110,  160, 160, 110, 60, 60}
var pgy = []int{610, 585, 610, 660, 690, 660, 610}

var plx = []int{60, 160, 260, 400 }
var ply = []int{520, 490, 520, 490}

func main() {
	width := 760
	height := 760
	svg.Start(width, height)
	svg.Title("Go SVG Package")
	svg.Rect(0, 0, width, height, svg.RGB(255, 255, 255))
	svg.Desc("Objects")
	svg.Gstyle(gtextstyle)
	svg.Text(60, 40, "Go SVG Package", "font-size:24; text-anchor:start")

	showcircle(355, 320, 50, "svg.Circle(x, y, r, style)")
	showellipse(140, 320, 100, 50, "svg.Ellipse(x, y, w, h, style)")
	
	showrect(60, 100, 160, 100, "svg.Rect(x, y, w, h, style)")
	showsquare(300, 100, 100, "svg.Square(x, y, w, style)")
	showroundrect(480, 50, 250, 150, 40, 40, "svg.Roundrect(x, y, w, h, rx, ry, style)")

	showarc(480,700, 20,10, 0, false, true, 725,700, "svg.Arc(sx, sy, ax, ay, r, lflag, sflag, style)")
	showbez(480,350, 600,290, 600,400, 725,350, "svg.Bezier(sx, sy, cx, cy, px, py, ex, ey, style)")
	showqbez(480,500, 600,500, 600,550, 725,500, "svg.Qbezier(sx, sy, cx, cy, ex, ey, tx, ty, style)")
	
	showline(60,440, 400,440, "svg.Line(x1, y1, x2, y2, style)")	
	showpolygon(pgx, pgy, "svg.Polygon(x, y, style)")
	showpolyline(plx, ply, "svg.Polyline(x, y, style)")
	
	showimage(272,575, 128, 128, "images/gophercolor128x128.png", "svg.Image(x, y, w, h, name, style)")

	svg.Gend()
	svg.Desc("20 px grid")
	svg.Grid(0, 0, width, height, 20, "stroke:rgb(150,150,150); opacity:0.1")
	svg.End()
}
