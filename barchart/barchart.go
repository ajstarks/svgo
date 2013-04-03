// barchart - draw bar charts
package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/ajstarks/svgo"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	width, height, iscale, fontsize, barheight, gutter, cornerRadius int
	bgcolor, barcolor, title, inbar                                  string
	showtitle, showdata, showgrid, showscale, endtitle               bool
)

const (
	gstyle      = "font-family:Calibri,sans-serif;font-size:%dpx"
	borderstyle = "stroke:lightgray;stroke-width:1px"
	scalestyle  = "text-anchor:middle;font-size:75%"
	btitlestyle = "font-style:italic;font-size:150%;text-anchor:"
	notestyle   = "font-style:italic;text-anchor:"
	datastyle   = "text-anchor:end;fill:"
	titlestyle  = "text-anchor:start;font-size:300%"
	labelstyle  = "fill:black;baseline-shift:-25%"
)

// a Barchart Defintion
// <barchart title="Bullet Graph" top="50" left="250" right="50">
//    <note>This is a note</note>
//    <note>More expository text</note>
//    <bdata title="Browser Market Share" scale="0,100,20" showdata="on" color="red" unit="%"/>
//    	<bitem name="Firefox"  value="22.5" color="green"/>
//    	<bitem name="Chrome" value="12.3"/>
//		<bitem name="IE8" value="63.5"/>
//	  <bdata>
// </barchart>

type Barchart struct {
	Top   int     `xml:"top,attr"`
	Left  int     `xml:"left,attr"`
	Right int     `xml:"right,attr"`
	Title string  `xml:"title,attr"`
	Bdata []bdata `xml:"bdata"`
}

type bdata struct {
	Title    string  `xml:"title,attr"`
	Scale    string  `xml:"scale,attr"`
	Color    string  `xml:"color,attr"`
	Unit     string  `xml:"unit,attr"`
	Showdata string  `xml:"showdata,attr"`
	Showgrid string  `xml:"showgrid,attr"`
	Bitem    []bitem `xml:"bitem"`
	Note     []note  `xml:"note"`
}

type bitem struct {
	Name  string  `xml:"name,attr"`
	Value float64 `xml:"value,attr"`
	Color string  `xml:"color,attr"`
}

type note struct {
	Text string `xml:",chardata"`
}

// dobc does file i/o
func dobc(location string, s *svg.SVG) {
	var f *os.File
	var err error
	if len(location) > 0 {
		f, err = os.Open(location)
	} else {
		f = os.Stdin
	}
	if err == nil {
		readbc(f, s)
		f.Close()
	} else {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}

// readbc reads and parses the XML specification
func readbc(r io.Reader, s *svg.SVG) {
	var bc Barchart
	if err := xml.NewDecoder(r).Decode(&bc); err == nil {
		drawbc(bc, s)
	} else {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}

// drawbc draws the bar chart
func drawbc(bg Barchart, canvas *svg.SVG) {

	if bg.Left == 0 {
		bg.Left = 250
	}
	if bg.Right == 0 {
		bg.Right = 50
	}
	if bg.Top == 0 {
		bg.Top = 50
	}
	if len(title) > 0 {
		bg.Title = title
	}

	cr := cornerRadius
	maxwidth := width - (bg.Left + bg.Right)
	x := bg.Left
	y := bg.Top
	sep := 10
	color := barcolor
	scfmt := "%v"
	canvas.Title(bg.Title)

	// for each bdata element...
	for _, b := range bg.Bdata {

		// overide the color if specified
		if len(b.Color) > 0 {
			color = b.Color
		} else {
			color = barcolor
		}
		// extract the scale data from the XML attributes
		// if not specified, compute the scale factors
		sc := strings.Split(b.Scale, ",")
		var scalemin, scalemax, scaleincr float64
		if len(sc) != 3 {
			scalemin, scalemax, scaleincr = scalevalues(b.Bitem)
		} else {
			scalemin, _ = strconv.ParseFloat(sc[0], 64)
			scalemax, _ = strconv.ParseFloat(sc[1], 64)
			scaleincr, _ = strconv.ParseFloat(sc[2], 64)
		}
		// label the graph
		canvas.Text(x, y, b.Title, btitlestyle+anchor())

		y += sep * 2
		chartop := y
		chartbot := chartop + ((len(b.Bitem)) * (barheight + gutter))

		// draw the scale and borders
		if showgrid || b.Showgrid == "on" {
			canvas.Line(x, chartop, x+maxwidth, y, borderstyle)                       // top border
			canvas.Line(x, chartbot-gutter, x+maxwidth, chartbot-gutter, borderstyle) // bottom border
		}
		if showscale {
			if scaleincr < 1 {
				scfmt = "%.1f"
			} else {
				scfmt = "%0.f"
			}
			canvas.Gstyle(scalestyle)
			for sc := scalemin; sc <= scalemax; sc += scaleincr {
				scx := vmap(sc, scalemin, scalemax, 0, float64(maxwidth))
				canvas.Text(x+int(scx), chartbot+fontsize, fmt.Sprintf(scfmt, sc))
				if showgrid || b.Showgrid == "on" {
					canvas.Line(x+int(scx), chartbot, x+int(scx), chartop, borderstyle) // grid line
				}
			}
			canvas.Gend()
		}

		// draw the data items
		canvas.Gstyle(datastyle + color)
		for _, d := range b.Bitem {
			canvas.Text(x-sep, y+barheight/2, d.Name, labelstyle)
			dw := vmap(d.Value, scalemin, scalemax, 0, float64(maxwidth))
			if len(d.Color) > 0 {
				canvas.Roundrect(x, y, int(dw), barheight, cr, cr, "fill:"+d.Color)
			} else {
				canvas.Roundrect(x, y, int(dw), barheight, cr, cr)
			}
			if showdata || b.Showdata == "on" {
				var valuestyle = "font-style:italic;font-size:75%;text-anchor:start;baseline-shift:-25%;"
				var ditem string
				var datax int
				if len(b.Unit) > 0 {
					ditem = fmt.Sprintf("%v%s", d.Value, b.Unit)
				} else {
					ditem = fmt.Sprintf("%v", d.Value)
				}
				if len(inbar) > 0 {
					valuestyle += inbar
					datax = x + fontsize/2
				} else {
					valuestyle += "fill:black"
					datax = x + int(dw) + fontsize/2
				}
				canvas.Text(datax, y+barheight/2, ditem, valuestyle)
			}
			y += barheight + gutter
		}
		canvas.Gend()

		// apply the note if present
		if len(b.Note) > 0 {
			canvas.Gstyle(notestyle + anchor())
			y += fontsize * 2
			leading := 3
			for _, note := range b.Note {
				canvas.Text(bg.Left, y, note.Text)
				y += fontsize + leading
			}
			canvas.Gend()
		}
		y += sep * 7 // advance vertically for the next chart
	}
	// if requested, place the title below the last chart
	if showtitle && len(bg.Title) > 0 {
		y += fontsize * 2
		canvas.Text(bg.Left, y, bg.Title, titlestyle)
	}
}

func anchor() string {
	if endtitle {
		return "end"
	}
	return "start"
}

// vmap maps one interval to another
func vmap(value float64, low1 float64, high1 float64, low2 float64, high2 float64) float64 {
	return low2 + (high2-low2)*(value-low1)/(high1-low1)
}

func maxitem(data []bitem) float64 {
	max := -math.SmallestNonzeroFloat64
	for _, d := range data {
		if d.Value > max {
			max = d.Value
		}
	}
	return max
}

func scalevalues(data []bitem) (float64, float64, float64) {
	var m, max, increment float64
	rui := 5
	m = maxitem(data)
	max = roundup(m, 100)
	if max > 2 {
		increment = roundup(max/float64(rui), 10)
	} else {
		increment = 0.4
	}
	return 0, max, increment
}

func roundup(n float64, m int) float64 {
	i := int(n)
	if i <= 2 {
		return 2
	}
	for ; i%m != 0; i++ {
	}
	return float64(i)
}

// init sets up the command flags
func init() {
	flag.StringVar(&bgcolor, "bg", "white", "background color")
	flag.StringVar(&barcolor, "bc", "rgb(200,200,200)", "bar color")
	flag.IntVar(&width, "w", 1024, "width")
	flag.IntVar(&height, "h", 800, "height")
	flag.IntVar(&barheight, "bh", 20, "bar height")
	flag.IntVar(&gutter, "g", 5, "gutter")
	flag.IntVar(&cornerRadius, "cr", 0, "corner radius")
	flag.IntVar(&fontsize, "f", 18, "fontsize (px)")
	flag.BoolVar(&showscale, "showscale", true, "show scale")
	flag.BoolVar(&showgrid, "showgrid", false, "show grid")
	flag.BoolVar(&showdata, "showdata", false, "show data values")
	flag.BoolVar(&showtitle, "showtitle", false, "show title")
	flag.BoolVar(&endtitle, "endtitle", false, "align title to the end")
	flag.StringVar(&inbar, "inbar", "", "data in bar format")
	flag.StringVar(&title, "t", "", "title")
	flag.Parse()
}

// for every input file (or stdin), draw a bar graph
// as specified by command flags
func main() {
	canvas := svg.New(os.Stdout)
	canvas.Start(width, height)
	canvas.Rect(0, 0, width, height, "fill:"+bgcolor)
	canvas.Gstyle(fmt.Sprintf(gstyle, fontsize))
	if len(flag.Args()) == 0 {
		dobc("", canvas)
	} else {
		for _, f := range flag.Args() {
			dobc(f, canvas)
		}
	}
	canvas.Gend()
	canvas.End()
}
