// "Standard" main.go for a command to run an HTTP server
// which serves an svg document from svgHandler, 
// with any other files from the current directory (exposed at cssroot)
// +build !appengine

package main

import (
	"fmt"
	"net/http"
	"os"
	"github.com/ajstarks/svgo"
)

const (
	port = ":5050"
)
	
func main() {
	fmt.Fprintf(os.Stderr, "serving on: http://localhost%s%s\n", port, root)
	// svgHandler(os.Stdout)	// helpful debug
	
	http.HandleFunc(root, svg.HttpAdapter(svgHandler))
	http.Handle(cssroot, http.StripPrefix(cssroot, http.FileServer(http.Dir("."))))
	http.ListenAndServe(port, nil)
}
