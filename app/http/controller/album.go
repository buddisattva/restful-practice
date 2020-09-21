package controller

import (
	"io"
	"net/http"

	"github.com/zenazn/goji/web"
)

type Album struct{}

// GET /v1/albums
func (a Album) Index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "787878787878")
}

// GET /v1/albums/{{id}}
func (a Album) Show(c web.C, w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Gritter\n======\n\n")
	handle := c.URLParams["id"]
	// user, ok := Users[handle]
	// if !ok {
	// 	http.Error(w, http.StatusText(404), 404)
	// 	return
	// }

	// user.Write(w, handle)

	io.WriteString(w, "albumId:"+handle)
	// for i := len(Greets) - 1; i >= 0; i-- {
	// 	if Greets[i].User == handle {
	// 		Greets[i].Write(w)
	// 	}
	// }
}
