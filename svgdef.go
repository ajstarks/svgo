// svgdef: SVG Object Definition and Use

package main

import (
	"github.com/ajstarks/svgo"
	"os"
	"math"
)

const (
	textsize    = 20
	coordsize   = 4
	objstyle    = "fill:none; stroke-width:2; stroke:rgb(127,0,0)"
	fobjstyle   = "fill:rgb(127,0,0);fill-opacity:0.25"
	legendstyle = "fill:gray; text-anchor:middle"
	titlestyle  = "fill:black; text-anchor:middle;font-size:20px"
	linestyle   = "stroke:black; stroke-width:1"
	gtextstyle  = "font-family:Calibri; text-anchor:middle; font-size:18px"
	coordstring = "x, y"
)

var (
	canvas   = svg.New(os.Stdout)
	grayfill = canvas.RGB(220, 220, 220)
	oc1      = svg.Offcolor{0, "red", 1.0}
	oc2      = svg.Offcolor{50, "gray", 1.0}
	oc3      = svg.Offcolor{100, "black", 1.0}
	ga       = []svg.Offcolor{oc1, oc2, oc3}
)

func defcoordstr(x int, y int, s string) {
	canvas.Circle(x, y, coordsize, grayfill)
	canvas.Text(x, y-textsize, s, legendstyle)
}

func defcoord(x, y, n int) {
	canvas.Circle(x, y, coordsize, grayfill)
	canvas.Text(x, y+n, coordstring, legendstyle)
}

func deflegend(x int, y int, size int, legend string) {
	canvas.Text(x, y+size+textsize, legend, titlestyle)
}

func defcircle(id string, w, h int, legend string) {
	canvas.Gid(id)
	canvas.Translate(w, h)
	defcoord(0, 0, -textsize)
	canvas.Circle(0, 0, h, objstyle)
	canvas.Line(0, 0, h, 0, linestyle)
	canvas.Text(h/2, textsize, "r", legendstyle)
	deflegend(0, 0, h, legend)
	canvas.Gend()
	canvas.Gend()
}

func defellipse(id string, w int, h int, legend string) {
	canvas.Gid(id)
	canvas.Translate(w, h)
	defcoord(0, 0, -textsize)
	canvas.Ellipse(0, 0, w, h, objstyle)
	canvas.Line(0, 0, w, 0, linestyle)
	canvas.Line(0, 0, 0, h, linestyle)
	canvas.Text(w/2, textsize, "rx", legendstyle)
	canvas.Text(-textsize, (h / 2), "ry", legendstyle)
	deflegend(0, 0, h, legend)
	canvas.Gend()
	canvas.Gend()
}

func defrect(id string, w int, h int, legend string) {
	canvas.Gid(id)
	defcoord(0, 0, -textsize)
	canvas.Rect(0, 0, w, h, objstyle)
	canvas.Text(-textsize, (h / 2), "h", legendstyle)
	canvas.Text((w / 2), -textsize, "w", legendstyle)
	deflegend((w / 2), 0, h, legend)
	canvas.Gend()
}

func defcrect(id string, w int, h int, legend string) {
	canvas.Gid(id)
	defcoord(w/2, h/2, -textsize)
	canvas.Rect(0, 0, w, h, objstyle)
	canvas.Text(-textsize, (h / 2), "h", legendstyle)
	canvas.Text((w / 2), -textsize, "w", legendstyle)
	deflegend((w / 2), 0, h, legend)
	canvas.Gend()
}

func defsquare(id string, w int, legend string) {
	canvas.Gid(id)
	defcoord(0, 0, -textsize)
	canvas.Square(0, 0, w, objstyle)
	canvas.Text((w / 2), -textsize, "w", legendstyle)
	deflegend((w / 2), 0, w, legend)
	canvas.Gend()
}

func defimage(id string, w int, h int, s string, legend string) {
	canvas.Gid(id)
	defcoord(0, 0, -textsize)
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
	canvas.Arc(0, 0, h, h, 0, false, true, w*2, 0, objstyle)
	deflegend(w, h, 0, legend)
	canvas.Gend()
}

func defbez(id string, x int, y int, legend string) {
	sx, sy := 0, 0
	cx, cy := x, -y
	px, py := x, y
	ex, ey := x*2, 0
	canvas.Gid(id)
	defcoordstr(sx, sy, "sx, sy")
	defcoordstr(cx, cy, "cx, cy")
	defcoordstr(px, py, "px, py")
	defcoordstr(ex, ey, "ex, ey")
	canvas.Bezier(sx, sy, cx, cy, px, py, ex, ey, objstyle)
	deflegend(px, py, 0, legend)
	canvas.Gend()
}

func defqbezier(id string, x int, y int, legend string) {
	sx, sy := 0, 0
	cx, cy := x, -y
	px, py := x, y
	tx, ty := x*2, 0
	canvas.Gid(id)
	defcoordstr(sx, sy, "sx, sy")
	defcoordstr(cx, cy, "cx, cy")
	defcoordstr(px, py, "px, py")
	defcoordstr(tx, ty, "tx, ty")
	canvas.Qbezier(sx, sy, cx, cy, px, py, tx, ty, objstyle)
	deflegend(x, y, 40, legend)
	canvas.Gend()
}

func defqbez(id string, px int, py int, legend string) {
	sx, sy := 0, 0
	ex, ey := px*2, 0
	cx, cy := (ex-px)/3, -py-(py/2)
	canvas.Gid(id)
	defcoordstr(sx, sy, "sx, sy")
	defcoordstr(cx, cy, "cx, cy")
	defcoordstr(ex, ey, "ex, ey")
	canvas.Qbez(sx, sy, cx, cy, ex, ey, objstyle)
	deflegend(px, py, 0, legend)
	canvas.Gend()
}

func defroundrect(id string, w int, h int, rx int, ry int, legend string) {
	canvas.Gid(id)
	defcoord(0, 0, -textsize)
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
	var y = []int{0, -h / 4, 0, (h * 3) / 4, h / 2, (h * 3) / 4}
	canvas.Gid(id)
	for i := 0; i < len(x); i++ {
		defcoord(x[i], y[i], -textsize)
	}
	canvas.Polygon(x, y, objstyle)
	deflegend(w/2, h, 0, legend)
	canvas.Gend()
}

func defpolyline(id string, w int, h int, legend string) {
	var x = []int{0, w / 3, (w * 3) / 4, w}
	var y = []int{0, -(h / 2), -(h / 3), -h}
	canvas.Gid(id)
	for i := 0; i < len(x); i++ {
		defcoord(x[i], y[i], -textsize)
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
	defcoord(0, 0, -textsize)
	deflegend(x/2, y, textsize, legend)
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
	canvas.Rect(0, 0, tx, h, fobjstyle)
	canvas.Translate(w-tx, 0)
	canvas.Rect(0, 0, tx, h, objstyle)
	canvas.Gend()
	canvas.Gend()
}

func defgrid(id string, w, h int, legend string) {
	n := h/4
	canvas.Gid(id)
	defcoord(0, 0, -textsize)
	canvas.Text(-textsize, (h / 2), "h", legendstyle)
	canvas.Text((w / 2), -textsize, "w", legendstyle)
	canvas.Text(n+textsize, n/2, "n", legendstyle)
	canvas.Grid(0, 0, w, h, n, "stroke:rgb(127,0,0)")
	deflegend((w / 2), 0, h, legend)
	canvas.Gend()
}

func deftext(id string, w, h int, text string, legend string) {
	canvas.Gid(id)
	defcoord(0, h/2, textsize)
	canvas.Text(0, h/2, text, "text-anchor:start;font-size:32pt")
	deflegend(w/2, 0, h, legend)
	canvas.Gend()
}

func defscale(id string, w, h int, n float64, legend string) {
	canvas.Gid(id)
	defcoordstr(0, 0, "0, 0")
	canvas.Rect(0, 0, w, h, objstyle)
	canvas.Scale(n)
	canvas.Rect(0, 0, w, h, fobjstyle)
	canvas.Gend()
	deflegend(w/2, 0, h, legend)
	canvas.Gend()
}

func defrotate(id string, w, h int, deg float64, legend string) {
    t := deg * (math.Pi/180.0)
    r := float64(w/2)
    rx := r * math.Cos(t)
    ry := r * math.Sin(t)
	canvas.Gid(id)
	defcoordstr(0, 0, "0, 0")
	deflegend(w/2, 0, h, legend)
	canvas.Rect(0, 0, w, h, fobjstyle)
	//canvas.Arc(w/2, 0, 5, 5, 0, false, true, int(rx), int(ry), "fill:none;stroke:gray")
	canvas.Qbez(w/2, 0, (w/2)+10, int(ry)/2, int(rx), int(ry), "fill:none;stroke:gray")
	canvas.Text(w/4, textsize, "n", legendstyle)
	canvas.Rotate(deg)
	canvas.Rect(0, 0, w, h, objstyle)
	canvas.Gend()
	canvas.Gend()
}

func defobjects(w, h int) {
	h2 := h / 2
	canvas.Desc("Object Definitions")
	canvas.Def()
	canvas.LinearGradient("linear", 0, 0, 100, 0, ga)
	canvas.RadialGradient("radial", 0, 0, 100, 50, 50, ga)
	defsquare("square", 100, "Square(x, y, w, ...style)")
	defrect("rect", w, h, "Rect(x, y, w, h, ...style)")
	defcrect("crect", w, h, "CenterRect(x, y, w, h, ...style)")
	defroundrect("roundrect", w, h, 25, 25, "Roundrect(x, y, w, h, rx, ry, ...style)")
	defpolygon("polygon", w, h, "Polygon(x, y, ...style)")
	defcircle("circle", h, h2, "Circle(x, y, r, ...style)")
	defellipse("ellipse", h, h2, "Ellipse(x, y, rx, ry, ...style)")
	defline("line", w, "Line(x1, y1, x2, y2, ...style)")
	defpolyline("polyline", w, h2, "Polyline(x, y, ...style)")
	defarc("arc", h, h2, "Arc(sx, sy, ax, ay, r, lflag, sflag, ex, ey, ...style)")
	defpath("path", h, h2, "Path(s, ...style)")
	defqbez("qbez", h, h2, "Qbez(sx, sy, cx, cy, ex, ey, ...style)")
	defqbezier("qbezier", h, h2, "Qbezier(sx, sy, cx, cy, ex, ey, tx, ty, ...style)")
	defbez("bezier", h, h2, "Bezier(sx, sy, cx, cy, px, py, ex, ey, ...style)")
	defimage("image", 128, 128, "images/gophercolor128x128.png", "Image(x, y, w, h, path, ...style)")
	deflg("lgrad", w, h, "LinearGradient(id, x1, y1, x2, y2, oc)")
	defrg("rgrad", w, h, "RadialGradient(id, cx, cy, r, fx, fy, oc)")
	deftrans("trans", w, h, "Translate(x, y)")
	defgrid("grid", w, h, "Grid(x, y, w, h, n, ...style)")
	deftext("text", w, h, "hello, SVG", "Text(x, y, s, ...style)")
	defscale("scale", w, h, 0.5, "Scale(n)")
	defrotate("rotate", w, h, 30, "Rotate(n)")
	canvas.DefEnd()
}

func placerow(w int, s []string) {
	for x, name := range s {
		canvas.Use(x*w, 0, "#"+name)
	}
}

func placeobjects(x, y, w, h int, data [][]string) {
	canvas.Desc("Object Usage")
	for _, object := range data {
		canvas.Translate(x, y)
		placerow(w, object)
		canvas.Gend()
		y += h
	}
}


var roworder = [][]string{
	{"rect", "crect", "roundrect", "square", "line", "polyline"},
	{"polygon", "circle", "ellipse", "arc", "qbez", "bezier"},
	{"trans", "scale", "rotate", "grid", "lgrad", "rgrad"},
	{"text", "image", "path"},
}

func main() {
	width := 2200
	height := (width * 3) / 4
	canvas.Start(width, height)
	defobjects(250, 125)
	canvas.Title("SVG Go Library Description")
	canvas.Rect(0, 0, width, height, "fill:white;stroke:black;stroke-width:2")
	canvas.Gstyle(gtextstyle)
	canvas.Text(width/2, 75, "SVG Go Library", "font-size:50px")
	placeobjects(120, 200, 350, 400, roworder)
	canvas.Gend()
	canvas.End()

}
