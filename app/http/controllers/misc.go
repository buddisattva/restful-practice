package controllers

import (
	"io"
	"net/http"
)

type Misc struct{}

// Root route (GET "/"). Print a list of greets.
func (m Misc) Root(w http.ResponseWriter, r *http.Request) {
	// In the real world you'd probably use a template or something.
	io.WriteString(w, "Gritter\n======\n\n")
}

// NotFound is a 404 handler.
func (m Misc) NotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Umm... have you tried turning it off and on again?", 404)
}
