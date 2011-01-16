// Flickr50 -- given a search term, display 10x5 image grid, sorted by interestingness

package main

import (
	"github.com/ajstarks/svgo"
	"os"
	"http"
	"xml"
	"fmt"
)

type FlickrResp struct {
	Stat   string "attr"
	Photos Photos
}

type Photos struct {
	Page    string "attr"
	Pages   string "attr"
	Perpage string "attr"
	Total   string "attr"
	Photo   []Photo
}

type Photo struct {
	Id       string "attr"
	Owner    string "attr"
	Secret   string "attr"
	Server   string "attr"
	Farm     string "attr"
	Title    string "attr"
	Ispublic string "attr"
	Isfriend string "attr"
	IsFamily string "attr"
}

var (
	width  = 805
	height = 500
	canvas = svg.New(os.Stdout)
)

const (
	apifmt      = "http://api.flickr.com/services/rest/?method=%s&api_key=%s&%s=%s&per_page=50&sort=interestingness-desc"
	urifmt      = "http://farm%s.static.flickr.com/%s/%s.jpg"
	apiKey      = "APIKEYHERE"
	textStyle   = "font-family:Calibri,sans-serif; font-size:48px; fill:white; text-anchor:start"
	imageWidth  = 75
	imageHeight = 75
)

// FlickrAPI calls the API given a method with single name/value pair
func flickrAPI(method, name, value string) string {
	return fmt.Sprintf(apifmt, method, apiKey, name, value)
}

// makeURI converts the elements of a photo into a Flickr photo URI
func makeURI(p Photo, imsize string) string {
	im := p.Id + "_" + p.Secret

	if len(imsize) > 0 {
		im += "_" + imsize
	}
	return fmt.Sprintf(urifmt, p.Farm, p.Server, im)
}

// imageGrid reads the response from Flickr, and creates a grid of images
func imageGrid(f FlickrResp, x, y, cols, gutter int, imgsize string) {
	if f.Stat != "ok" {
		fmt.Fprintf(os.Stderr, "%v\n", f)
		return
	}
	xpos := x
	for i, p := range f.Photos.Photo {
		if i%cols == 0 && i > 0 {
			xpos = x
			y += (imageHeight + gutter)
		}
		canvas.Link(makeURI(p, ""), p.Title)
		canvas.Image(xpos, y, imageWidth, imageHeight, makeURI(p, "s"))
		canvas.LinkEnd()
		xpos += (imageWidth + gutter)
	}
}

// fs calls the Flickr API to perform a photo search
func fs(s string) {
	var f FlickrResp
	r, _, weberr := http.Get(flickrAPI("flickr.photos.search", "text", s))
	if weberr != nil || r.StatusCode != 200 {
		fmt.Fprintf(os.Stderr, "%v (status=%d)\n", weberr, r.StatusCode)
		return
	}
	xmlerr := xml.Unmarshal(r.Body, &f)
	if xmlerr != nil {
		fmt.Fprintf(os.Stderr, "%v\n", xmlerr)
		return
	}
	canvas.Title(s)
	imageGrid(f, 5, 5, 10, 5, "s")
	canvas.Text(20, height-40, s, textStyle)
}

// for each search term on the commandline, create a photo grid
func main() {
	for i := 1; i < len(os.Args); i++ {
		canvas.Start(width, height)
		canvas.Rect(0, 0, width, height, "fill:black")
		fs(http.URLEscape(os.Args[i]))
		canvas.End()
	}
}
