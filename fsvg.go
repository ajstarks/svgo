// svg: generate SVG objects
//
// Anthony Starks, ajstarks@gmail.com

package svg

import (
	"fmt"
	"os"
	"xml"
	"strings"
)

const svginit = `<?xml version="1.0"?>
<svg xmlns="http://www.w3.org/2000/svg"
     xmlns:xlink="http://www.w3.org/1999/xlink"
     width="%d" height="%d">
`

type coord int
type dimen int
type rgbval byte


// Structure and Metadata

func Start(w dimen, h dimen)  { fmt.Printf(svginit, w, h) }
func End()                { fmt.Println("</svg>") }
func Gstyle(s string)     { fmt.Println(group("style", s)) }
func Gtransform(s string) { fmt.Println(group("transform", s)) }
func Gid(s string)        { fmt.Println(group("id", s)) }
func Gend()               { fmt.Println("</g>") }
func Def()                { fmt.Println("<defs>") }
func DefEnd()             { fmt.Println("</defs>") }
func Desc(s string)       { tt("desc", "", s) }
func Title(s string)      { tt("title", "", s) }
func Use(x coord, y coord, link string, s ...string) {
	fmt.Printf(`<use %s %s %s`, loc(x, y), href(link), endstyle(s))
}

// Shapes

func Circle(x coord, y coord, r dimen, s ...string) {
	fmt.Printf(`<circle cx="%d" cy="%d" r="%d" %s`, x, y, r, endstyle(s))
}

func Ellipse(x coord, y coord, w dimen, h dimen, s ...string) {
	fmt.Printf(`<ellipse cx="%d" cy="%d" rx="%d" ry="%d" %s`,
		x, y, w, h, endstyle(s))
}

func Polygon(x []coord, y []coord, s ...string) { poly(x, y, "polygon", s) }

func Rect(x coord, y coord, w dimen, h dimen, s ...string) {
	fmt.Printf(`<rect %s %s`, dim(x, y, w, h), endstyle(s))
}

func Roundrect(x coord, y coord, w dimen, h dimen, rx dimen, ry dimen, s ...string) {
	fmt.Printf(`<rect %s rx="%d" ry="%d" %s`, dim(x, y, w, h), rx, ry, endstyle(s))
}

func Square(x coord, y coord, s dimen, style ...string) {
	Rect(x, y, s, s, style)
}

// Curves

func Arc(sx coord, sy coord, ax coord, ay coord, r dimen, large bool, sweep bool, ex coord, ey coord, s ...string) {
	fmt.Printf(`%s A%s %d %s %s %s" %s`,
		ptag(sx, sy), coordinate(ax, ay), r, onezero(large), onezero(sweep), coordinate(ex, ey), endstyle(s))
}

func Bezier(sx coord, sy coord, cx coord, cy coord, px coord, py coord, ex coord, ey coord, s ...string) {
	fmt.Printf(`%s C%s %s %s" %s`,
		ptag(sx, sy), coordinate(cx, cy), coordinate(px, py), coordinate(ex, ey), endstyle(s))
}

func Qbezier(sx coord, sy coord, cx coord, cy coord, ex coord, ey coord, tx coord, ty coord, s ...string) {
	fmt.Printf(`%s Q%s %s T%s" %s`,
		ptag(sx, sy), coordinate(cx, cy), coordinate(ex, ey), coordinate(tx, ty), endstyle(s))
}

// Lines

func Line(x1 coord, y1 coord, x2 coord, y2 coord, s ...string) {
	fmt.Printf(`<line x1="%d" y1="%d" x2="%d" y2="%d" %s`, x1, y1, x2, y2, endstyle(s))
}

func Polyline(x []coord, y []coord, s ...string) { poly(x, y, "polyline", s) }


// Image

func Image(x coord, y coord, w dimen, h dimen, link string, s ...string) {
	fmt.Printf("<image %s %s %s", dim(x, y, w, h), href(link), endstyle(s))
}

// Text

func Text(x coord, y coord, t string, s ...string) {
	if len(s) > 0 {
		tt("text", " "+loc(x, y)+" "+style(s[0]), t)
	} else {
		tt("text", " "+loc(x, y)+" ", t)
	}
}

// Color

func RGB(r rgbval, g rgbval, b rgbval) string { return fmt.Sprintf(`fill:rgb(%d,%d,%d)`, r, g, b) }
func RGBA(r rgbval, g rgbval, b rgbval, a float) string {
	return fmt.Sprintf(`fill-opacity:%.2f; %s`, a, RGB(r, g, b))
}

// Utility

func Grid(x coord, y coord, w coord, h coord, n coord, s ...string) {

	if len(s) > 0 {
		Gstyle(s[0])
	}

	for ix := x; ix <= x+w; ix += n {
		Line(ix, y, ix, y+h)
	}
/*
	for iy := y; iy <= y+h; iy += n {
		Line(x, iy, x+w, iy)
	}
	*/
	if len(s) > 0 {
		Gend()
	}

}

// Support functions

func style(s string) string {
	if len(s) > 0 {
		return fmt.Sprintf(`style="%s"`, s)
	}
	return s
}

func pp(x []coord, y []coord, tag string) {
	if len(x) != len(y) {
		return
	}
	fmt.Print(tag)
	for i := 0; i < len(x); i++ {
		fmt.Print(coordinate(x[i], y[i]) + " ")
	}
}

func endstyle(s []string) string {
	if len(s) > 0 {
		if strings.Index(s[0], "=") > 0 {
			return s[0] + "/>\n"
		} else {
			return style(s[0]) + "/>\n"
		}
	}
	return "/>\n"
}

func tt(tag string, attr string, s string) {
	fmt.Print("<" + tag + attr + ">")
	xml.Escape(os.Stdout, []byte(s))
	fmt.Println("</" + tag + ">")
}

func poly(x []coord, y []coord, tag string, s ...string) {
	pp(x, y, "<"+tag+` points="`)
	fmt.Print(`" ` + endstyle(s))
}

func onezero(flag bool) string {
	if flag {
		return "1"
	}
	return "0"
}

func coordinate(x coord, y coord) string { return fmt.Sprintf(`%d,%d`, x, y) }
func ptag(x coord, y coord) string  { return fmt.Sprintf(`<path d="M%s`, coordinate(x, y)) }
func loc(x coord, y coord) string   { return fmt.Sprintf(`x="%d" y="%d"`, x, y) }
func href(s string) string      { return fmt.Sprintf(`xlink:href="%s"`, s) }
func dim(x coord, y coord, w dimen, h dimen) string {
	return fmt.Sprintf(`x="%d" y="%d" width="%d" height="%d"`, x, y, w, h)
}
func group(tag string, value string) string { return fmt.Sprintf(`<g %s="%s">`, tag, value) }
