package controller

import (
	"encoding/json"
	"gateway"
	"io"
	"net/http"
	"strconv"

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
	// TODO: validate input id. input id should be unsigned int
	id, err := strconv.Atoi(c.URLParams["id"])

	albumGateway := new(gateway.AlbumGateway)
	album := albumGateway.GetById(id)
	json, err := json.Marshal(album)
	if err != nil {
		err.Error()
	}

	io.WriteString(w, string(json))
}
