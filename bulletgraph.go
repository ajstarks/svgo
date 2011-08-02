// bulletgraph - bullet graphs
// Bullet Graph Design Specification, Steven Few
// http://www.perceptualedge.com/articles/misc/Bullet_Graph_Design_Spec.pdf

package main

import (
	"github.com/ajstarks/svgo"
	"os"
	"fmt"
	"flag"
	"xml"
	"strconv"
	"strings"
	"io"
)

var (
	width, height, iscale, fontsize, barheight, gutter int
	bgcolor, barcolor, datacolor, compcolor, title     string
	showtitle, circlemark                              bool
	gstyle                                             = "font-family:Calibri;font-size:%dpx"
)

// a Bulletgraph Defintion
// <bulletgraph top="50" left="250" right="50">
//    <bdata title="Revenue 2005" subtitle="USD (1,000)" scale="0,300,50" qmeasure="150,225" cmeasure="250" measure="275"/>
//    <bdata title="Profit"  subtitle="%" scale="0,30,5" qmeasure="20,25" cmeasure="27" measure="22.5"/>
//    <bdata title="Avg Order Size subtitle="USD" scale="0,600,100" qmeasure="350,500" cmeasure="550" measure="320"/>
//    <bdata title="New Customers" subtitle="Count" scale="0,2500,500" qmeasure="1700,2000" cmeasure="2100" measure="1750"/>
//    <bdata title="Cust Satisfaction" subtitle="Top rating of 5" scale="0,5,1" qmeasure="3.5,4.5" cmeasure="4.7" measure="4.85"/>
// </bulletgraph>
type Bulletgraph struct {
	Top   string `xml:"attr"`
	Left  string `xml:"attr"`
	Right string `xml:"attr"`
	Bdata []bdata
}

type bdata struct {
	Title    string `xml:"attr"`
	Subtitle string `xml:"attr"`
	Scale    string `xml:"attr"`
	Qmeasure string `xml:"attr"`
	Cmeasure string `xml:"attr"`
	Measure  string `xml:"attr"`
}

// dobg does file i/o
func dobg(location string, s *svg.SVG) {
	var f *os.File
	var err os.Error
	if len(location) > 0 {
		f, err = os.Open(location)
	} else {
		f = os.Stdin
	}
	defer f.Close()
	if err == nil {
		readbg(f, s)
	} else {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}

// readbg reads and parses the XML specification
func readbg(r io.Reader, s *svg.SVG) {
	var bg Bulletgraph
	if err := xml.Unmarshal(r, &bg); err == nil {
		drawbg(bg, s)
	} else {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}

// drawbg draws the bullet graph
func drawbg(bg Bulletgraph, canvas *svg.SVG) {
	qmheight := barheight / 3
	top, _ := strconv.Atoi(bg.Top)
	left, _ := strconv.Atoi(bg.Left)
	right, _ := strconv.Atoi(bg.Right)
	maxwidth := width - (left + right)
	x := left
	y := top
	scalesep := 4
	tx := x - fontsize

	canvas.Title(title)
	// for each bdata element...
	for _, v := range bg.Bdata {

		// extract the data from the XML attributes
		sc := strings.Split(v.Scale, ",")
		qm := strings.Split(v.Qmeasure, ",")

		// you must have min,max,increment for the scale, at least one qualitative measure
		if len(sc) != 3 || len(qm) < 1 {
			println(len(sc), len(qm))
			continue
		}
		// get the qualitative measures
		qmeasures := make([]float64, len(qm))
		for i, q := range qm {
			qmeasures[i], _ = strconv.Atof64(q)
		}
		scalemin, _ := strconv.Atof64(sc[0])
		scalemax, _ := strconv.Atof64(sc[1])
		scaleincr, _ := strconv.Atof64(sc[2])
		measure, _ := strconv.Atof64(v.Measure)
		cmeasure, _ := strconv.Atof64(v.Cmeasure)

		// label the graph
		canvas.Text(tx, y+barheight/3, fmt.Sprintf("%s (%g)", v.Title, measure), "text-anchor:end;font-weight:bold")
		canvas.Text(tx, y+(barheight/3)+fontsize, v.Subtitle, "text-anchor:end;font-size:75%")

		// draw the scale
		scfmt := "%g"
		if fraction(scaleincr) > 0 {
			scfmt = "%.1f"
		}
		canvas.Gstyle("text-anchor:middle;font-size:75%")
		for sc := scalemin; sc <= scalemax; sc += scaleincr {
			scx := vmap(sc, scalemin, scalemax, 0, float64(maxwidth))
			canvas.Text(x+int(scx), y+scalesep+barheight+fontsize/2, fmt.Sprintf(scfmt, sc))
		}
		canvas.Gend()

		// draw the qualitative measures
		canvas.Gstyle("fill-opacity:0.5;fill:" + barcolor)
		canvas.Rect(x, y, maxwidth, barheight)
		for _, q := range qmeasures {
			qbarlength := vmap(q, scalemin, scalemax, 0, float64(maxwidth))
			canvas.Rect(x, y, int(qbarlength), barheight)
		}
		canvas.Gend()

		// draw the measure and the comparative measure
		barlength := int(vmap(measure, scalemin, scalemax, 0, float64(maxwidth)))
		canvas.Rect(x, y+qmheight, barlength, qmheight, "fill:"+datacolor)
		cmx := int(vmap(cmeasure, scalemin, scalemax, 0, float64(maxwidth)))
		if circlemark {
			canvas.Circle(x+cmx, y+barheight/2, barheight/6, "fill-opacity:0.3;fill:"+compcolor)
		} else {
			cbh := barheight / 4
			canvas.Line(x+cmx, y+cbh, x+cmx, y+barheight-cbh, "stroke-width:3;stroke:"+compcolor)
		}

		y += barheight + gutter // adjust vertical position for the next iteration
	}
	// if requested, place the title below the last bar
	if showtitle {
		canvas.Text(left, y+fontsize*2, title, "text-anchor:start;font-size:200%")
	}
}

//vmap maps one interval to another
func vmap(value float64, low1 float64, high1 float64, low2 float64, high2 float64) float64 {
	return low2 + (high2-low2)*(value-low1)/(high1-low1)
}

// fraction returns the fractions portion of a floating point number
func fraction(n float64) float64 {
	i := int(n)
	return n - float64(i)
}

// init sets up the command flags
func init() {
	flag.StringVar(&bgcolor, "bg", "white", "background color")
	flag.StringVar(&barcolor, "bc", "rgb(200,200,200)", "bar color")
	flag.StringVar(&datacolor, "dc", "darkgray", "data color")
	flag.StringVar(&compcolor, "cc", "black", "comparative color")
	flag.IntVar(&width, "w", 1024, "width")
	flag.IntVar(&height, "h", 800, "height")
	flag.IntVar(&barheight, "bh", 48, "bar height")
	flag.IntVar(&gutter, "g", 30, "gutter")
	flag.IntVar(&fontsize, "f", 18, "fontsize (px)")
	flag.BoolVar(&circlemark, "circle", false, "circle mark")
	flag.BoolVar(&showtitle, "showtitle", false, "show title")
	flag.StringVar(&title, "t", "Bullet Graphs", "title")
	flag.Parse()
}

// for every input file (or stdin), draw a bullet graph
// as specified by command flags
func main() {
	canvas := svg.New(os.Stdout)
	canvas.Start(width, height)
	canvas.Rect(0, 0, width, height, "fill:"+bgcolor)
	canvas.Gstyle(fmt.Sprintf(gstyle, fontsize))
	if len(flag.Args()) == 0 {
		dobg("", canvas)
	} else {
		for _, f := range flag.Args() {
			dobg(f, canvas)
		}
	}
	canvas.Gend()
	canvas.End()
}
