package main

import (
	"./svg"
	"fmt"
	"os"
	"bufio"
	"strings"
	"flag"
)

var (
  width = flag.Int("w", 1024, "Overall width")
	height = flag.Int("h", 768, "Overall height")
	sw = flag.Int("sw", 350, "Server width")
	sh = flag.Int("sh", 150, "Server height")
	lh = flag.Int("lh", 25, "Label height")
	ft = flag.Bool("ft", false, "Flat top flag")
	linkf = flag.Bool("l", false, "Link flag")
	inputfile = flag.String("f", "/dev/stdin", "Input file")
	grid = flag.Bool("g", false, "show coordinate grid")
	topcolor = flag.String("tc", svg.RGB(127, 0, 0), "topcolor")
	appcolor = flag.String("ac", svg.RGB(100, 100, 100), "appcolor")
	appshape = flag.String("as", "n", "Application shape")
)

func ngrid(x, y, w, h, n, m int) {
	svg.Grid(x, y, w, h, n, "stroke:rgb(210,210,210)")
	fs := 12
	svg.Gstyle(fmt.Sprintf("font-size:%d; font-family:monospace", fs))
	for ix := x; ix <= w; ix += m {
		svg.Text(ix, (h / 2), fmt.Sprintf("%d", ix), "fill:red;text-anchor:end")
	}
	for iy := y; iy <= h; iy += m {
		svg.Text((w / 2), iy, fmt.Sprintf("%d", iy), "fill:blue;text-anchor:start")
	}
	svg.Gend()
}

func roundtop(x, y, w, h, r int, s string) {
	svg.Gstyle("stroke:none")
	svg.Roundrect(x, y, w, h, r, r, s)
	svg.Rect(x, y+r, w, h-r, s)
	svg.Gend()
}

func note(x, y, w, h, n int, style string) {
	var xp = []int{x, x + (w - n), x + w, x + w, x, x}
	var yp = []int{y, y, y + n, y + h, y + h, y}
	svg.Polyline(xp, yp, style)
}

func folder(x, y, w, h, tw, th int, style string) {
	var xp = []int{x, x + th, x + (tw - th), x + tw}
	var yp = []int{y + th, y, y, y + th}
	svg.Gstyle(style)
	svg.Polyline(xp, yp, "")
	svg.Rect(x, y+th, w, h-th, "")
	svg.Gend()
}

func rfolder(x, y, w, h, tw, th int, style string) {
	var xp = []int{x + th, x + (2 * th), x + th + (tw - th), x + tw + th}
	var yp = []int{y + th, y, y, y + th}
	svg.Gstyle(style)
	svg.Polyline(xp, yp)
	svg.Roundrect(x, y+th, w, h-th, th, th)
	svg.Gend()
}

func rnote(x, y, w, h, n int, style string, label ...string) {
	note(x, y, w, h-n, n, style)
	svg.Roundrect(x, y+(n*2), w, h-(n*2), n, n, style)
	if len(label) > 0 {
		svg.Text(x+(w/2), y+(h/2), label[0], "text-anchor:middle")
	}
}

func labelshape(x, y, w, h int, s string, ts int, scolor, tcolor string, rtype string) {

	tx := x + (w / 2)
	ty := y + (h / 2) + (ts / 4)

	switch rtype {
	case "r":
		svg.Rect(x, y, w, h, scolor)
	case "rr":
		svg.Roundrect(x, y, w, h, ts/2, ts/2, scolor)
	case "n":
		note(x, y, w, h, h/4, scolor)
	case "rn":
		rnote(x, y, w, h, h/4, scolor)
	case "f":
		folder(x, y, w, h, h/2, h/2, scolor)
	case "rf":
		rfolder(x, y, w, h, h/4, h/4, scolor)
	case "c":
		svg.Circle(x, y, w/4, scolor)
		tx = x
		ty = y + (ts / 4)
	case "e":
		svg.Ellipse(x, y, w/4, h/4, scolor)
		tx = x
		ty = y + (ts / 4)
	}
	svg.Text(tx, ty, s, fmt.Sprintf("%s; font-size:%d; text-anchor:middle", tcolor, ts))
}

func server(x, y, w, h, r, o int, data []string) {

	if *linkf {
	  svg.Link(data[0]+".svg", data[0])
	}

	if *ft {
	  svg.Rect(x, y, w, r*2, *topcolor)
	} else {
	  roundtop(x, y, w, r*2, r, *topcolor)
	}

	// Draw the boxes

	yr := y + r // beginning of box area, after the top
	ih := h - r
	o2 := o * 2
	svg.Rect(x, yr, w, ih, "fill:rgb(220,220,220); stroke:none")
	svg.Rect(x+o, yr+o, w-o, ih-o, "fill:rgb(200, 200, 200); stroke:none")
	svg.Rect(x+o2, yr+o2, w-o2, ih-o2, "fill:rgb(180,180,180); stroke:none")
	
	
	if *linkf {
	  svg.LinkEnd()
	}

	// Adjust font sizes based on r

	var fa int

	if r > o {
		fa = o
	} else {
		fa = r
	}

	basefs := fa - (fa / 5)
	afs := basefs / 2
	hwfs := basefs - (basefs / 4)
	osfs := basefs - (2*basefs)/5
	adj := fa / 4

	// Text annotations

	oscomp := strings.Split(data[5], "~", 3)
	if len(data) > 4 && len(oscomp) > 0 {
		svg.Gstyle("font-family:Calibri")
		svg.Text(x+r, yr-adj, data[0], fmt.Sprintf(`fill:white;text-anchor:start;font-size:%d`, basefs))
		svg.Text((x+w)-r, yr-adj, data[1], fmt.Sprintf(`fill:white;text-anchor:end;font-size:%d`, afs))
		svg.Text(x+o, (yr+o)-adj, data[2], fmt.Sprintf(`text-anchor:start;font-size:%d`, hwfs))
		svg.Text((x+w)-r, (yr+o)-adj, data[3]+" GB/"+data[4]+" GB", fmt.Sprintf(`text-anchor:end;font-size:%d`, afs))
		svg.Text(x+o2, (yr+o2)-adj, oscomp[0], fmt.Sprintf(`font-size:%d`, osfs))
		svg.Gend()

		apps := strings.Split(oscomp[1], ",", 10)
		/*
			fmt.Println("<!--")
			fmt.Println("data:  ", data)
			fmt.Println("oscomp:", oscomp)
			fmt.Println("osname:", oscomp[0])
			fmt.Println("apps:  ", apps)
			fmt.Println("-->")
		*/
		// Components

		cx := x + o2 + adj
		cy := yr + o2 + adj
		cfs := float(basefs) - float(basefs)*.50
		if len(apps) > 0 {
			components(cx, cy, cx+(w-o2), o, int(cfs), apps)
		}
	}
}

func components(x int, y int, w int, h int, ts int, labels []string) {
	wc := 0
	gutter := 5
	xo := x
	ta := ts / 2
	svg.Gstyle("font-family:Calibri")
	for i := 0; i < len(labels); i++ {
		if len(labels[i]) < 6 {
			wc = len(labels[i]) * ts
		} else {
			wc = len(labels[i]) * ta
		}

		if x > (w - wc) {
			y += (h + gutter)
			x = xo
		}
		labelshape(x, y, wc, h, labels[i], ts, *appcolor, "fill:white", *appshape)
		x += (wc + gutter)
	}
	svg.Gend()
}

func readserver(file string) {
	r, _ := os.Open(file, os.O_RDONLY, 0444)
	defer r.Close()
	in := bufio.NewReader(r)

	maxf := 6
	y := 0
	x := 0
	ng := 0
	nr := 0
	gutter := 10
	w := *sw
	h := *sh
	for {
		line, err := in.ReadString('\n')
		nr++
		if err != nil {
			break
		}
		parts := strings.Split(line[0:len(line)-1], "\t", maxf)
		if len(parts) != maxf {
			continue
		}
		ng++

		if x > (*width - w) {
			y += (h + gutter)
			x = 0
		}
		server(x, y, w, h, *lh, *lh, parts)
		x += (w + gutter)
	}
	fmt.Println("<!--", ng, nr, "-->")
}

func background(v int) { svg.Rect(0, 0, *width, *height, svg.RGB(v, v, v)) }


func main() {
	flag.Parse()
	svg.Start(*width, *height)
	background(255)
	readserver(*inputfile)
	if *grid {
		ngrid(0, 0, *width, *height, *lh, *lh)
	}
	svg.End()
}
