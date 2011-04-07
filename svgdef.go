// svgdef: SVG Object Definition and Use

package main

import (
	"svg"
	"os"
)

const (
	textsize    = 15
	coordsize   = 4
	objstyle    = "fill:none; stroke-width:2; stroke:rgb(127,0,0)"
	fobjstyle   = "fill:rgb(127,0,0);fill-opacity:0.25"
	legendstyle = "fill:gray; text-anchor:middle"
	titlestyle  = "fill:black; text-anchor:middle;font-size:16px"
	linestyle   = "stroke:black; stroke-width:1"
	gtextstyle  = "font-family:Calibri; text-anchor:middle; font-size:14px"
	coordstring = "x, y"
)

var (
	canvas   = svg.New(os.Stdout)
	grayfill = canvas.RGB(200, 200, 200)
	oc1      = svg.Offcolor{0, "red", 1.0}
	oc2      = svg.Offcolor{50, "gray", 1.0}
	oc3      = svg.Offcolor{100, "black", 1.0}
	ga       = []svg.Offcolor{oc1, oc2, oc3}
)

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
	defcoord(w/2, h/2)
	canvas.Rect(0, 0, w, h, objstyle)
	canvas.Text(-textsize, (h / 2), "h", legendstyle)
	canvas.Text((w / 2), -textsize, "w", legendstyle)
	deflegend((w / 2), 0, h, legend)
	deflegend(w/2, 0, h+(textsize)+(textsize/2), "CenterRect(x, y, w, h, ...style)")
	canvas.Gend()
}

func defcrect(id string, w int, h int, legend string) {
	canvas.Gid(id)
	defcoord(0, 0)
	defcoord(w/2, h/2)
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

func defqbezier(id string, ex int, ey int, legend string) {
	canvas.Gid(id)
	defcoordstr(0, 0, "sx, sy")
	defcoordstr(ex, -ey, "cx, cy")
	defcoordstr(ex, ey, "ex, ey")
	defcoordstr(ex*2, 0, "tx, ty")
	canvas.Qbezier(0, 0, ex, -ey, ex, ey, ex*2, 0, objstyle)
	deflegend(ex, ey, 40, legend)
	canvas.Gend()
}

func defqbez(id string, px int, py int, legend string) {
	canvas.Gid(id)
	defcoordstr(0, 0, "sx, sy")
	defcoordstr(px, py, "cx, cy")
	defcoordstr(px*2, 0, "ex, ey")
	canvas.Qbez(0, 0, px, py, px*2, 0, objstyle)
	deflegend(px, py, 10, legend)
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
	var x = []int{0, w / 3, (w * 3) / 4, w}
	var y = []int{0, -(h / 2), -(h / 3), -h}
	canvas.Gid(id)
	for i := 0; i < len(x); i++ {
		defcoord(x[i], y[i])
	}
	canvas.Polyline(x, y, objstyle)
	deflegend(x[1], y[1], 30, legend)
	canvas.Gend()
}

func defpath(id string, x, y int, legend string) {
	var w3path = `M36,5l12,41l12-41h33v4l-13,21c30,10,2,69-21,28l7-2c15,27,33,-22,3,-19v-4l12-20h-15l-17,59h-1l-13-42l-12,42h-1l-20-67h9l12,41l8-28l-4-13h9`
	var cpath = `M94,53c15,32,30,14,35,7l-1-7c-16,26-32,3-34,0M122,16c-10-21-34,0-21,30c-5-30 16,-38 23,-21l5-10l-2-9`
	canvas.Gid(id)
	canvas.Path(w3path, `fill="#AA0000"`)
	canvas.Path(cpath)
	defcoord(0, 0)
	deflegend(x/2, y, 10, legend)
	canvas.Gend()
}

func deflg(id string, w int, h int, legend string) {
	canvas.Gid(id)
	defcoordstr(0, 0, "x1%, y1%")
	defcoordstr(w, 0, "x2%, y2%")
	canvas.Rect(0, 0, w, h, "fill:url(#linear)")
	deflegend((w / 2), 0, h, legend)
	canvas.Gend()
}

func defrg(id string, w int, h int, legend string) {
	canvas.Gid(id)
	defcoordstr(0, 0, "cx%, cy%")
	canvas.Rect(0, 0, w, h, "fill:url(#radial)")
	defcoordstr(w/2, h/2, "fx%, fy%")
	deflegend((w / 2), 0, h, legend)
	canvas.Gend()
}

func deftrans(id string, w, h int, legend string) {
	tx := w / 3
	canvas.Gid(id)
	defcoordstr(0, 0, "0, 0")
	defcoordstr(w-tx, 0, "x, y")
	deflegend(w/2, 0, h, legend)
	canvas.Rect(0, 0, tx, h, objstyle)
	canvas.Translate(w-tx, 0)
	canvas.Rect(0, 0, tx, h, fobjstyle)
	canvas.Gend()
	canvas.Gend()
}

func defgrid(id string, w, h int, legend string) {
	canvas.Gid(id)
	canvas.Grid(0, 0, w, h, 25, "stroke:rgb(127,0,0)")
	deflegend((w / 2), 0, h, legend)
	canvas.Gend()
}

func deftext(id string, w, h int, text string, legend string) {
	canvas.Gid(id)
	canvas.Text(w/2, h, "hello", "font-size:48pt")
	deflegend(w/2, 0, h, legend)
	canvas.Gend()
}

func defscale(id string, w, h int, n float64, legend string) {
	canvas.Gid(id)
	canvas.Rect(0, 0, w, h, objstyle)
	canvas.Scale(n)
	canvas.Rect(0, 0, w, h, fobjstyle)
	canvas.Gend()
	deflegend(w/2, 0, h, legend)
	canvas.Gend()
}

func defrotate(id string, w, h int, n float64, legend string) {
	canvas.Gid(id)
	deflegend(w/2, 0, h, legend)
	canvas.Rect(0, 0, w, h, fobjstyle)
	canvas.Rotate(n)
	canvas.Rect(0, 0, w, h, objstyle)
	canvas.Gend()
	canvas.Gend()
}

func deftr(id string, tx, ty int, r float64, legend string) {
	canvas.Gid(id)
	canvas.Gend()
}

func defrt(id string, tx, ty int, r float64, legend string) {
	canvas.Gid(id)
	canvas.Gend()
}

func main() {
	width := 1400
	height := 1400
	canvas.Start(width, height)
	canvas.Title("SVG Go Library Description")
	canvas.Def()
	canvas.LinearGradient("linear", 0, 0, 100, 0, ga)
	canvas.RadialGradient("radial", 0, 0, 100, 50, 50, ga)
	canvas.DefEnd()
	canvas.Rect(0, 0, width, height, "fill:white;stroke:black;stroke-width:2")
	canvas.Gstyle(gtextstyle)
	canvas.Text(width/2, 100, "SVG Go Library", "font-size:50px")

	canvas.Desc("Object Definitions")
	canvas.Def()
	defsquare("square", 100, "Square(x, y, w,...style)")
	defrect("rectangle", 150, 100, "Rect(x, y, w, h,...style)")
	defcrect("crect", 150, 100, "CenterRect(x, y, w, h,...style)")
	defroundrect("roundrect", 200, 100, 25, 25, "Roundrect(x,y,w,h,rx,ry, ...style)")
	defpolygon("polygon", 200, 120, "Polygon(x, y, ...style)")

	defcircle("circle", 50, "Circle(x, y, r,...style)")
	defellipse("ellipse", 75, 50, "Ellipse(x, y, rx ,ry,...style)")
	defline("line", 200, "Line(x1, y1, x2, y2, ...style)")
	defpolyline("polyline", 200, 50, "Polyline(x, y, ...style)")

	defarc("arc", 50, 40, "Arc(sx, sy, ax, ay, r, lflag, sflag, ex, ey, ...style)")
	defpath("path", 100, 80, "Path(s, ...style)")
	defqbez("qbez", 100, 80, "Qbez(sx, sy, cx, cy, ex, ey, ...style)")
	defqbezier("qbezier", 100, 50, "Qbezier(sx, sy, cx, cy, ex, ey, tx, ty, ...style)")

	defbez("bezier", 100, 50, "Bezier(sx, sy, cx, cy, px, py, ex, ey, ...style)")
	defimage("image", 128, 128, "images/gophercolor128x128.png", "Image(x, y, w, h, path, ...style)")
	deflg("lgrad", 200, 100, "LinearGradient(id, x1, y1, x2, y2, Offcolor)")
	defrg("rgrad", 200, 100, "RadialGradient(id, cx, cy, r, fx, fy, Offcolor)")

	deftrans("trans", 200, 100, "Translate(x, y)")
	defgrid("grid", 200, 100, "Grid(x, y, w, h, n, ...style)")
	deftext("text", 200, 100, "hello", "Text(x, y, s, ...style)")
	defscale("scale", 200, 100, 0.5, "Scale(n)")
	defrotate("rotate", 200, 100, 30, "Rotate(n)")
	canvas.DefEnd()

	canvas.Desc("Object Usage")

	left := 140
	y := 200
	rowsize := y + 50
	
	// row 0
	canvas.Translate(left, y)
	canvas.Use(0, 0,    "#square")
	canvas.Use(300, 0,  "#rectangle")
	canvas.Use(600, 0,  "#roundrect")
	canvas.Use(950, 0,  "#polygon")
	canvas.Gend()
	y += rowsize

	// row 1
	canvas.Translate(left, y)
	canvas.Use(50, 0,   "#circle")
	canvas.Use(375, 0,  "#ellipse")
	canvas.Use(600, 0,  "#line")
	canvas.Use(950, 0,  "#polyline")
	canvas.Gend()
	y += rowsize

	// row 2
	canvas.Translate(left, y)
	canvas.Use(0, 0,    "#arc")
	canvas.Use(300, 0,  "#path")
	canvas.Use(600, 0,  "#qbez")
	canvas.Use(950, 0,  "#qbezier")
	canvas.Gend()
	y += rowsize

	// row 3
	canvas.Translate(left, y)
	canvas.Use(0, 0,    "#bezier")
	canvas.Use(300, 0,  "#image")
	canvas.Use(600, 0,  "#lgrad")
	canvas.Use(950, 0, "#rgrad")
	canvas.Gend()
	y += rowsize

	// row 4 
	canvas.Translate(left, y)
	canvas.Use(0, 0,    "#trans")
	canvas.Use(300, 0,  "#scale")
	canvas.Use(600, 0,  "#rotate")
	canvas.Use(950, 0,  "#grid")
	canvas.Gend()

	canvas.Gend()
	canvas.End()

}
