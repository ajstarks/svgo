// ltr: Layer Tennis remixes

package main

import (
	"github.com/ajstarks/svgo"
	"os"
	"flag"
)

var (
	canvas                            = svg.New(os.Stdout)
	poster, opacity, row, col, offset bool
	title                             string
)

const (
	width  = 900
	height = 280
	ni     = 11
)

func imagefiles(directory string) []string {
	f, ferr := os.Open(directory)
	defer f.Close()
	if ferr != nil {
		return nil
	}
	files, derr := f.Readdir(-1)
	if derr != nil || len(files) == 0 {
		return nil
	}
	names := make([]string, len(files))
	for i, v := range files {
		names[i] = directory + "/" + v.Name
	}
	return names
}

func ltposter(x, y, w, h int, f []string) {
	canvas.Image(x, y, w*2, h*2, f[0]) // first file, assumed to be the banner
	y = y + (h * 2)
	for i := 1; i < len(f); i += 2 {
		canvas.Image(x, y, w, h, f[i])
		canvas.Image(x+w, y, w, h, f[i+1])
		if i%2 == 1 {
			y += h
		}
	}
}

func ltcol(x, y, w, h int, f []string) {
	for i := 0; i < len(f); i++ {
		canvas.Image(x, y, w, h, f[i])
		y += h
	}
}

func ltop(x, y, w, h int, f []string) {
	for i := 1; i < len(f); i++ { // skip the first file, assumed to be the banner
		canvas.Image(x, y, w, h, f[i], "opacity:0.2")
	}
}

func ltrow(x, y, w, h int, f []string) {
	for i := 0; i < len(f); i++ {
		canvas.Image(x, y, w, h, f[i])
		x += w
	}
}

func ltoffset(x, y, w, h int, f []string) {
	for i := 1; i < len(f); i++ { // skip the first file, assumed to be the banner

		if i%2 == 0 {
			x += w
		} else {
			x = 0
		}
		canvas.Image(x, y, w, h, f[i])
		y += h
	}
}
func dotitle(s string) {
	if len(title) > 0 {
		canvas.Title(title)
	} else {
		canvas.Title(s)
	}
}

func init() {
	flag.BoolVar(&poster, "poster", false, "poster style")
	flag.BoolVar(&opacity, "opacity", false, "opacity style")
	flag.BoolVar(&row, "row", false, "display is a single row")
	flag.BoolVar(&col, "col", false, "display in a single column")
	flag.BoolVar(&offset, "offset", false, "display in a row, even layers offset")
	flag.StringVar(&title, "title", "", "title")
	flag.Parse()
}

func main() {
	x := 0
	y := 0
	nd := len(flag.Args())
	for i, dir := range flag.Args() {
		filelist := imagefiles(dir)
		if len(filelist) != ni || filelist == nil {
			continue
		}
		switch {

		case opacity:
			if i == 0 {
				canvas.Start(width*nd, height*nd)
				dotitle(dir)
			}
			ltop(x, y, width, height, filelist)
			y += height

		case poster:
			if i == 0 {
				canvas.Start(width, ((height*(ni-1)/4)+height)*nd)
				dotitle(dir)
			}
			ltposter(x, y, width/2, height/2, filelist)
			y += (height * 3) + (height / 2)

		case col:
			if i == 0 {
				canvas.Start(width*nd, height*ni)
				dotitle(dir)
			}
			ltcol(x, y, width, height, filelist)
			x += width

		case row:
			if i == 0 {
				canvas.Start(width*ni, height*nd)
				dotitle(dir)
			}
			ltrow(x, y, width, height, filelist)
			y += height

		case offset:
			n := ni - 1
			pw := width * 2
			ph := nd * (height * (n))
			if i == 0 {
				canvas.Start(pw, ph)
				canvas.Rect(0, 0, pw, ph, "fill:white")
				dotitle(dir)
			}
			ltoffset(x, y, width, height, filelist)
			y += n * height

		}
	}
	canvas.End()
}
