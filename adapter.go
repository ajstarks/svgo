// Package svg provides an API for generating Scalable Vector Graphics (SVG)
package svg

import (
	"io"
	"net/http"
)

// HttpAdapter wraps an f(io.Writer) function to become an http.HandlerFunc  
// function, and also sets the correct HTTP header SVG mime type for the browser
func HttpAdapter(f func(w io.Writer)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		f(w)
	}
}
