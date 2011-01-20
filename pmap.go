// pmap: percentage maps

package main

import (
	"flag"
	"strings"
	"strconv"
	"fmt"
	"svg"
	"os"
	"io"
	"xml"
)

type Pmap struct {
	Top   string "attr"
	Left  string "attr"
	Pdata []Pdata
}
type Pdata struct {
	Legend    string "attr"
	Stagger   string "attr"
	Alternate string "attr"
	Item      []Item
}
type Item struct {
	Name  string "chardata"
	Value string "attr"
}

var (
	width, height, fontsize, fontscale, round, gutter, pred, pgreen, pblue, oflen int
	bgcolor, colorspec, title                                                     string
	showpercent, showdata, alternate, showtitle, stagger, showlegend, showtotal   bool
	ofpct                                                                         float64
	leftmargin                                                                    = 40
	topmargin                                                                     = 40
	canvas                                                                        = svg.New(os.Stdout)
)

const (
	globalfmt   = "stroke-width:1;font-family:Calibri,sans-serif;text-anchor:middle;font-size:%dpt"
	legendstyle = "text-anchor:start;font-size:150%"
	linefmt     = "stroke:%s"
)

func dopmap(location string) {
	var f *os.File
	var err os.Error
	if len(location) > 0 {
		f, err = os.Open(location, os.O_RDONLY, 0)
	} else {
		f = os.Stdin
	}
	defer f.Close()
	if err == nil {
		readpmap(f)
	} else {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}

func readpmap(r io.Reader) {
	var pm Pmap
	if err := xml.Unmarshal(r, &pm); err == nil {
		drawpmap(pm)
	} else {
		fmt.Fprintf(os.Stderr, "Unable to parse pmap (%v)\n", err)
	}
}

func drawpmap(m Pmap) {
	fs := fontsize
	if len(m.Left) > 0 {
		leftmargin, _ = strconv.Atoi(m.Left)
	}
	if len(m.Top) > 0 {
		topmargin, _ = strconv.Atoi(m.Top)
	} else {
		topmargin = fs * fontscale
	}
	x := leftmargin
	y := topmargin
	for _, p := range m.Pdata {
		pmap(x, y, fs, p)
		y += fs*fontscale + (gutter + fs*2)
	}
}

func pmap(x, y, fs int, m Pdata) {
	var tfill, vfmt string
	var up bool
	h := fs * fontscale
	fw := fs * 80
	slen := fs + (fs / 2)
	up = false

	sum := 0.0
	for _, v := range m.Item {
		fv, _ := strconv.Atof64(v.Value)
		sum += fv
	}

	loffset := (fs * fontscale) + fs
	gline := fmt.Sprintf(linefmt, "gray")
	wline := fmt.Sprintf(linefmt, bgcolor)
	if len(m.Legend) > 0 && showlegend {
		if showtotal {
			canvas.Text(x, y-fs, fmt.Sprintf("%s (total: %.1f)", m.Legend, sum), legendstyle)
		} else {
			canvas.Text(x, y-fs, m.Legend, legendstyle)
		}
	}
	for i, p := range m.Item {
		k := p.Name
		v, _ := strconv.Atof64(p.Value)
		if v == 0.0 {
			continue
		}
		pct := v / sum
		pw := int(pct * float64(fw))
		xw := x + (pw / 2)
		yh := y + (h / 2)
		if pct >= .4 {
			tfill = "fill:white"
		} else {
			tfill = "fill:black"
		}
		if round > 0 {
			canvas.Roundrect(x, y, pw, h, round, round, pctfill(pred, pgreen, pblue, pct))
		} else {
			canvas.Rect(x, y, pw, h, pctfill(pred, pgreen, pblue, pct))
		}

		dy := yh + fs + (fs / 2)
		if pct <= ofpct || len(k) > oflen { // overflow label
			if up {
				dy -= loffset
				yh -= loffset
				canvas.Line(xw, y, xw, dy+(fs/2), gline)
			} else {
				dy += loffset
				yh += loffset
				canvas.Line(xw, y+h, xw, dy-(fs*3), gline)
			}
			if alternate {
				up = !up
				slen = fs * 2
			} else {
				slen = fs * 3
			}
			if stagger {
				loffset += slen
			}
			tfill = "fill:black"
		}
		canvas.Text(xw, yh, k, tfill)
		dpfmt := tfill + ";font-size:75%"
		if v-float64(int(v)) == 0.0 {
			vfmt = "%.0f"
		} else {
			vfmt = "%.1f"
		}
		switch {
		case showpercent && !showdata:
			canvas.Text(xw, dy, fmt.Sprintf("%.1f%%", pct*100), dpfmt)
		case showpercent && showdata:
			canvas.Text(xw, dy, fmt.Sprintf(vfmt+", %.1f%%", v, pct*100), dpfmt)
		case showdata && !showpercent:
			canvas.Text(xw, dy, fmt.Sprintf(vfmt, v), dpfmt)
		}
		x += pw
		if i < len(m.Item)-1 {
			canvas.Line(x, y, x, y+h, wline)
		}
	}
}

func pctfill(r, g, b int, v float64) string {
	d := int(255.0*v) - 255
	return canvas.RGB(r-d, g-d, b-d)
}

func colorparse(c string) (int, int, int) {
	s := strings.Split(c, ",", -1)
	if len(s) != 3 {
		return 0, 0, 0
	}
	r, _ := strconv.Atoi(s[0])
	g, _ := strconv.Atoi(s[1])
	b, _ := strconv.Atoi(s[2])
	return r, g, b
}

func dotitle(s string) {
	offset := 40
	canvas.Text(leftmargin, height-offset, s, "text-anchor:start;font-size:250%")
}

func init() {
	flag.IntVar(&width, "w", 1024, "width")
	flag.IntVar(&height, "h", 768, "height")
	flag.IntVar(&fontsize, "f", 12, "font size (pt)")
	flag.IntVar(&fontscale, "s", 5, "font scaling factor")
	flag.IntVar(&round, "r", 0, "rounded corner size")
	flag.IntVar(&gutter, "g", 100, "gutter")
	flag.IntVar(&oflen, "ol", 20, "overflow length")
	flag.StringVar(&bgcolor, "bg", "white", "background color")
	flag.StringVar(&colorspec, "c", "0,0,0", "color (r,g,b)")
	flag.StringVar(&title, "t", "Proportions", "title")
	flag.BoolVar(&showpercent, "p", false, "show percentage")
	flag.BoolVar(&showdata, "d", false, "show data")
	flag.BoolVar(&alternate, "a", false, "alternate overflow labels")
	flag.BoolVar(&stagger, "stagger", false, "stagger labels")
	flag.BoolVar(&showlegend, "showlegend", true, "show the legend")
	flag.BoolVar(&showtitle, "showtitle", false, "show the title")
	flag.BoolVar(&showtotal, "showtotal", false, "show totals in the legend")
	flag.Float64Var(&ofpct, "op", 0.05, "overflow percentage")
	flag.Parse()
	pred, pgreen, pblue = colorparse(colorspec)
}

func main() {
	canvas.Start(width, height)
	canvas.Rect(0, 0, width, height, "fill:"+bgcolor)
	canvas.Title(title)
	canvas.Gstyle(fmt.Sprintf(globalfmt, fontsize))

	if len(flag.Args()) == 0 {
		dopmap("")
	} else {
		for _, f := range flag.Args() {
			dopmap(f)
		}
	}

	if showtitle {
		dotitle(title)
	}
	canvas.Gend()
	canvas.End()
}
