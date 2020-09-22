package controller

import (
	"encoding/json"
	"gateway"
	"io"
	"net/http"

	"github.com/zenazn/goji/web"
)

type AlbumController struct{}

// GET /v1/albums
func (a AlbumController) Index(w http.ResponseWriter, r *http.Request) {
	albumGateway := new(gateway.AlbumGateway)
	albums := albumGateway.ListAll()
	json, err := json.Marshal(albums)
	if err != nil {
		err.Error()
	}

	io.WriteString(w, string(json))
}

// GET /v1/albums/{{id}}
func (a AlbumController) Show(c web.C, w http.ResponseWriter, r *http.Request) {
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
