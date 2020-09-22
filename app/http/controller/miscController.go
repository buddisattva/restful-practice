package controller

import (
	"net/http"
)

type Misc struct{}

// NotFound is a 404 handler.
func (m Misc) NotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Umm... have you tried turning it off and on again?", 404)
}
