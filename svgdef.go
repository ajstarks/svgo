// svgdef: SVG Object Definition and Use

package main

import (
	"svg"
	"os"
)

var canvas = svg.New(os.Stdout)

const textsize = 15
const coordsize = 4
const objstyle = "fill:none; stroke-width:2; stroke:rgb(127,0,0); opacity:0.75"
const legendstyle = "fill:gray; text-anchor:middle"
const titlestyle = "fill:black; text-anchor:middle"
const linestyle = "stroke:black; stroke-width:1"
const gtextstyle = "font-family:Calibri; text-anchor:middle; font-size:14px"
const coordstring = "x, y"

var grayfill = canvas.RGB(200, 200, 200)

func defcoordstr(x int, y int, s string) {
	canvas.Circle(x, y, coordsize, grayfill)
	canvas.Text(x, y-textsize, s, legendstyle)
}

func defcoord(x int, y int) {
	canvas.Circle(x, y, coordsize, grayfill)
	canvas.Text(x, y-textsize, coordstring, legendstyle)
}

func deflegend(x int, y int, size int, legend string) {
	canvas.Text(x, y+size+textsize, legend, titlestyle)
}

func defcircle(id string, r int, legend string) {
	canvas.Gid(id)
	defcoord(0, 0)
	canvas.Circle(0, 0, r, objstyle)
	canvas.Line(0, 0, r, 0, linestyle)
	canvas.Text((r / 2), textsize, "r", legendstyle)
	deflegend(0, 0, r, legend)
	canvas.Gend()
}

func defellipse(id string, w int, h int, legend string) {
	canvas.Gid(id)
	defcoord(0, 0)
	canvas.Ellipse(0, 0, w, h, objstyle)
	canvas.Line(0, 0, w, 0, linestyle)
	canvas.Line(0, 0, 0, h, linestyle)
	canvas.Text((w / 2), textsize, "rx", legendstyle)
	canvas.Text(-textsize, (h / 2), "ry", legendstyle)
	deflegend(0, 0, h, legend)
	canvas.Gend()
}

func defrect(id string, w int, h int, legend string) {
	canvas.Gid(id)
	defcoord(0, 0)
	canvas.Rect(0, 0, w, h, objstyle)
	canvas.Text(-textsize, (h / 2), "h", legendstyle)
	canvas.Text((w / 2), -textsize, "w", legendstyle)
	deflegend((w / 2), 0, h, legend)
	canvas.Gend()
}

func defsquare(id string, w int, legend string) {
	canvas.Gid(id)
	defcoord(0, 0)
	canvas.Square(0, 0, w, objstyle)
	canvas.Text((w / 2), -textsize, "w", legendstyle)
	deflegend((w / 2), 0, w, legend)
	canvas.Gend()
}

func defimage(id string, w int, h int, s string, legend string) {
	canvas.Gid(id)
	defcoord(0, 0)
	canvas.Rect(0, 0, w, h, objstyle)
	canvas.Text(-textsize, (h / 2), "h", legendstyle)
	canvas.Text((w / 2), -textsize, "w", legendstyle)
	canvas.Image(0, 0, w, h, s)
	deflegend(w/2, h, 0, legend)
	canvas.Gend()
}

func defline(id string, size int, legend string) {
	canvas.Gid(id)
	defcoordstr(0, 0, "x1, y1")
	defcoordstr(size, 0, "x2, y2")
	canvas.Line(0, 0, size, 0, objstyle)
	deflegend(size/2, textsize, -5, legend)
	canvas.Gend()
}

func defarc(id string, w int, h int, legend string) {
	canvas.Gid(id)
	defcoordstr(0, 0, "sx, sy")
	defcoordstr(w*2, 0, "ex, ey")
	canvas.Arc(0, 0, w, h, 0, false, true, w*2, 0, objstyle)
	deflegend(w, 0, 0, legend)
	canvas.Gend()
}

func defbez(id string, px int, py int, legend string) {
	canvas.Gid(id)
	defcoordstr(0, 0, "sx, sy")
	defcoordstr(px, -py, "cx, cy")
	defcoordstr(px, py, "px, py")
	defcoordstr(px*2, 0, "ex, ey")
	canvas.Bezier(0, 0, px, -py, px, py, px*2, 0, objstyle)
	deflegend(px, py, 10, legend)
	canvas.Gend()
}

func defqbez(id string, ex int, ey int, legend string) {
	canvas.Gid(id)
	defcoordstr(0, 0, "sx, sy")
	defcoordstr(ex, -ey, "cx, cy")
	defcoordstr(ex, ey, "ex, ey")
	defcoordstr(ex*2, 0, "tx, ty")
	canvas.Qbezier(0, 0, ex, -ey, ex, ey, ex*2, 0, objstyle)
	deflegend(ex, ey, 30, legend)
	canvas.Gend()
}

func defroundrect(id string, w int, h int, rx int, ry int, legend string) {
	canvas.Gid(id)
	defcoord(0, 0)
	canvas.Roundrect(0, 0, w, h, rx, ry, objstyle)
	canvas.Text(-textsize, (h / 2), "h", legendstyle)
	canvas.Text((w / 2), -textsize, "w", legendstyle)
	canvas.Line(rx, 0, rx, ry, linestyle)
	canvas.Line(0, ry, rx, ry, linestyle)
	canvas.Text(rx+textsize, ry-(ry/2), "ry", legendstyle)
	canvas.Text((rx / 2), ry+textsize, "rx", legendstyle)
	deflegend((w / 2), 0, h, legend)
	canvas.Gend()
}

func defpolygon(id string, w int, h int, legend string) {
	var x = []int{0, w / 2, w, w, w / 2, 0}
	var y = []int{0, h / 5, 0, (h * 3) / 4, h, (h * 3) / 4}
	canvas.Gid(id)
	for i := 0; i < len(x); i++ {
		defcoord(x[i], y[i])
	}
	canvas.Polygon(x, y, objstyle)
	deflegend(x[4], y[4], 10, legend)
	canvas.Gend()
}

func defpolyline(id string, w int, h int, legend string) {
	var x = []int{0, w / 2, (w * 3) / 4, w}
	var y = []int{0, -(h / 2), -(h / 3), -h}
	canvas.Gid(id)
	for i := 0; i < len(x); i++ {
		defcoord(x[i], y[i])
	}
	canvas.Polyline(x, y, objstyle)
	deflegend(x[1], y[1], 20, legend)
	canvas.Gend()
}


func main() {

	width := 760
	height := 760
	canvas.Start(width, height)
	canvas.Title("SVG Go Library Description")
	canvas.Rect(0, 0, width, height, "fill:white")
	canvas.Gstyle(gtextstyle)
	canvas.Text(width/2, 40, "SVG Go Library", "font-size:24")

	canvas.Desc("Object Definitions")
	canvas.Def()
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
	canvas.DefEnd()

	canvas.Desc("Object Usage")
	canvas.Use(40, 80, "#rectangle")
	canvas.Use(40, 240, "#roundrect")
	canvas.Use(40, 580, "#polygon")
	canvas.Use(120, 420, "#ellipse")
	canvas.Use(260, 280, "#arc")
	canvas.Use(280, 580, "#image")
	canvas.Use(300, 80, "#square")
	canvas.Use(345, 420, "#circle")
	canvas.Use(480, 280, "#polyline")
	canvas.Use(480, 140, "#line")
	canvas.Use(480, 420, "#bezier")
	canvas.Use(480, 620, "#qbez")
	canvas.Gend()

	canvas.Grid(0, 0, width, height, 20, "stroke:black;opacity:0.1")
	canvas.End()

}
