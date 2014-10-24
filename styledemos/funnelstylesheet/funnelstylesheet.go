// funnel draws a funnel-like shape
// +build !appengine

// Demo of SVG's Style entity to embed rules using StyleStart(), StyleEnd() to 
// enclose the CSS rules, and StyleRule() or StyleString() to insert the 
// CSS rules text
// Otherwise it is identical to the web-served funnel demo

package main


import (
	"io"

	"github.com/ajstarks/svgo"
)

const (
	root = "/funnel/"
	cssroot = root+"look/"
)

const width = 320
const height = 480

const rectRule = `rect`
const ellipseFmt = `ellipse { %s }`

func funnel(canvas *svg.SVG, bg int, fg int, grid int, dim int) {
	h := dim / 2
	canvas.Rect(0, 0, width, height)				// style removed to external style sheet 
	for size := grid; size < width; size += grid {
		canvas.Ellipse(h, size, size/2, size/2)		// style removed to external style sheet 
	}
}

func svgHandler(w io.Writer) {
	const bg = 0
	const fg = 255
	canvas := svg.New(w)
		// Use an external style sheet for CSS rules
	canvas.ExternalStyleSheet(cssroot+"style.css")

	canvas.Start(width, height)
	canvas.Title("Funnel")
	funnel(canvas, bg, fg, 25, width)
	canvas.End()
}

