#A Go library for SVG generation#

The library generates SVG as defined by the Scalable Vector Graphics 1.1 Specification (<http://www.w3.org/TR/SVG11/>). 
Output goes to the specified io.Writer.

## Supported SVG elements ##

 circle, ellipse, polygon, polyline, rect (including roundrects), paths (arc,
 cubic and quadratic bezier paths), line, image, text, linearGradient, radialGradient

## Metadata elements ##

 desc, defs, g (style, transform, id), title, (a)ddress, link, use

## Building and Usage ##

See svgdef.[svg|png|pdf] for a graphical view of the function calls

Usage: 

	goinstall github.com/ajstarks/svgo
	
to install into your GO environment. 

a minimal program, to generate SVG to standard output.

	package main
	
	import (
		"github.com/ajstaks/svgo"
		"os"
	)
	
	var (
		width = 500
		height = 500
		canvas = svg.New(os.Stdout)
	)
	
	func main() {
		canvas.Start(width, height)
		canvas.Circle(width/2, height/2, 100)
		canvas.Text(width/2, height/2, "Hello, SVG", "text-anchor;font-size:30px;fill:white")
		canvas.End()
	}

Drawing in a web server: (http://localhost:2003/circle)

	package main
	
	import (
		"log"
		"github.com/ajstaks/svgo"
		"http"
	)
	
	func main() {
		http.Handle("/circle", http.HandlerFunc(circle))
		err := http.ListenAndServe(":2003", nil)
		if err != nil {
			log.Exit("ListenAndServe:", err)
		}
	}
	
	func circle(w http.ResponseWriter, req *http.Request) {
	  w.SetHeader("Content-Type", "image/svg+xml")
	  s := svg.New(w)
	  s.Start(500, 500)
	  s.Circle(250, 250, 125, "fill:none;stroke:black")
	  s.End()
	}

You may view the SVG output with a browser that supports SVG (tested on Chrome, Opera, Firefox and Safari), or any other SVG user-agent such as Batik Squiggle. The test-svgo script tries to use reasonable defaults
based on the GOOS and GOARCH environment variables.

The command:

	$ ./newsvg foo.go
   
creates Go source file ready for your code, using $EDITOR

To create browsable documentation:

	$ godoc -path=<svgo directory> -<http=:6060>
  
and click on the "Package documentation for svg" link

### Tutorial Video ###

A video describing how to use the package can be seen on YouTube at <http://www.youtube.com/watch?v=ze6O2Dj5gQ4> 

## Package contents ##

* svg.go:		Library
* test-svgo:	Compiles the library, builds the clients and displays the results
* newsvg:		Coding template command
* svgdef.go:	Creates a SVG representation of the API
* android.go:	The Android logo
* bubtrail.go: Bubble trails
* colortab.go: Display SVG named colors with RGB values
* flower.go:	Random "flowers"
* funnel.go:	Funnel from transparent circles
* gradient.go:	Linear and radial gradients
* imfade.go:	Show image fading
* lewitt.go:	Version of Sol Lewitt's Wall Drawing 91
* ltr.go:		Layer Tennis Remixes
* planets.go:	Show the scale of the Solar system
* pmap.go:		Proportion maps
* randcomp.go:	Compare random number generators
* richter.go:	Gerhard Richter's 256 colors
* rl.go:			Random lines (port of a Processing demo)
* svgopher.go:	SVGo Mascot
* vismem.go:	Visualize data from files
* websvg.go:	Generate SVG as a web server
* images/*		Images used by the client programs


## Functions ##

Many functions use x, y to specify an object's location, and w, h to specify the object's width and height.
Where applicable, a final optional argument specifies the style to be applied to the object. 
The style strings follow the SVG standard; name:value pairs delimited by semicolons, or a
series of name="value" pairs.
For example: "fill:none; opacity:0.3" or  fill="none", opacity="0.3" (see: <http://www.w3.org/TR/SVG11/styling.html>)


### Structure, Metadata and Links ###

`New(w io.Writer) *SVG`
  Constructor, Specify the output destination
  
`Start(w int, h int)`
  begin the SVG document with the width w and height h
  <http://www.w3.org/TR/SVG11/struct.html#SVGElement>

`End()`
  end the SVG document         

`Gstyle(s string)`
  begin a group, with the specified style
  <http://www.w3.org/TR/SVG11/struct.html#GElement>

`Gtransform(s string)`
  begin a group, with the specified transform

`Gid(s string)`
  begin a group, with the specified id

`Gend()`
  end the group (must be paired with Gstyle, Gtransform, Gid)

`Def()`
  begin a definition block
  <http://www.w3.org/TR/SVG11/struct.html#DefsElement>

`DefEnd()`
  end a definition block

`Desc(s string)`
  specify the text of the description
  <http://www.w3.org/TR/SVG11/struct.html#DescElement>

`Title(s string)`
  specify the text of the title
  <http://www.w3.org/TR/SVG11/struct.html#TitleElement>

`Link(name string, title string)`
  begin a link named "name", with the specified title
  <http://www.w3.org/TR/SVG11/linking.html#Links>

`LinkEnd()`
  end the link

`Use(x int, y int, link string, s ...string)`
  place the object referenced at link at the location x, y
  <http://www.w3.org/TR/SVG11/struct.html#UseElement>

### Shapes ###


 
`Circle(x int, y int, r int, s ...string)`
  draw a circle, centered at x,y with radius r
  <http://www.w3.org/TR/SVG11/shapes.html#CircleElement>
  
 
`Ellipse(x int, y int, w int, h int, s ...string)`
  draw an ellipse, centered at x,y with radii w, and h
  <http://www.w3.org/TR/SVG11/shapes.html#EllipseElement>


 
`Polygon(x []int, y []int, s ...string)`
  draw a series of line segments using an array of x, y coordinates
  <http://www.w3.org/TR/SVG11/shapes.html#PolygonElement>

 
`Rect(x int, y int, w int, h int, s ...string)`
  draw a rectangle with upper left-hand corner at x,y, with width w, and height h
  <http://www.w3.org/TR/SVG11/shapes.html#RectElement>

 
`Roundrect(x int, y int, w int, h int, rx int, ry int, s ...string)`
  draw a rounded rectangle with upper the left-hand corner at x,y, 
  with width w, and height h. The radii for the rounded portion 
  is specified by rx (width), and ry (height)
  
 
`Square(x int, y int, s int, style ...string)`
  draw a square with upper left corner at x,y with sides of length s

### Paths ###

 
`Arc(sx int, sy int, ax int, ay int, r int, large bool, sweep bool, ex int, ey int, s ...string)`
  draw an elliptical arc beginning coordinate at sx,sy, ending coordinate at ex, ey
  width and height of the arc are specified by ax, ay, the x axis rotation is r
  if sweep is true, then the arc will be drawn in a "positive-angle" direction (clockwise), if false,
  the arc is drawn counterclockwise.
  if large is true, the arc sweep angle is greater than or equal to 180 degrees, 
  otherwise the arc sweep is less than 180 degrees
  <http://www.w3.org/TR/SVG11/paths.html#PathDataEllipticalArcCommands>

 
`Bezier(sx int, sy int, cx int, cy int, px int, py int, ex int, ey int, s ...string)`
  draw a cubic bezier curve, beginning at sx,sy, ending at ex,ey
  with control points at cx,cy and px,py
  <http://www.w3.org/TR/SVG11/paths.html#PathDataCubicBezierCommands>

 
`Qbezier(sx int, sy int, cx int, cy int, ex int, ey int, tx int, ty int, s ...string)`
  draw a quadratic bezier curve, beginning at sx, sy, ending at tx,ty
  with control points are at cx,cy, ex,ey
  <http://www.w3.org/TR/SVG11/paths.html#PathDataQuadraticBezierCommands>

### Lines ###

 
`Line(x1 int, y1 int, x2 int, y2 int, s ...string)`
  draw a line segment between x1,y1 and x2,y2
  <http://www.w3.org/TR/SVG11/shapes.html#LineElement>

 
`Polyline(x []int, y []int, s ...string)`
  draw a polygon using coordinates specified in x,y arrays
  <http://www.w3.org/TR/SVG11/shapes.html#PolylineElement>

### Image and Text ###

 
`Image(x int, y int, w int, h int, link string, s ...string)`
  place at x,y (upper left hand corner), the image with width w, and height h, referenced at link
  <http://www.w3.org/TR/SVG11/struct.html#ImageElement>

`Text(x int, y int, t string, s ...string)`
  Place the specified text, t at x,y according to the style specified in s
  <http://www.w3.org/TR/SVG11/text.html#TextElement>

### Color ###

`RGB(r int, g int, b int) string` 
  creates a style string for the fill color designated 
  by the (r)ed, g(reen), (b)lue components
  <http://www.w3.org/TR/css3-color/>
  
`RGBA(r int, g int, b int, a float) string`
  as above, but includes the color's opacity as a value
  between 0.0 (fully transparent) and 1.0 (opaque)
  
### Gradients ###


`LinearGradient(id string, x1, y1, x2, y2 uint8, sc []Offcolor)`
  constructs a linear color gradient identified by id, 
  along the vector defined by (x1,y1), and (x2,y2).
  The stop color sequence defined in sc. Coordinates are expressed as percentages.
  <http://www.w3.org/TR/SVG11/pservers.html#LinearGradients>
  
  
`RadialGradient(id string, cx, cy, r, fx, fy uint8, sc []Offcolor)`
  constructs a radial color gradient identified by id, 
  centered at (cx,cy), with a radius of r.
  (fx, fy) define the location of the focal point of the light source. 
  The stop color sequence defined in sc.
  Coordinates are expressed as percentages.
  <http://www.w3.org/TR/SVG11/pservers.html#RadialGradients>

### Utility ###


`Grid(x int, y int, w int, h int, n int, s ...string)`
  draws a grid of straight lines starting at x,y, with a width w, and height h, and a size of n
