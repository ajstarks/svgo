// svgrid -- composite SVG files in a grid
package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/ajstarks/svgo"
)

// SVG is a SVG document
type SVG struct {
	Width  int    `xml:"width,attr"`
	Height int    `xml:"height,attr"`
	Doc    string `xml:",innerxml"`
}

var (
	byrow                                          bool
	startx, starty, count, gutter, gwidth, gheight int
	canvas                                         = svg.New(os.Stdout)
)

// init sets up command line options
func init() {
	flag.BoolVar(&byrow, "r", true, "order row wise")
	flag.IntVar(&startx, "x", 0, "begin x")
	flag.IntVar(&starty, "y", 0, "begin y")
	flag.IntVar(&count, "c", 3, "columns or rows")
	flag.IntVar(&gutter, "g", 100, "gutter")
	flag.IntVar(&gwidth, "w", 1024, "width")
	flag.IntVar(&gheight, "h", 768, "height")
	flag.Parse()
}

// placepic puts a SVG file at a location
func placepic(x, y int, filename string) (int, int) {
	var s SVG
	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return 0, 0
	}
	defer f.Close()
	if err := xml.NewDecoder(f).Decode(&s); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse (%v)\n", err)
		return 0, 0
	}
	canvas.Group(`clip-path="url(#pic)"`, fmt.Sprintf(`transform="translate(%d,%d)"`, x, y))
	canvas.ClipPath(`id="pic"`)
	canvas.Rect(0, 0, s.Width, s.Height)
	canvas.ClipEnd()
	io.WriteString(canvas.Writer, s.Doc)
	canvas.Gend()
	return s.Width, s.Height
}

// compose places files row or column-wise
func compose(x, y, n int, rflag bool, files []string) {
	px := x
	py := y
	var pw, ph int
	for i, f := range files {
		if i > 0 && i%n == 0 {
			if rflag {
				px = x
				py += gutter + ph
			} else {
				px += gutter + pw
				py = y
			}
		}
		pw, ph = placepic(px, py, f)
		if rflag {
			px += gutter + pw
		} else {
			py += gutter + ph
		}
	}
}

// main lays out files as specified on the command line
func main() {
	canvas.Start(gwidth, gheight)
	compose(startx, starty, count, byrow, flag.Args())
	canvas.End()
}
