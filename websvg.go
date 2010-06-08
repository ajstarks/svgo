package main

import (
	"flag"
	"http"
	"log"
	"strings"
	"svg"
)

const defaultstyle = "fill:rgb(127,0,0);stroke:black"

var port = flag.String("port", ":2003", "http service address")

func main() {
	flag.Parse()
	http.Handle("/circle/", http.HandlerFunc(circle))
	http.Handle("/rect/", http.HandlerFunc(rect))
	http.Handle("/arc/", http.HandlerFunc(arc))
	http.Handle("/text/", http.HandlerFunc(text))
	err := http.ListenAndServe(*port, nil)
	if err != nil {
		log.Exit("ListenAndServe:", err)
	}
}

func shapestyle(path string) string {
	i := strings.LastIndex(path, "/") + 1
	if i > 0 && len(path[i:]) > 0 {
		return "fill:" + path[i:]
	}
	return defaultstyle
}

func circle(c *http.Conn, req *http.Request) {
	c.SetHeader("Content-Type", "image/svg+xml")
	s := svg.New(c)
	s.Start(500, 500)
	s.Title("Circle")
	s.Circle(250, 250, 125, shapestyle(c.Req.URL.Path))
	s.End()
}

func rect(c *http.Conn, req *http.Request) {
	c.SetHeader("Content-Type", "image/svg+xml")
	s := svg.New(c)
	s.Start(500, 500)
	s.Title("Rectangle")
	s.Rect(250, 250, 100, 200, shapestyle(c.Req.URL.Path))
	s.End()
}

func arc(c *http.Conn, req *http.Request) {
	c.SetHeader("Content-Type", "image/svg+xml")
	s := svg.New(c)
	s.Start(500, 500)
	s.Title("Arc")
	s.Arc(250, 250, 100, 100, 0, false, false, 100, 125, shapestyle(c.Req.URL.Path))
	s.End()
}

func text(c *http.Conn, req *http.Request) {
	c.SetHeader("Content-Type", "image/svg+xml")
	s := svg.New(c)
	s.Start(500, 500)
	s.Title("Text")
	s.Text(250, 250, "Hello, world", "text-anchor:middle;font-size:32px;"+shapestyle(c.Req.URL.Path))
	s.End()
}
