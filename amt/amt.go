package main

import (
	"os"

	svg "github.com/wildberries-ru/svgo"
)

func main() {
	canvas := svg.New(os.Stdout)
	width, height := 500, 500
	duration, repeat := 1.0, 0

	canvas.Start(width, height)
	canvas.Rect(0, 0, width, height)

	// Translate
	canvas.CenterRect(0, 0, 40, 30, "fill:red", `id="redbox"`)
	canvas.CenterRect(0, 0, 40, 30, "fill:blue", `id="bluebox"`)
	canvas.AnimateTranslate("#redbox", 100, 100, 200, 200, duration, repeat)
	canvas.AnimateTranslate("#bluebox", 200, 200, 100, 100, duration, repeat)

	// Scale
	canvas.Translate(200, 100)
	canvas.CenterRect(0, 0, 40, 30, "fill:green", `id="greenbox"`)
	canvas.Gend()
	canvas.AnimateScale("#greenbox", 1, 3, duration, repeat)

	// SkewX
	canvas.Translate(300, 100)
	canvas.CenterRect(0, 0, 40, 30, "fill:purple", `id="purplebox"`)
	canvas.Gend()
	canvas.AnimateSkewX("#purplebox", 0, 45, duration, repeat)

	// SkewY
	canvas.Translate(400, 100)
	canvas.CenterRect(0, 0, 40, 30, "fill:lightsteelblue", `id="lsbox"`)
	canvas.Gend()
	canvas.AnimateSkewY("#lsbox", 0, 45, duration, repeat)

	// Rotate
	canvas.Translate(width/4, height/2)
	canvas.CenterRect(0, 0, 40, 30, "fill:maroon", `id="rotbox"`)
	canvas.Gend()
	canvas.AnimateRotate("#rotbox", 0, 75, 75, 360, 75, 75, duration*2, repeat)

	canvas.End()

}
