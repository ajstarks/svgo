package main

import (
	"os"
	"rand"
	"time"
	"fmt"
	"strconv"
	"./svg"
)

func main() {
	width := 512
	height := 256
	var n int = 256
	var rx, ry int

	if len(os.Args) > 1 {
		n, _ = strconv.Atoi(os.Args[1])
	}

	f, _ := os.Open("/dev/urandom", os.O_RDONLY, 0)
	x := make([]byte, n)
	y := make([]byte, n)
	f.Read(x)
	f.Read(y)
	f.Close()

	rand.Seed(time.Nanoseconds() % 1e9)
	svg.Start(600, 400)
	svg.Title("Random Integer Comparison")
	svg.Desc("Comparison of Random integers: the random device & the Go rand package")
	svg.Rect(0, 0, width/2, height, "fill:white; stroke:gray")
	svg.Rect(width/2, 0, width/2, height, "fill:white; stroke:gray")

	svg.Desc("Left: Go rand package (red), Right: /dev/urandom")
	svg.Gstyle("stroke:none; fill-opacity:0.5")
	for i := 0; i < n; i++ {
		rx = rand.Intn(255)
		ry = rand.Intn(255)
		svg.Circle(rx, ry, 5, svg.RGB(127, 0, 0))
		svg.Circle(int(x[i])+255, int(y[i]), 5, "fill:black")
	}
	svg.Gend()

	svg.Desc("Legends")
	svg.Gstyle("text-anchor:middle; font-size:18; font-family:Calibri")
	svg.Text(128, 280, "Go rand package", "")
	svg.Text(384, 280, "/dev/urandom")
	svg.Text(256, 280, fmt.Sprintf("n=%d", n), "font-size:12")
	svg.Gend()
	svg.End()
}
