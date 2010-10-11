package main

import (
	"fmt"
	"os"
	"svg"
	"strings"
	"bufio"
	"flag"
)

var (
	canvas   = svg.New(os.Stdout)
	filename = flag.String("f", "svgcolors.txt", "input file")
	fontname = flag.String("font", "Calibri,sans-serif", "fontname")
	outline  = flag.Bool("o", false, "outline")
	neg      = flag.Bool("n", false, "negative")
	showrgb  = flag.Bool("rgb", true, "show RGB")
	fontsize = flag.Int("fs", 12, "fontsize")
	width    = flag.Int("w", 1600, "width")
	height   = flag.Int("h", 900, "height")
	rowsize  = flag.Int("r", 32, "rowsize")
	colw     = flag.Int("c", 320, "column size")
	swatch   = flag.Int("s", 16, "swatch size")
	gutter   = flag.Int("g", 11, "gutter")
)

func main() {
	top := 32
	left := 32
	var (
		colorfmt, tcolor, line string
		err                    os.Error = nil
	)
	flag.Parse()
	f, oerr := os.Open(*filename, os.O_RDONLY, 0444)
	if oerr != nil {
		fmt.Fprintf(os.Stderr, "%v\n", oerr)
		return
	}
	in := bufio.NewReader(f)
	canvas.Start(*width, *height)
	if *neg {
		canvas.Rect(0, 0, *width, *height, "fill:black")
		tcolor = "white"
	} else {
		canvas.Rect(0, 0, *width, *height, "fill:white")
		tcolor = "black"
	}
	canvas.Gstyle(fmt.Sprintf("font-family:%s;font-size:%dpt;fill:%s;stroke-width:2",
		*fontname, *fontsize, tcolor))

	for x, y, nr := left, top, 0; err == nil; nr++ {
		line, err = in.ReadString('\n')
		fields := strings.Split(strings.TrimSpace(line), "\t", -1)
		if nr%*rowsize == 0 && nr > 0 {
			x += *colw
			y = top
		}
		if len(fields) == 3 {
			colorfmt = "fill:" + fields[1]
			if *outline {
				colorfmt = colorfmt + ";stroke:" + tcolor
			}
			canvas.Circle(x, y, *swatch/2, colorfmt)
			canvas.Text(x+*swatch, y+(*swatch/4), fields[0], "stroke:none")
			if *showrgb {
				canvas.Text(x+((*colw*4)/5), y+(*swatch/4), fields[2], "text-anchor:end;fill:gray")
			}
		}
		y += (*swatch + *gutter)
	}
	canvas.Gend()
	canvas.End()
}
