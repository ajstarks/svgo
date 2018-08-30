// personal: make persona slides
package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"os"

	"github.com/ajstarks/deck/generate"
)

// Persona is a title and 1 to n persons
type Persona struct {
	Title  string   `xml:"title,attr"`
	Person []Person `xml:"person"`
}

// Person describes the attributes of a persona
type Person struct {
	Name     string     `xml:"name"`
	Role     string     `xml:"role"`
	Picture  pic        `xml:"picture"`
	Location string     `xml:"location"`
	About    string     `xml:"about"`
	Lists    []list     `xml:"list"`
	Scores   []scoreset `xml:"scoreset"`
	Summary  string     `xml:"summary"`
}

type score struct {
	LName string  `xml:"lname,attr"`
	RName string  `xml:"rname,attr"`
	Score float64 `xml:"score,attr"`
}

type scoreset struct {
	Name   string  `xml:"name,attr"`
	Scores []score `xml:"score"`
}

type list struct {
	Name  string   `xml:"name,attr"`
	Items []string `xml:"item"`
}

type pic struct {
	Name   string `xml:"name,attr"`
	Width  int    `xml:"width,attr"`
	Height int    `xml:"height,attr"`
}

const (
	scorewidth      = 25.0
	scoreheight     = 2.5
	ptop            = 89.0
	left1           = 15.0
	left2           = 30.0
	scorecolor      = "rgb(230,230,230)"
	scorelabelcolor = "rgb(75,75,75)"
)

// makeslides reads and decode the persona file, and makes slides
func makeslides(filename, color string, deck *generate.Deck) {
	r, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	defer r.Close()
	var p Persona
	if err := xml.NewDecoder(r).Decode(&p); err == nil {
		pslide(&p, deck, color)
	} else {
		fmt.Fprintf(os.Stderr, "Unable to parse personas (%v)\n", err)
	}
}

// pslide makes a persona slide
func pslide(p *Persona, deck *generate.Deck, color string) {
	for _, person := range p.Person {
		deck.StartSlide()
		deck.TextMid(left1, 95, p.Title, "sans", 2, "darkgray")
		deck.TextMid(left1, ptop, person.Name, "sans", 2.5, "black")
		deck.Image(left1, 72, person.Picture.Width, person.Picture.Height, person.Picture.Name, "")
		deck.Rect(left1, 42, 20, 30, color)
		deck.TextEnd(11.5, 53, "ROLE", "sans", 1.2, "white")
		deck.Text(12, 53, person.Role, "sans", 1.2, "white")
		deck.TextEnd(11.5, 50, "LOCATION", "sans", 1.2, "white")
		deck.Text(12, 50, person.Location, "sans", 1.2, "white")
		deck.TextBlock(6, 42, person.About, "sans", 1.2, 15, "white")
		deck.Text(left2, 55, "SUMMARY", "sans", 1.5, color)
		deck.Rect(66, 30, 75, 60, color, 7)
		deck.TextBlock(left2, 51, person.Summary, "sans", 1.2, 60, "black")
		measures(person.Scores, deck, color)
		lists(person.Lists, deck, color)
		deck.EndSlide()
	}
}

// list displays list data
func lists(pl []list, deck *generate.Deck, color string) {
	var lx, ly float64
	for _, l := range pl {
		switch l.Name {
		case "GOALS/LIKES":
			lx = left2
			ly = ptop
		case "FRUSTRATIONS":
			lx = left2
			ly = 72
		case "RELATED QUESTIONS":
			lx = left2
			ly = 25
		}
		deck.Text(lx, ly, l.Name, "sans", 1.5, color)
		ly -= 1.5 * scoreheight
		if len(l.Items) < 7 {
			deck.List(lx, ly, 1.1, 1.4, 40, l.Items, "bullet", "sans", "black")
		} else {
			deck.List(lx, ly, 1.1, 1.4, 40, l.Items[0:6], "bullet", "sans", "black")
			deck.List(lx+32, ly, 1.1, 1.4, 40, l.Items[6:], "bullet", "sans", "black")
		}
	}
}

// measures places a percentage indicator in a rectangle
func measures(ps []scoreset, deck *generate.Deck, color string) {
	var sx, sy float64
	for _, s := range ps {
		switch s.Name {
		case "TECHNOLOGY":
			sx = left1
			sy = 20
		case "PERSONALITY":
			sx = 85
			sy = ptop
		}
		deck.Text(sx-scorewidth/2, sy, s.Name, "sans", 1.5, color)
		sy -= 2.0 * scoreheight
		for _, ss := range s.Scores {
			makescore(sx, sy, ss, deck, color)
			sy -= 2.5 * scoreheight
		}
	}
}

// make score builds the rect and circle to indicate a percentage score
func makescore(x, y float64, s score, deck *generate.Deck, color string) {
	deck.Text(x-scorewidth/2, y+2, s.LName, "sans", 1, scorelabelcolor)
	if len(s.RName) > 0 {
		deck.TextEnd(x+scorewidth/2, y+2, s.RName, "sans", 1, scorelabelcolor)
	}
	deck.Rect(x, y, scorewidth, scoreheight, scorecolor)
	deck.Circle((x-scorewidth/2)+(scorewidth*(s.Score/100)), y, scoreheight/2, color)
}

func main() {
	hcolor := flag.String("h", "rgb(67,74,154)", "highlight color")
	flag.Parse()
	deck := generate.NewSlides(os.Stdout, 1600, 900)
	deck.StartDeck()
	for _, f := range flag.Args() {
		makeslides(f, *hcolor, deck)
	}
	deck.EndDeck()
}
