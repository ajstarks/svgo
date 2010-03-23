package main
import "./svg"

func main() {
 svg.Start(500, 500)
 svg.Circle(250, 250, 100)
 svg.Text(250,250,"Hello, SVG", "fill:white;text-anchor:middle;font-size:30")
 svg.End()
}
