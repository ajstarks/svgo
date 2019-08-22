// gradient shows sample gradient fills
// +build !appengine

package main

import (
	"os"
	"strconv"

	svg "github.com/ajstarks/svgo"
)

func main() {
	width := 500
	height := 500

	lg := []svg.Offcolor{
		{Offset: 0, Color: "rgb(255,255,0)", Opacity: 1.0},
		{Offset: 100, Color: "rgb(255,0,0)", Opacity: .5},
		{Offset: 0, Color: "rgb(200,200,200)", Opacity: 0.0},
		{Offset: 100, Color: "rgb(0,0,255)", Opacity: 1.0}}

	rainbow := []svg.Offcolor{
		{Offset: 10, Color: "#00cc00", Opacity: 1},
		{Offset: 30, Color: "#006600", Opacity: 1},
		{Offset: 70, Color: "#cc0000", Opacity: 1},
		{Offset: 90, Color: "#000099", Opacity: 1}}

	rg := []svg.Offcolor{
		{Offset: 1, Color: "powderblue", Opacity: 1},
		{Offset: 10, Color: "lightskyblue", Opacity: 1},
		{Offset: 100, Color: "darkblue", Opacity: 1}}

	g := svg.New(os.Stdout)
	g.Start(width, height)
	g.Title("Gradients")
	g.Rect(0, 0, width, height, "fill:white")
	g.Def()
	g.LinearGradient("h", 0, 100, 0, 0, lg)
	g.LinearGradient("v", 0, 0, 100, 0, lg)
	g.LinearGradient("rainbow", 0, 0, 100, 0, rainbow)
	g.RadialGradient("rad100", 50, 50, 100, 25, 25, rg)
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
