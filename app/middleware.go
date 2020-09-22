package main

import (
	"net/http"
)

// ApplicationJSON sets the content-type of responses to application/json.
func ApplicationJSON(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
