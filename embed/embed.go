// embed -- including the contents of an SVG within your own
// +build !appengine

package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"

	svg "github.com/ajstarks/svgo"
)

var (
	canvas   = svg.New(os.Stdout)
	filename string
)

var (
	titleHeight = 64
	borderSize  = 4
	borderfmt   = "fill:transparent;stroke:red;stroke-width:%d"
	titlefmt    = `text-anchor:middle;width:%d;font-size:32px`
	groupfmt    = `transform="translate(%d, %d)"`
)

func embed() error {
	// load and parse svg
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	s, err := parseSVG(f)
	if err != nil {
		return err
	}
	width := s.Width + 2*borderSize
	height := s.Height + titleHeight + 2*borderSize
	// create svg
	canvas.Start(width, height)
	placeHeader(width, height, f.Name())
	// embed svg
	placeSVG(s)
	return nil
}

// SVG contains the parsed attributes and xml from the given file.
type SVG struct {
	// Width and Height are attributes of the <svg> tag
	Width  int `xml:"width,attr"`
	Height int `xml:"height,attr"`
	// Doc is all all of the contents within the <svg> tags, specified by the
	// `innerxml` struct tag
	Doc []byte `xml:",innerxml"`
}

func parseSVG(src io.Reader) (SVG, error) {
	var s SVG
	data, err := io.ReadAll(src)
	if err != nil {
		return SVG{}, err
	}
	if xml.Unmarshal(data, &s); err != nil {
		return SVG{}, err
	}
	return s, nil
}

func placeHeader(width, height int, name string) {
	// add border
	canvas.Rect(0, 0, width, height, fmt.Sprintf(borderfmt, borderSize))
	// add title from file name
	canvas.Text(width/2, titleHeight*3/4, name, fmt.Sprintf(titlefmt, width))
}

func placeSVG(s SVG) {
	// create clip path of svg size
	canvas.Group(`clip-path="url(#embed)"`, fmt.Sprintf(groupfmt, borderSize, titleHeight+borderSize))
	canvas.ClipPath(`id="embed"`)
	canvas.Rect(0, 0, s.Width, s.Height)
	canvas.ClipEnd()
	// append embedded svg
	canvas.Writer.Write(s.Doc)
	canvas.Gend()
	canvas.End()
}

func init() {
	flag.StringVar(&filename, "f", "embed.svg", "file name")
	flag.Parse()
}

func main() {
	if err := embed(); err != nil {
		fmt.Fprintf(os.Stdout, "error: %v\n", err)
		os.Exit(1)
	}
}
