// codepic -- produce code+output sample suitable for slides

package main

import (
	"os"
	"fmt"
	"bufio"
	"flag"
	"strings"
	"image/png"
	"github.com/ajstarks/svgo"
)

var (
	canvas                                                    = svg.New(os.Stdout)
	font                                                      string
	codeframe, picframe                                       bool
	linespacing, fontsize, top, left, boxwidth, width, height int
	framestyle                                                = "stroke:gray;stroke-dasharray:1,1;fill:none"
	labelstyle                                                = "text-anchor:middle;font-size:200%"
	codefmt                                                   = "font-family:%s;font-size:%dpx"
)

func slide(filename string) {
	var line, basename, imgname string
	var rerr os.Error
	bn := strings.Split(filename, ".")
	if len(bn) > 0 {
		basename = bn[0]
		imgname = basename + ".png"
	} else {
		fmt.Fprintf(os.Stderr, "cannot get the basename for %s\n", filename)
		return
	}
	f, oerr := os.Open(filename)
	defer f.Close()
	if oerr != nil {
		fmt.Fprintf(os.Stderr, "%v\n", oerr)
		return
	}
	imgf, imgerr := os.Open(imgname)
	defer imgf.Close()
	if imgerr != nil {
		fmt.Fprintf(os.Stderr, "%v\n", imgerr)
		return
	}
	imginfo, imgerr := png.DecodeConfig(imgf)
	if imgerr != nil {
		fmt.Fprintf(os.Stderr, "%v\n", imgerr)
		return
	}

	lx := width/2 + imginfo.Width/2
	ly := height - 10
	in := bufio.NewReader(f)
	canvas.Start(width, height)
	canvas.Title(basename)

	canvas.Rect(0,0,width,height,"fill:white")
	if picframe {
		canvas.Rect(width/2, top, imginfo.Width, imginfo.Height, framestyle)
	}
	canvas.Image(width/2, top, imginfo.Width, imginfo.Height, imgname, "image")
	canvas.Gstyle(fmt.Sprintf(codefmt, font, fontsize))
	canvas.Text(lx, ly, filename, labelstyle)
	y := top + fontsize*2
	for x := left + fontsize; rerr == nil; y += linespacing {
		line, rerr = in.ReadString('\n')
		if len(line) > 0 {
			canvas.Text(x, y, line[0:len(line)-1], `xml:space="preserve"`)
		}
	}
	if codeframe {
		canvas.Rect(top, left, left+boxwidth, y, framestyle)
	}
	canvas.Gend()
	canvas.End()
}

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

func main() {
	for _, f := range flag.Args() {
		slide(f)
	}
}
