#A Go library for SVG generation#

The library generates SVG as defined by the Scalable Vector Graphics 1.1 Specification (<http://www.w3.org/TR/SVG11/>). 
Output goes to the specified io.Writer.

## Supported SVG elements ##

 circle, ellipse, polygon, polyline, rect (including roundrects), paths (general, arc,
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
		"github.com/ajstarks/svgo"
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
		canvas.Text(width/2, height/2, "Hello, SVG", "text-anchor:middle;font-size:30px;fill:white")
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
* f50.go:		Get 50 photos from Flickr based on a query
* funnel.go:	Funnel from transparent circles
* gradient.go:	Linear and radial gradients
* imfade.go:	Show image fading
* lewitt.go:	Version of Sol Lewitt's Wall Drawing 91
* ltr.go:		Layer Tennis Remixes
* paths.go		Demonstrate SVG paths
* planets.go:	Show the scale of the Solar system
* pmap.go:		Proportion maps
* randcomp.go:	Compare random number generators
* richter.go:	Gerhard Richter's 256 colors
* rl.go:			Random lines (port of a Processing demo)
* svgopher.go:	SVGo Mascot
* vismem.go:	Visualize data from files
* websvg.go:	Generate SVG as a web server
* images/*		Images used by the client programs


## Functions and types ##

Many functions use x, y to specify an object's location, and w, h to specify the object's width and height.
Where applicable, a final optional argument specifies the style to be applied to the object. 
The style strings follow the SVG standard; name:value pairs delimited by semicolons, or a
series of name="value" pairs. For example: `"fill:none; opacity:0.3"` or  `fill="none" opacity="0.3"` (see: <http://www.w3.org/TR/SVG11/styling.html>)

The Offcolor type:

	type Offcolor struct {
		Offset  uint8
		Color   string
		Opacity float
	}

is used to specify the offset, color, and opacity of stop colors in linear and radial gradients


### Structure, Metadata and Links ###

	New(w io.Writer) *SVG
  Constructor, Specify the output destination
  
	Start(w int, h int)
  begin the SVG document with the width w and height h
  <http://www.w3.org/TR/SVG11/struct.html#SVGElement>
  
	Startview(w, h, minx, miny, vh, vh)
  begin the SVG document with the width w, height h, with a viewBox at minx, miny, vw, vh
  <http://www.w3.org/TR/SVG11/struct.html#SVGElement>

	End()
  end the SVG document         

	Gstyle(s string)
  begin a group, with the specified style
  <http://www.w3.org/TR/SVG11/struct.html#GElement>

	Gtransform(s string)
  begin a group, with the specified transform

	Gid(s string)
  begin a group, with the specified id

	Gend()
  end the group (must be paired with Gstyle, Gtransform, Gid)

	Def()
  begin a definition block
  <http://www.w3.org/TR/SVG11/struct.html#DefsElement>

	DefEnd()
  end a definition block

	Desc(s string)
  specify the text of the description
  <http://www.w3.org/TR/SVG11/struct.html#DescElement>

	Title(s string)
  specify the text of the title
  <http://www.w3.org/TR/SVG11/struct.html#TitleElement>

	Link(name string, title string)
  begin a link named "name", with the specified title
  <http://www.w3.org/TR/SVG11/linking.html#Links>

	LinkEnd()
  end the link

	Use(x int, y int, link string, s ...string)
  place the object referenced at link at the location x, y
  <http://www.w3.org/TR/SVG11/struct.html#UseElement>

### Shapes ###

![Circle](http://farm5.static.flickr.com/4144/5187953823_01a1741489_m.jpg)
 
	Circle(x int, y int, r int, s ...string)
  draw a circle, centered at x,y with radius r
  <http://www.w3.org/TR/SVG11/shapes.html#CircleElement>
 
![Ellipse](http://farm2.static.flickr.com/1271/5187953773_a9d1fc406c_m.jpg)
 
	Ellipse(x int, y int, w int, h int, s ...string)
  draw an ellipse, centered at x,y with radii w, and h
  <http://www.w3.org/TR/SVG11/shapes.html#EllipseElement>

![Polygon](http://farm2.static.flickr.com/1006/5187953873_337dc26597_m.jpg)
 
	Polygon(x []int, y []int, s ...string)
  draw a series of line segments using an array of x, y coordinates
  <http://www.w3.org/TR/SVG11/shapes.html#PolygonElement>

![Rect](http://farm2.static.flickr.com/1233/5188556032_86c90e354b_m.jpg)
 
	Rect(x int, y int, w int, h int, s ...string)
  draw a rectangle with upper left-hand corner at x,y, with width w, and height h
  <http://www.w3.org/TR/SVG11/shapes.html#RectElement>

![Roundrect](http://farm2.static.flickr.com/1275/5188556120_e2a9998fee_m.jpg)
 
	Roundrect(x int, y int, w int, h int, rx int, ry int, s ...string)
  draw a rounded rectangle with upper the left-hand corner at x,y, 
  with width w, and height h. The radii for the rounded portion 
  is specified by rx (width), and ry (height)
  
![Square](http://farm5.static.flickr.com/4110/5187953659_54dcce242e_m.jpg)
 
	Square(x int, y int, s int, style ...string)
  draw a square with upper left corner at x,y with sides of length s

### Paths ###

	Path(p string, s ...style)
 draw the arbitrary path as specified in p, according to the style specified in s. <http://www.w3.org/TR/SVG11/paths.html>


 ![Arc](http://farm2.static.flickr.com/1300/5188556148_df1a176074_m.jpg)
 
	Arc(sx int, sy int, ax int, ay int, r int, large bool, sweep bool, ex int, ey int, s ...string)
  draw an elliptical arc beginning coordinate at sx,sy, ending coordinate at ex, ey
  width and height of the arc are specified by ax, ay, the x axis rotation is r
  if sweep is true, then the arc will be drawn in a "positive-angle" direction (clockwise), if false,
  the arc is drawn counterclockwise.
  if large is true, the arc sweep angle is greater than or equal to 180 degrees, 
  otherwise the arc sweep is less than 180 degrees
  <http://www.w3.org/TR/SVG11/paths.html#PathDataEllipticalArcCommands>

![Bezier](http://farm2.static.flickr.com/1233/5188556246_a03e67d013.jpg)
 
	Bezier(sx int, sy int, cx int, cy int, px int, py int, ex int, ey int, s ...string)
  draw a cubic bezier curve, beginning at sx,sy, ending at ex,ey
  with control points at cx,cy and px,py
  <http://www.w3.org/TR/SVG11/paths.html#PathDataCubicBezierCommands>

 ![Qbezier](http://farm2.static.flickr.com/1018/5187953917_9a43cf64fb.jpg)
 
	Qbezier(sx int, sy int, cx int, cy int, ex int, ey int, tx int, ty int, s ...string)
  draw a quadratic bezier curve, beginning at sx, sy, ending at tx,ty
  with control points are at cx,cy, ex,ey
  <http://www.w3.org/TR/SVG11/paths.html#PathDataQuadraticBezierCommands>
  
	Qbez(sx int, sy int, cx int, cy int, ex int, ey int, s...string)
   draws a quadratic bezier curver, with optional style beginning at sx,sy, ending at ex, sy
   with the control point at cx, cy
   <http://www.w3.org/TR/SVG11/paths.html#PathDataQuadraticBezierCommands>

### Lines ###

![Line](http://farm5.static.flickr.com/4154/5188556080_0be19da0bc.jpg)
 
	Line(x1 int, y1 int, x2 int, y2 int, s ...string)
  draw a line segment between x1,y1 and x2,y2
  <http://www.w3.org/TR/SVG11/shapes.html#LineElement>

![Polyline](http://farm2.static.flickr.com/1266/5188556384_a863273a69.jpg)
 
	Polyline(x []int, y []int, s ...string)
  draw a polygon using coordinates specified in x,y arrays
  <http://www.w3.org/TR/SVG11/shapes.html#PolylineElement>

### Image and Text ###

![Image](http://farm5.static.flickr.com/4058/5188556346_e5ce3dcbc2_m.jpg)

	Image(x int, y int, w int, h int, link string, s ...string)
  place at x,y (upper left hand corner), the image with width w, and height h, referenced at link
  <http://www.w3.org/TR/SVG11/struct.html#ImageElement>

	Text(x int, y int, t string, s ...string)
  Place the specified text, t at x,y according to the style specified in s
  <http://www.w3.org/TR/SVG11/text.html#TextElement>

### Color ###

	RGB(r int, g int, b int) string
  creates a style string for the fill color designated 
  by the (r)ed, g(reen), (b)lue components
  <http://www.w3.org/TR/css3-color/>
  
	RGBA(r int, g int, b int, a float) string
  as above, but includes the color's opacity as a value
  between 0.0 (fully transparent) and 1.0 (opaque)
  
### Gradients ###

![LinearGradient](http://farm5.static.flickr.com/4153/5187954033_3972f63fa9.jpg) 

	LinearGradient(id string, x1, y1, x2, y2 uint8, sc []Offcolor)
  constructs a linear color gradient identified by id, 
  along the vector defined by (x1,y1), and (x2,y2).
  The stop color sequence defined in sc. Coordinates are expressed as percentages.
  <http://www.w3.org/TR/SVG11/pservers.html#LinearGradients>

![RadialGradient](http://farm2.static.flickr.com/1302/5187954065_7ddba7b819.jpg)
  
	RadialGradient(id string, cx, cy, r, fx, fy uint8, sc []Offcolor)
  constructs a radial color gradient identified by id, 
  centered at (cx,cy), with a radius of r.
  (fx, fy) define the location of the focal point of the light source. 
  The stop color sequence defined in sc.
  Coordinates are expressed as percentages.
  <http://www.w3.org/TR/SVG11/pservers.html#RadialGradients>

### Utility ###

![Grid](http://farm5.static.flickr.com/4133/5190957924_7a31d0db34.jpg)

	Grid(x int, y int, w int, h int, n int, s ...string)
  draws a grid of straight lines starting at x,y, with a width w, and height h, and a size of n
