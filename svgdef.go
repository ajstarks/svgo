// svgdef: SVG Object Definition and Use

package main

import (
	svglib "svg"
	"os"
)

var svg = svglib.New(os.Stdout)

const textsize = 15
const coordsize = 4
const objstyle = "fill:none; stroke-width:2; stroke:rgb(127,0,0); opacity:0.75"
const legendstyle = "fill:gray; text-anchor:middle"
const titlestyle = "fill:black; text-anchor:middle"
const linestyle = "stroke:black; stroke-width:1"
const gtextstyle = "font-family:Calibri; text-anchor:middle; font-size:14px"
const coordstring = "x, y"

var grayfill = svg.RGB(200, 200, 200)

func defcoordstr(x int, y int, s string) {
	svg.Circle(x, y, coordsize, grayfill)
	svg.Text(x, y-textsize, s, legendstyle)
}

func defcoord(x int, y int) {
	svg.Circle(x, y, coordsize, grayfill)
	svg.Text(x, y-textsize, coordstring, legendstyle)
}

func deflegend(x int, y int, size int, legend string) {
	svg.Text(x, y+size+textsize, legend, titlestyle)
}

func defcircle(id string, r int, legend string) {
	svg.Gid(id)
	defcoord(0, 0)
	svg.Circle(0, 0, r, objstyle)
	svg.Line(0, 0, r, 0, linestyle)
	svg.Text((r / 2), textsize, "r", legendstyle)
	deflegend(0, 0, r, legend)
	svg.Gend()
}

func defellipse(id string, w int, h int, legend string) {
	svg.Gid(id)
	defcoord(0, 0)
	svg.Ellipse(0, 0, w, h, objstyle)
	svg.Line(0, 0, w, 0, linestyle)
	svg.Line(0, 0, 0, h, linestyle)
	svg.Text((w / 2), textsize, "rx", legendstyle)
	svg.Text(-textsize, (h / 2), "ry", legendstyle)
	deflegend(0, 0, h, legend)
	svg.Gend()
}

func defrect(id string, w int, h int, legend string) {
	svg.Gid(id)
	defcoord(0, 0)
	svg.Rect(0, 0, w, h, objstyle)
	svg.Text(-textsize, (h / 2), "h", legendstyle)
	svg.Text((w / 2), -textsize, "w", legendstyle)
	deflegend((w / 2), 0, h, legend)
	svg.Gend()
}

func defsquare(id string, w int, legend string) {
	svg.Gid(id)
	defcoord(0, 0)
	svg.Square(0, 0, w, objstyle)
	svg.Text((w / 2), -textsize, "w", legendstyle)
	deflegend((w / 2), 0, w, legend)
	svg.Gend()
}

func defimage(id string, w int, h int, s string, legend string) {
	svg.Gid(id)
	defcoord(0, 0)
	svg.Rect(0, 0, w, h, objstyle)
	svg.Text(-textsize, (h / 2), "h", legendstyle)
	svg.Text((w / 2), -textsize, "w", legendstyle)
	svg.Image(0, 0, w, h, s)
	deflegend(w/2, h, 0, legend)
	svg.Gend()
}

func defline(id string, size int, legend string) {
	svg.Gid(id)
	defcoordstr(0, 0, "x1, y1")
	defcoordstr(size, 0, "x2, y2")
	svg.Line(0, 0, size, 0, objstyle)
	deflegend(size/2, textsize, -5, legend)
	svg.Gend()
}

func defarc(id string, w int, h int, legend string) {
	svg.Gid(id)
	defcoordstr(0, 0, "sx, sy")
	defcoordstr(w*2, 0, "ex, ey")
	svg.Arc(0, 0, w, h, 0, false, true, w*2, 0, objstyle)
	deflegend(w, 0, 0, legend)
	svg.Gend()
}

func defbez(id string, px int, py int, legend string) {
	svg.Gid(id)
	defcoordstr(0, 0, "sx, sy")
	defcoordstr(px, -py, "cx, cy")
	defcoordstr(px, py, "px, py")
	defcoordstr(px*2, 0, "ex, ey")
	svg.Bezier(0, 0, px, -py, px, py, px*2, 0, objstyle)
	deflegend(px, py, 10, legend)
	svg.Gend()
}

func defqbez(id string, ex int, ey int, legend string) {
	svg.Gid(id)
	defcoordstr(0, 0, "sx, sy")
	defcoordstr(ex, -ey, "cx, cy")
	defcoordstr(ex, ey, "ex, ey")
	defcoordstr(ex*2, 0, "tx, ty")
	svg.Qbezier(0, 0, ex, -ey, ex, ey, ex*2, 0, objstyle)
	deflegend(ex, ey, 30, legend)
	svg.Gend()
}

func defroundrect(id string, w int, h int, rx int, ry int, legend string) {
	svg.Gid(id)
	defcoord(0, 0)
	svg.Roundrect(0, 0, w, h, rx, ry, objstyle)
	svg.Text(-textsize, (h / 2), "h", legendstyle)
	svg.Text((w / 2), -textsize, "w", legendstyle)
	svg.Line(rx, 0, rx, ry, linestyle)
	svg.Line(0, ry, rx, ry, linestyle)
	svg.Text(rx+textsize, ry-(ry/2), "ry", legendstyle)
	svg.Text((rx / 2), ry+textsize, "rx", legendstyle)
	deflegend((w / 2), 0, h, legend)
	svg.Gend()
}

func defpolygon(id string, w int, h int, legend string) {
	var x = []int{0, w / 2, w, w, w / 2, 0}
	var y = []int{0, h / 5, 0, (h * 3) / 4, h, (h * 3) / 4}
	svg.Gid(id)
	for i := 0; i < len(x); i++ {
		defcoord(x[i], y[i])
	}
	svg.Polygon(x, y, objstyle)
	deflegend(x[4], y[4], 10, legend)
	svg.Gend()
}

func defpolyline(id string, w int, h int, legend string) {
	var x = []int{0, w / 2, (w * 3) / 4, w}
	var y = []int{0, -(h / 2), -(h / 3), -h}
	svg.Gid(id)
	for i := 0; i < len(x); i++ {
		defcoord(x[i], y[i])
	}
	svg.Polyline(x, y, objstyle)
	deflegend(x[1], y[1], 20, legend)
	svg.Gend()
}


func main() {

	width := 760
	height := 760
	svg.Start(width, height)
	svg.Title("SVG Go Library Description")
	svg.Rect(0, 0, width, height, "fill:white")
	svg.Gstyle(gtextstyle)
	svg.Text(width/2, 40, "SVG Go Library", "font-size:24")

	svg.Desc("Object Definitions")
	svg.Def()
	defcircle("circle", 50, "Circle(x, y, r,...)")
	defellipse("ellipse", 75, 50, "Ellipse(x, y, rx ,ry,...)")
	defrect("rectangle", 160, 100, "Rect(x, y, w, h,...)")
	defroundrect("roundrect", 160, 100, 25, 25, "Roundrect(x,y,rx,ry,...)")
	defsquare("square", 100, "Square(x, y, w,...)")
	defimage("image", 128, 128, "images/gophercolor128x128.png", "Image(x, y, w, h, path,...)")
	defarc("arc", 90, 40, "Arc(sx, sy, ax, ay, r, lflag, sflag, ex, ey, ...)")
	defline("line", 240, "Line(x1, y1, x2, y2, ...)")
	defbez("bezier", 120, 60, "Bezier(sx, sy, cx, cy, px, py, ex, ey, ...)")
	defqbez("qbez", 120, 40, "Qbezier(sx, sy, cx, cy, ex, ey, tx, ty, ...)")
	defpolygon("polygon", 160, 120, "Polygon(x, y, ...)")
	defpolyline("polyline", 240, 40, "Polyline(x, y, ...)")
	svg.DefEnd()

	svg.Desc("Object Usage")
	svg.Use(40, 80, "#rectangle")
	svg.Use(40, 240, "#roundrect")
	svg.Use(40, 580, "#polygon")
	svg.Use(120, 420, "#ellipse")
	svg.Use(260, 280, "#arc")
	svg.Use(280, 580, "#image")
	svg.Use(300, 80, "#square")
	svg.Use(345, 420, "#circle")
	svg.Use(480, 280, "#polyline")
	svg.Use(480, 140, "#line")
	svg.Use(480, 420, "#bezier")
	svg.Use(480, 620, "#qbez")
	svg.Gend()

	svg.Grid(0, 0, width, height, 20, "stroke:black;opacity:0.1")
	svg.End()

}
