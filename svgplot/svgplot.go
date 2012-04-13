//svgplot -- plot data (a stream of x,y coordinates)
package main

import (
	"flag"
	"fmt"
	"io"
	"math"
	"os"

	"github.com/ajstarks/svgo"
)

// rawdata defines data as float64 x,y coordinates
type rawdata struct {
	x float64
	y float64
}

type options map[string]bool
type attributes map[string]string
type measures map[string]int

// plotset defines plot metadata
type plotset struct {
	opt  options
	attr attributes
	size measures
}

var (
	canvas                                                       = svg.New(os.Stdout)
	plotopt                                                      = options{}
	plotattr                                                     = attributes{}
	plotnum                                                      = measures{}
	ps                                                           = plotset{plotopt, plotattr, plotnum}
	plotw, ploth, plotc, gwidth, gheight, gutter, beginx, beginy int
)

const (
	globalfmt = "font-family:%s;font-size:%dpt;stroke-width:%dpx"
	linefmt   = "stroke:%s"
	barfmt    = linefmt + ";stroke-width:%dpx"
	ticfmt    = "stroke:rgb(200,200,200);stroke-width:1px"
	labelfmt  = ticfmt + ";text-anchor:end;fill:black"
	textfmt   = "stroke:none;baseline-shift:-33.3%"
)

// init initializes command flags and sets default options
func init() {
	// boolean options 
	showx := flag.Bool("showx", true, "show the xaxis")
	showy := flag.Bool("showy", true, "show the yaxis")
	showbar := flag.Bool("showbar", false, "show data bars")
	connect := flag.Bool("connect", true, "connect data points")
	showdot := flag.Bool("showdot", false, "show dots")
	showbg := flag.Bool("showbg", true, "show the background color")
	showfile := flag.Bool("showfile", false, "show the filename")
	
	// attributes
	bgcolor := flag.String("bgcolor", "rgb(240,240,240)", "plot background color")
	barcolor := flag.String("barcolor", "gray", "bar color")
	dotcolor := flag.String("dotcolor", "black", "dot color")
	linecolor := flag.String("linecolor", "gray", "line color")
	font := flag.String("font", "Calibri,sans", "font")
	plotlabel := flag.String("label", "", "plot label")
	
	// sizes
	dotsize := flag.Int("dotsize", 2, "dot size")
	linesize := flag.Int("linesize", 2, "line size")
	barsize := flag.Int("barsize", 2, "bar size")
	fontsize := flag.Int("fontsize", 11, "font size")
	xinterval := flag.Int("xint", 10, "x axis interval")
	yinterval := flag.Int("yint", 4, "y axis interval")

	// meta options
	flag.IntVar(&beginx, "bx", 100, "initial x")
	flag.IntVar(&beginy, "by", 50, "initial y")
	flag.IntVar(&plotw, "pw", 500, "plot width")
	flag.IntVar(&ploth, "ph", 500, "plot height")
	flag.IntVar(&plotc, "pc", 2, "plot columns")
	flag.IntVar(&gutter, "gutter", ploth/10, "gutter")
	flag.IntVar(&gwidth, "width", 1024, "canvas width")
	flag.IntVar(&gheight, "height", 768, "canvas height")

	flag.Parse()

	// fill in the plotset -- all options, attributes, and sizes
	plotopt["showx"] = *showx
	plotopt["showy"] = *showy
	plotopt["showbar"] = *showbar
	plotopt["connect"] = *connect
	plotopt["showdot"] = *showdot
	plotopt["showbg"] = *showbg
	plotopt["showfile"] = *showfile

	plotattr["bgcolor"] = *bgcolor
	plotattr["barcolor"] = *barcolor
	plotattr["linecolor"] = *linecolor
	plotattr["dotcolor"] = *dotcolor
	plotattr["font"] = *font
	plotattr["label"] = *plotlabel

	plotnum["dotsize"] = *dotsize
	plotnum["linesize"] = *linesize
	plotnum["fontsize"] = *fontsize
	plotnum["xinterval"] = *xinterval
	plotnum["yinterval"] = *yinterval
	plotnum["barsize"] = *barsize
}

// fmap maps world data to document coordinates
func fmap(value float64, low1 float64, high1 float64, low2 float64, high2 float64) float64 {
	return low2 + (high2-low2)*(value-low1)/(high1-low1)
}

// doplot opens a file and makes a plot
func doplot(x, y int, location string) {
	var f *os.File
	var err error
	if len(location) > 0 {
		if plotopt["showfile"] {
			plotattr["label"] = location
		}
		f, err = os.Open(location)
	} else {
		f = os.Stdin
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	nd, data := readxy(f)
	if nd >= 2 {
		plot(x, y, plotw, ploth, ps, data)
	}
	f.Close()
}

// plot places a plot at the specified location with the specified dimemsions
// usinng the specified settings, using the specified data
func plot(x, y, w, h int, settings plotset, d []rawdata) {
	if len(d) < 2 {
		fmt.Fprintf(os.Stderr, "%d is not enough points to plot\n", len(d))
		return
	}
	maxx, minx := d[0].x, d[0].x
	maxy, miny := d[0].y, d[0].y
	for _, v := range d {
		if v.x > maxx {
			maxx = v.x
		}
		if v.y > maxy {
			maxy = v.y
		}
		if v.x < minx {
			minx = v.x
		}
		if v.y < miny {
			miny = v.y
		}
	}
	if settings.opt["showbg"] {
		canvas.Rect(x, y, w, h, "fill:"+settings.attr["bgcolor"])
	}
	spacer := 10
	canvas.Gstyle(fmt.Sprintf(globalfmt,
		settings.attr["font"], settings.size["fontsize"], settings.size["linesize"]))
	if len(settings.attr["label"]) > 0 {
		canvas.Text(x, y-spacer, settings.attr["label"], textfmt+";font-size:120%")
	}
	px, py := 0, 0
	for i, v := range d {
		xp := int(fmap(v.x, minx, maxx, float64(x), float64(x+w)))
		yp := int(fmap(v.y, miny, maxy, float64(y), float64(y-h)))
		if settings.opt["showbar"] {
			canvas.Line(xp, yp+h, xp, y+h,
				fmt.Sprintf(barfmt, settings.attr["barcolor"], settings.size["barsize"]))
		}
		if settings.opt["showdot"] {
			canvas.Circle(xp, yp+h, settings.size["dotsize"], "fill:"+settings.attr["dotcolor"])
		}
		if settings.opt["connect"] && i > 0 && i < len(d) {
			canvas.Line(xp, yp+h, px, py+h, fmt.Sprintf(linefmt, settings.attr["linecolor"]))
		}
		if settings.opt["showx"] {
			if i%settings.size["xinterval"] == 0 {
				canvas.Text(xp, (y+h)+(spacer*2), fmt.Sprintf("%d", int(v.x)), "text-anchor:middle")
				canvas.Line(xp, (y + h), xp, (y+h)+spacer, ticfmt)
			}
		}
		px = xp
		py = yp
	}
	if settings.opt["showy"] {
		bot := math.Floor(miny)
		top := math.Ceil(maxy)
		yrange := top - bot
		interval := yrange / float64(settings.size["yinterval"])
		canvas.Gstyle(labelfmt)
		for yax := bot; yax <= top; yax += interval {
			yaxp := fmap(yax, bot, top, float64(y), float64(y-h))
			canvas.Text(x-spacer, int(yaxp)+h, fmt.Sprintf("%.1f", yax), textfmt)
			canvas.Line(x-spacer, int(yaxp)+h, x, int(yaxp)+h)
		}
		canvas.Gend()
	}
	canvas.Gend()
}

// readxy reads coordinates (x,y float64 values) from a io.Reader
func readxy(f io.Reader) (int, []rawdata) {
	var (
		r     rawdata
		err   error = nil
		n, nf int
	)
	data := make([]rawdata, 1)
	for ; err == nil; n++ {
		if n > 0 {
			data = append(data, r)
		}
		nf, err = fmt.Fscan(f, &data[n].x, &data[n].y)
		if nf != 2 {
			continue
		}
	}
	return n - 1, data[0 : n-1]
}

// plotgrid places plots on a grid, governed by a number of columns.
func plotgrid(x, y int, files []string) {
	px := x
	for i, f := range files {
		if i > 0 && i%plotc == 0 {
			px = x
			y += (ploth + gutter)
		}
		doplot(px, y, f)
		px += (plotw + gutter)
	}
}

// main plots data from specified files or standard input in a 
// grid where plotc specifies the number of columns.
func main() {
	canvas.Start(gwidth, gheight)
	canvas.Rect(0, 0, gwidth, gheight, "fill:white")
	filenames := flag.Args()
	if len(filenames) == 0 {
		doplot(beginx, beginy, "")
	} else {
		plotgrid(beginx, beginy, filenames)
	}
	canvas.End()
}