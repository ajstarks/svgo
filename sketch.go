package main

import (
	"./svg"
	"time"
	"rand"
	"fmt"
	"flag"
	"math"
)


type coord float
type dimension float

var width = 500
var height = 500
var niter = flag.Int("n", 200, "number of iterations")
var font = flag.String("f", "Calibri", "font name")
var fontsize = flag.Int("s", 72, "font size")
var ns = flag.Int("r", 30, "r")
var rstyle = flag.String("rs", "stroke:gray; stroke-opacity:0.3; stroke-width:3", "radial style")

func randletters() {
	fc := "fill:black"
	svg.Gstyle("fill-opacity:0.6; font-family:" + *font)
	c := 0
	s := 5
	x := 0
	y := 0
	for i := 0; i < *niter; {
		s = rand.Intn(*fontsize)
		if s < 5 {
			s = 5
		}
		if i%2 == 0 {
			fc = fmt.Sprint("fill:white; font-size:", s)
		} else {
			fc = fmt.Sprint("fill:rgb(127,0,0); font-size:", s)
		}
		c = rand.Intn(127)
		if c < 90 && c > 64 || c < 123 && c > 96 {
			x = rand.Intn(width)
			y = rand.Intn(height)
			svg.Text(x+1, y+1, string(c), fc+";fill:black")
			svg.Text(x, y, string(c), fc)
			i++
		}
	}
	svg.Gend()
}

func dotgrid(xb int, yb int, w int, h int, size int, dotsize int) {
	svg.Gstyle("fill:gray; fill-opacity:0.2")
	for y := yb; y <= h; y += size {
		for x := xb; x <= w; x += size {
			svg.Circle(x, y, dotsize, "")
		}
	}
	svg.Gend()
}

func ipad(x int, y int, h int) {
	bezel := 50
	b2 := bezel / 2
	aspect := 0.75
	w := float(h) * aspect
	svg.Roundrect(x, y, int(w), h, b2, b2, "fill:black")
	svg.Rect(x+b2, y+b2, int(w)-bezel, h-bezel, "fill:white")
	svg.Grid(x+b2, y+b2, int(w)-bezel, h-bezel, 10, "stroke-width:1; stroke-opacity:0.4; stroke:black")
}


func money() {
	svg.Rect(0, 0, width, height/3, "fill:green; fill-opacity:0.1")
	svg.Rect(0, height/3, width, height/3, "fill:green; fill-opacity:0.3")
	svg.Rect(0, 2*(height/3), width, height/3, "fill:red")
	svg.Line(0, 0, width, height, "stroke:black")
	svg.Circle(width/2, height/2, (height / 10), "fill:white; stroke:none")
	svg.Circle((width/2)-85, (height/2)-85, height/15, "fill:none; stroke:black")
}
/*
func archobj() {
	note(100, 100, 100, 100, 10, "fill:rgb(127,0,0); stroke:none")
	svg.Text(100+(100/2), 100+(100/2), "Server1", "text-anchor:middle; font-size:14; fill:white")
	folder(250, 100, 150, 100, 50, 10, svg.RGB(130, 42, 132))
	svg.Text(250+(150/2), 100+(100)/2, "UTIL03", "text-anchor:middle;font-size:18; fill:white")

	rfolder(250, 250, 150, 100, 50, 10, svg.RGB(130, 42, 132))
	rnote(100, 250, 100, 100, 10, "fill:rgb(127,0,0); stroke:none")
	svg.Gstyle("font-family:Calibri; font-size:18; fill:white")
	rnote(100, 360, 100, 100, 10, "fill:rgb(127,0,0); stroke:none", "hello")
	svg.Gend()
}
*/
func useit(s string) {
	svg.Def()
	svg.Gid("circle")
	svg.Circle(0, 0, 100)
	svg.Gend()
	svg.DefEnd()
	svg.Use(100, 100, "#circle")
	svg.Use(300, 100, "#circle", s)
}

func randshapes(w, h, n int) {
	for i := 0; i < n; i++ {
		x := rand.Intn(w)
		y := rand.Intn(h)
		r := rand.Intn(255)
		g := rand.Intn(255)
		b := rand.Intn(255)
		s1 := rand.Intn(100)
		s2 := rand.Intn(100)
		if i%2 == 0 {
			svg.Ellipse(x, y, s1/2, s2/2, svg.RGB(r, g, b))
		} else {
			svg.Rect(x, y, s1, s2, svg.RGB(r, g, b))
		}
	}
}

func smile(x, y, r int, style ...string) {

	if len(style) > 0 {
		svg.Gstyle(style[0])
	}
	svg.Roundrect(x-(r*2), y-(r*2), r*7, r*20, r*2, r*2, svg.RGB(200, 200, 200))
	svg.Circle(x, y, r, svg.RGB(127, 0, 0))
	svg.Link("planets.svg", "Planets")
	svg.Circle(x, y, r/4, "fill:white")
	svg.LinkEnd()
	svg.Circle(x+(r*3), y, r)
	svg.Arc(x-r, y+(r*3), r/4, r/4, 0, true, false, x+(r*4), y+(r*3))
	if len(style) > 0 {
		svg.Gend()
	}
}

func hotpotato() {

	smile(200, 100, 10)
	svg.Gtransform("rotate(30)")
	svg.Link("http://hotpotato.com/", "Hot Potato Site")
	smile(200, 100, 10)
	svg.LinkEnd()
	svg.Gend()

	svg.Gtransform("translate(50,0) scale(2,2)")
	smile(200, 100, 30, "opacity:0.3")
	svg.Gend()

}

func progress(x, y, w, h, percentage int) {
	pct := float(percentage) / 100.0
	loc := x + int(float(w)*pct)
	svg.Roundrect(x, y, w+h, h, h/2, h/2, "fill:gray")
	svg.Circle(loc, y+(h/2), h/2)
}

func background(v int) { svg.Rect(0, 0, width, height, svg.RGB(v, v, v)) }

func init() {
	flag.Parse()
	rand.Seed(time.Nanoseconds() % 1e9)
}


func lewitt139() {
	svg.Gstyle("stroke:black; fill:none")
	svg.Arc(0, 0, 20, 20, 0, false, true, 0, 240, "opacity:0.5; fill:red")
	svg.Arc(320, 0, 20, 20, 0, true, false, 320, 240, "opacity:0.5; fill:blue")
	svg.Arc(0, 240, 10, 10, 0, false, true, 320, 240, "opacity:0.5; fill:green")
	svg.Arc(0, 0, 10, 10, 0, false, false, 320, 0, "")
	svg.Gend()
	svg.Grid(0, 0, 320, 240, 20, "stroke:gray; stroke-opacity:0.2")
}

func roundtop(x, y, w, h, r int, s string) {
	svg.Roundrect(x, y, w, h, r, r, s)
	svg.Rect(x, y+r, w, h-r, s+";stroke:none")
}


func ngrid(x, y, w, h, n, m int) {
	svg.Grid(x, y, w, h, n, "stroke:rgb(210,210,210)")
	fs := 12
	svg.Gstyle(fmt.Sprintf("font-size:%d; font-family:monospace", fs))
	for ix := x; ix <= w; ix += m {
		svg.Text(ix, (h / 2), fmt.Sprintf("%d", ix), "fill:red;text-anchor:end")
	}
	for iy := y; iy <= h; iy += m {
		svg.Text((w / 2), iy, fmt.Sprintf("%d", iy), "fill:blue;text-anchor:start")
	}
	svg.Gend()
}


func radial(xp int, yp int, n int, l int, style ...string) {
	var x float64
	var y float64
	var r float64
	var t float64
	var limit float64
	limit = 2.0 * math.Pi

	r = float64(l)
	svg.Gstyle(style[0])
	for t = 0.0; t < limit; t += limit / float64(n) {
		x = r * math.Cos(t)
		y = r * math.Sin(t)
		svg.Line(xp, yp, xp+int(x), yp+int(y))
	}
	svg.Gend()
}

func random(howsmall, howbig int) int {
	if howsmall >= howbig {
		return howsmall
	}
	return rand.Intn(howbig-howsmall) + howsmall
}

func randrad(w int, h int, n int) {
	for i := 0; i < n; i++ {
		x := rand.Intn(w)
		y := rand.Intn(h)
		r := rand.Intn(255)
		g := rand.Intn(255)
		b := rand.Intn(255)
		oi := random(10, 50)
		s := random(10, 60)
		t := random(2, 10)
		p := random(10, 15)
		radial(x, y, p, s,
			fmt.Sprintf("stroke:rgb(%d,%d,%d); stroke-opacity:%.2f; stroke-width:%d",
				r, g, b, float64(oi)/100.0, t))
	}
}



func diamond(x, y, w, h1, h2 int) {
  var xp = []int{x,     x+(w/2),    x,    x-(w/2)}
  var yp = []int{y-h1,  y,          y+h2, y}
  svg.Polygon(xp, yp)
}

func main() {

	svg.Start(width, height)
	background(0)
	
	
	x := 100
	y := 30
	//svg.Rect(x-40, y-25, 80, 500, "fill:gray")
	svg.Gstyle("fill:white")
	for i:=0; i < 10; i++ {
	  diamond(x, y, 20, 10, 10)
	  diamond(x, y+20, 20, 10, 10)
	  diamond(x+20, y+20, 20, 10, 10)
	  diamond(x-20, y+20, 20, 10, 10)
	  y += 40
	}
	svg.Gend()
	
	
	//randrad(width, height, *niter)
	//radial(width/2, height/2, *ns, width/2)
	//background(255)
	//archobj()
	//lewitt139()
	//hotpotato()
	//randshapes(500, 500, 500)
	//useit("fill:red")
	//archobj()
	//ipad(50, 50, 300)
	//ipad(300, 200, 200)
	//dotgrid(0,0,width,height,50, 5)
	//randletters()
	//money()
	svg.End()
}
