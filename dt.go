package main

import (
	"./svg"
)

var width = 500
var height = 500


func main() {

	svg.Start(width, height)
	svg.Def()
	svg.Gid("foo")
	svg.Circle(0, 0, 50, `id="mycircle"`)
	svg.Circle(0, 0, 25, "")
	svg.Gend()
	svg.DefEnd()
	svg.Use(200, 300, "#mycircle")
	svg.Use(200, 100, "#foo")
	svg.Use(200, 200, "/Users/ajstarks/svg/mrklogo.svg#mrklogo")

	svg.Grid(0, 0, width, height, 10, "stroke:black; opacity:0.2")
	svg.End()

}
