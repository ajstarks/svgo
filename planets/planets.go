// Planets: an exploration of scale
// Anthony Starks, ajstarks@gmail.com

package main

import (
	"github.com/ajstarks/svgo"
	"os"
	"flag"
	"image/png"
)

var ssDist = []float64{
	0.00,  // Sun
	0.34,  // Mercury
	0.72,  // Venus
	1.00,  // Earth
	1.54,  // Mars
	5.02,  // Jupiter
	9.46,  // Saturn
	20.11, // Uranus
	30.08} // Netpune


var ssRad = []float64{ // Miles
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
	"sun.png",
	"mercury.png",
	"venus.png",
	"earth.png",
	"mars.png",
	"jupiter.png",
	"saturn.png",
	"uranus.png",
	"neptune.png"}

var showimages *bool = flag.Bool("i", true, "show images")
var canvas = svg.New(os.Stdout)

func vmap(value float64, low1 float64, high1 float64, low2 float64, high2 float64) float64 {
	return low2 + (high2-low2)*(value-low1)/(high1-low1)
}

func main() {

	width := 1300
	height := 200

	flag.Parse()
	canvas.Start(width, height)
	canvas.Title("Planets")
	canvas.Rect(0, 0, width, height, "fill:black")
	canvas.Gstyle("stroke:none")
	nobj := len(ssDist)
	y := height / 2
	margin := 100
	minsize := 7.0
	labeloc := height / 4

	var x, r, imScale, maxh float64
	var px, po int

	if *showimages {
		maxh = float64(height) / 4.0
	} else {
		maxh = float64(height) / minsize
	}
	for i := 1; i < nobj; i++ {
		x = vmap(ssDist[i], ssDist[1], ssDist[nobj-1], float64(margin), float64(width-margin))
		r = (vmap(ssRad[i], ssRad[1], ssRad[nobj-1], minsize, maxh)) / 2
		px = int(x)
		if *showimages {
			f, err := os.Open(ssImages[i])
			defer f.Close()
			if err != nil {
				println("bad image file:", ssImages[i])
				continue
			}
			p, perr := png.DecodeConfig(f)
			if perr != nil {
				println("bad decode:", ssImages[i])
				continue
			}
			imScale = r / float64(p.Width)
			hs := float64(p.Height) * imScale
			dy := y - (int(hs) / 2) // center the image
			po = int(r) / 2
			canvas.Image(px, dy, int(r), int(hs), ssImages[i])
		} else {
			po = 0
			canvas.Circle(px, y, int(r), "fill:#"+ssColor[i])
		}
		if ssDist[i] == 1.0 { // earth
			canvas.Line(px + po, y-po, px + po, y-labeloc, 
						"stroke-width:1px;stroke:white")
			canvas.Text(px + po, y-labeloc-10, "You are here", 
						"fill:white; font-size:14; font-family:Calibri; text-anchor:middle")
		}
	}
	canvas.Gend()
	canvas.End()
}
