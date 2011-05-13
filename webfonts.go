package main

import (
	"os"
	"fmt"
	"http"
	"io/ioutil"
	"strings"
	"github.com/ajstarks/svgo"
)

var (
	g        = svg.New(os.Stdout)
	width    = 500
	height   = 1100
	fontlist = "Sue Ellen Francisco|Over the Rainbow|Pacifico|Inconsolata|Miltonian|Megrim|Monofett|Permanent Marker|Homemade Apple|Ultra"
)

const (
	gwfURI  = "http://fonts.googleapis.com/css?family=%s"
	fontfmt = "<style type=\"text/css\">\n<![CDATA[\n%s]]>\n</style>\n"
	gfmt    = "fill:white;font-size:36pt;text-anchor:middle"
)

func googlefont(fontname string) string {
	r, _, err := http.Get(fmt.Sprintf(gwfURI, http.URLEscape(fontname)))
	defer r.Body.Close()
	if err != nil {
		return ""
	}
	b, rerr := ioutil.ReadAll(r.Body)
	if rerr != nil || r.StatusCode != http.StatusOK {
		return ""
	}
	return string(b)
}

func defineFont(s string) {
	g.Def()
	fmt.Fprintf(g.Writer, fontfmt, googlefont(s))
	g.DefEnd()
}

func main() {
	g.Start(width, height)
	defineFont(fontlist)
	g.Rect(0, 0, width, height)
	g.Ellipse(width/2, height+50, width/2, height/5, "fill:rgb(44,77,232)")
	g.Gstyle(gfmt)
	for i, f := range strings.Split(fontlist, "|", -1) {
		g.Text(width/2, (i+1)*100, "Hello, World", "font-family:"+f)
	}
	g.Gend()
	g.End()
}
