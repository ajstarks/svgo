// codepic -- produce code+output sample suitable for slides

package main

import (
	"bufio"
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/ajstarks/svgo"
	"io"
	"os"
	"strings"
)

var (
	canvas                                                    = svg.New(os.Stdout)
	font                                                      string
	codeframe, picframe                                       bool
	linespacing, fontsize, top, left, boxwidth, width, height int
	framestyle                                                = "stroke:gray;stroke-dasharray:1,1;fill:none"
	labelstyle                                                = "text-anchor:middle"
	codefmt                                                   = "font-family:%s;font-size:%dpx"
)

// incoming SVG file, capture everything into between <svg..> and </svg> 
// in the Doc string.  This code will be translated to form the "picture" portion
type SVG struct {
	Width  int    `xml:"attr"`
	Height int    `xml:"attr"`
	Doc    string `xml:"innerxml"`
}

// codepic makes a code+picture SVG file, given a go source file
// and conventionally named output -- given <name>.go, <name>.svg
func codepic(filename string) {
	var basename string

	bn := strings.Split(filename, ".")
	if len(bn) > 0 {
		basename = bn[0]
	} else {
		fmt.Fprintf(os.Stderr, "cannot get the basename for %s\n", filename)
		return
	}
	canvas.Start(width, height)
	canvas.Title(basename)
	canvas.Rect(0, 0, width, height, "fill:white")
	placepic(width/2, top, basename)
	canvas.Gstyle(fmt.Sprintf(codefmt, font, fontsize))
	placecode(left+fontsize, top+fontsize*2, filename)
	canvas.Gend()
	canvas.End()
}

// placecode places the code section on the left
func placecode(x, y int, filename string) {
	var rerr error
	var line string
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	in := bufio.NewReader(f)
	for xp := left + fontsize; rerr == nil; y += linespacing {
		line, rerr = in.ReadString('\n')
		if len(line) > 0 {
			canvas.Text(xp, y, line[0:len(line)-1], `xml:space="preserve"`)
		}
	}
	if codeframe {
		canvas.Rect(top, left, left+boxwidth, y, framestyle)
	}
}

// placepic places the picture on the right
func placepic(x, y int, basename string) {
	var s SVG
	f, err := os.Open(basename + ".svg")
	defer f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	if err := xml.Unmarshal(f, &s); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse (%v)\n", err)
		return
	}
	canvas.Text(x+s.Width/2, height-10, basename+".go", fmt.Sprintf(codefmt, font, fontsize*2))
	canvas.Group(`clip-path="url(#pic)"`, fmt.Sprintf(`transform="translate(%d,%d)"`, x, y))
	canvas.ClipPath(`id="pic"`)
	canvas.Rect(0, 0, s.Width, s.Height)
	canvas.ClipEnd()
	io.WriteString(canvas.Writer, s.Doc)
	canvas.Gend()
	if picframe {
		canvas.Rect(x, y, s.Width, s.Height, framestyle)
	}
}

// init initializes flags
func init() {
	flag.BoolVar(&codeframe, "codeframe", true, "frame the code")
	flag.BoolVar(&picframe, "picframe", true, "frame the picture")
	flag.IntVar(&width, "w", 1024, "width")
	flag.IntVar(&height, "h", 768, "height")
	flag.IntVar(&linespacing, "ls", 16, "linespacing")
	flag.IntVar(&fontsize, "fs", 14, "fontsize")
	flag.IntVar(&top, "top", 20, "top")
	flag.IntVar(&left, "left", 20, "left")
	flag.IntVar(&boxwidth, "boxwidth", 450, "boxwidth")
	flag.StringVar(&font, "font", "Inconsolata", "font name")
	flag.Parse()
}

// for every file, make a code+pic SVG file
func main() {
	for _, f := range flag.Args() {
		codepic(f)
	}
}
