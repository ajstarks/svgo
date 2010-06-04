package main

import (
	"svg"
	"os"
	"strconv"
)

var g = svg.New(os.Stdout)


func main() {
	width := 500
	height := 500
	oc1 := svg.Offcolor{0, "rgb(255,255,0)", 1.0}
	oc2 := svg.Offcolor{100, "rgb(255,0,0)", .5}
	oc3 := svg.Offcolor{0, "rgb(200,200,200)", 0.0}
	oc4 := svg.Offcolor{100, "rgb(0,0,255)", 1.0}
	oc5 := svg.Offcolor{10, "#00cc00", 1}
	oc6 := svg.Offcolor{30, "#006600", 1}
	oc7 := svg.Offcolor{70, "#cc0000", 1}
	oc8 := svg.Offcolor{90, "#000099", 1}

	lg := []svg.Offcolor{oc1, oc2}
	rg := []svg.Offcolor{oc3, oc4}
	rainbow := []svg.Offcolor{oc5, oc6, oc7, oc8}

	g.Start(width, height)
	g.Title("Gradients")

	g.Def()
	g.LinearGradient("h", 0, 100, 0, 0, lg)
	g.LinearGradient("v", 0, 0, 100, 0, lg)
	g.LinearGradient("rainbow", 0, 0, 100, 0, rainbow)
	g.RadialGradient("rad100", 50, 50, 100, 20, 50, rg)
	g.RadialGradient("rad50", 50, 50, 50, 20, 50, rg)
	for i := 50; i < 100; i += 10 {
		g.RadialGradient("grad"+strconv.Itoa(i), 50, 50, uint8(i), 20, 50, rg)
	}
	g.DefEnd()

	g.Ellipse(width/2, height/2, 100, 100, "fill:url(#rad100)")
	g.Rect(300, 200, 100, 100, "fill:url(#h)")
	g.Rect(100, 200, 100, 100, "fill:url(#v)")
	g.Roundrect(10, 10, width-20, 50, 10, 10, "fill:url(#rainbow)")

	for i := 50; i < 100; i += 10 {
		g.Circle(i*5, 100, 15, "fill:url(#grad"+strconv.Itoa(i)+")")
	}
	g.End()
}
