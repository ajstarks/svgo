package main

import (
	"fmt"
	"github.com/swill/svgo"
	"os"
)

var (
	width  = 33.333
	height = 33.333
	style  = fmt.Sprintf("fill:none;stroke:black;stroke-width:0.2")
)

func main() {
	terminal := os.Stdout // just print the examples to the terminal.

	fmt.Fprintln(terminal, "\nDifferent doc starts")
	fmt.Fprintln(terminal, "--------------------")

	// StartF
	fmt.Fprintln(terminal, "\nStartF")
	a := svg.New(terminal)
	a.StartF(width, height)
	a.End()

	// StartunitF and SetFloatDecimals
	fmt.Fprintln(terminal, "\nStartunitF with optional FloatDecimals = 1")
	b := svg.New(terminal)
	b.FloatDecimals = 1
	b.StartunitF(width, height, "mm")
	b.End()

	// StartviewF and SetFloatDecimals
	fmt.Fprintln(terminal, "\nStartviewF with optional FloatDecimals = 3")
	c := svg.New(terminal)
	c.FloatDecimals = 0
	c.StartviewF(width, height, 0, 0, width, height)
	c.End()

	// StartviewUnitF
	fmt.Fprintln(terminal, "\nStartviewUnitF default FloatDecimals")
	d := svg.New(terminal)
	d.StartviewUnitF(width, height, "mm", 0, 0, width, height)
	d.End()

	// StartpercentF
	fmt.Fprintln(terminal, "\nStartpercentF with optional FloatDecimals = 0")
	e := svg.New(terminal)
	e.FloatDecimals = 0
	e.StartpercentF(width, height)
	e.End()

	// -- -- --

	// Create an actual SVG file that can be viewed in a browser and verified...
	file, err := os.Create("float_examples.svg")
	if err != nil {
		panic("ERROR: Unable to create the 'float_example.svg' file...")
	}
	defer file.Close()

	// setup
	out := svg.New(terminal)
	out.FloatDecimals = 0                                // the number of decimals our floats will have
	out.StartviewUnitF(1024, 768, "mm", 0, 0, 1024, 768) // setup the view
	// squares in top left
	out.CenterRectF(15, 15, 20, 20, style)
	out.RoundrectF(7, 7, 16, 16, 2, 2, style)
	out.SquareF(9, 9, 12, style)
	out.FloatDecimals = 2
	out.PolygonF([]float64{10, 20, 15, 10}, []float64{20, 20, 10, 20}, style) // triangle
	out.CircleF(15, 15, 1.5, style)
	out.EllipseF(15, 18.25, 3.5, 1, style)
	// translated ellipses right of squares
	out.TranslateF(30, 12.5)       // translate to the right of the boxes
	out.TranslateRotateF(5, 0, 45) // translate then rotate
	out.EllipseF(0, 0, 2, 10, fmt.Sprintf("fill:red;fill-opacity:0.5;stroke:black;stroke-width:0.2"))
	out.Gend()
	out.RotateTranslateF(5, 0, 45) // rotate then translate
	out.EllipseF(0, 0, 2, 10, fmt.Sprintf("fill:blue;fill-opacity:0.5;stroke:black;stroke-width:0.2"))
	out.Gend()
	out.Gend()
	// rect below squares and lines
	out.RectF(5, 27.5, 20, 5, style)
	out.LineF(25, 26.25, 35, 26.25, fmt.Sprintf("stroke:black;stroke-width:0.2"))
	out.Def()
	out.MarkerF("dot", 5, 5, 8, 8)
	out.CircleF(5, 5, 3, "fill:black")
	out.MarkerEnd()
	out.DefEnd()
	out.PolylineF(
		[]float64{27.5, 37.5, 37.5},
		[]float64{30, 30, 22.5},
		`fill="none"`,
		`stroke="black"`,
		`stroke-width="0.2"`,
		`marker-end="url(#dot)"`)
	out.End()
	fmt.Fprintln(terminal, "\n---  ---  ---")
	fmt.Fprintln(terminal, "\nSaved: float_examples.svg")
	fmt.Fprintln(terminal, "\nYou can review the elements of this SVG in your browser with 'Inspect Element'.")
	fmt.Fprintln(terminal, "Creation code is at the bottom of 'float_examples.go'.")
}
