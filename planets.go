// Planets: an exploration of scale
// Anthony Starks, ajstarks@gmail.com

package main

import (
	svglib "svg"
	"os"
	"flag"
	"image"
	"image/png"
)


var ssDist = []float{
	0.00,  // Sun
	0.34,  // Mercury
	0.72,  // Venus
	1.00,  // Earth
	1.54,  // Mars
	5.02,  // Jupiter
	9.46,  // Saturn
	20.11, // Uranus
	30.08} // Netpune


var ssRad = []float{ // Miles
	423200.0, // Sun
	1516.0,   // Mercury
	3760.0,   // Venus
	3957.0,   // Earth
	2104.0,   // Mars
	42980.0,  // Jupiter
	35610.0,  // Saturn
	15700.0,  // Uranus
	15260.0}  // Neptune


var ssColor = []string{ // R, G, B
	// Computed from images
	//  "CE3903, // Sun
	//  "7B5628, // Mercury
	//  "5F5D56, // Venus
	//  "555864, // Earth
	//  "614F3C, // Mars
	//  "735F52, // Jupiter
	//  "9C9383, // Saturn
	//  "556769, // Uranus
	//  "324A72  // Neptune

	//  Eyeballed from image
	"F7730C", // Sun
	"FAF8F2", // Mercury
	"FFFFF2", // Venus
	"0B5CE3", // Earth
	"F0C61D", // Mars
	"FDC791", // Jupiter
	"E0C422", // Saturn
	"DCF1F5", // Uranus
	"39B6F7"} // Neptune

var ssImages = []string{
	"images/sun.png",
	"images/mercury.png",
	"images/venus.png",
	"images/earth.png",
	"images/mars.png",
	"images/jupiter.png",
	"images/saturn.png",
	"images/uranus.png",
	"images/neptune.png"}


var showimages *bool = flag.Bool("i", true, "show images")

var svg = svglib.New(os.Stdout)

func loadimage(path string) image.Image {
	f, err := os.Open(path, os.O_RDONLY, 0)
	defer f.Close()
	if err != nil {
		return nil
	}
	p, perr := png.Decode(f)
	if perr != nil {
		return nil
	}
	return p
}


func vmap(value float, low1 float, high1 float, low2 float, high2 float) float {
	return low2 + (high2-low2)*(value-low1)/(high1-low1)
}


func main() {

	width := 1300
	height := 200

	flag.Parse()
	svg.Start(width, height)
	svg.Title("Planets")
	svg.Rect(0, 0, width, height, "fill:black")
	svg.Gstyle("stroke:none")
	nobj := len(ssDist)
	y := height / 2
	margin := 100
	minsize := 7.0
	labeloc := height / 4

	var x, r, imScale, maxh float

	if *showimages {
		maxh = float(height) / 4.0
	} else {
		maxh = float(height) / minsize
	}

	for i := 1; i < nobj; i++ {
		x = vmap(ssDist[i], ssDist[1], ssDist[nobj-1], float(margin), float(width-margin))
		r = (vmap(ssRad[i], ssRad[1], ssRad[nobj-1], minsize, maxh)) / 2
		if *showimages {
			p := loadimage(ssImages[i])
			if p == nil {
				continue
			}
			imScale = r / float(p.Width())
			hs := float(p.Height()) * imScale
			dy := y - (int(hs) / 2) // center the image
			svg.Image(int(x), dy, int(r), int(hs), ssImages[i])
		} else {
			svg.Circle(int(x), int(y), int(r), "fill:#"+ssColor[i])
		}
		if ssDist[i] == 1.0 { // earth
			svg.Line(int(x), int(y), int(x), int(y)-labeloc, "stroke:white")
			svg.Text(int(x), int(y)-labeloc-10, "You are here", "fill:white; font-size:14; font-family:Calibri; text-anchor:middle")
		}
	}
	svg.Gend()
	svg.End()
}
