package controller

import (
	"encoding/json"
	"gateway"
	"io"
	"net/http"
	"strconv"

	"github.com/zenazn/goji/web"
)

var albumGateway *gateway.AlbumGateway

func init() {
	albumGateway = new(gateway.AlbumGateway)
}

type AlbumController struct{}

// Index lists all albums, GET /v1/albums
func (a AlbumController) Index(w http.ResponseWriter, r *http.Request) {
	albums := albumGateway.ListAll()
	json, err := json.Marshal(albums)
	if err != nil {
		panic(err.Error())
	}

	io.WriteString(w, string(json))
}

// Show shows an album having input ID, GET /v1/albums/{{id}}
func (a AlbumController) Show(c web.C, w http.ResponseWriter, r *http.Request) {
	// TODO: validate input id. input id should be an unsigned int
	id, err := strconv.Atoi(c.URLParams["id"])
	if err != nil {
		panic(err.Error())
	}

	album := albumGateway.GetById(id)
	json, err := json.Marshal(album)
	if err != nil {
		panic(err.Error())
	}

	io.WriteString(w, string(json))
}

// Destroy deletes an album idempotently
func (a AlbumController) Destroy(c web.C, w http.ResponseWriter, r *http.Request) {
	// TODO: validate input id. input id should be an unsigned int
	id, err := strconv.Atoi(c.URLParams["id"])
	if err != nil {
		panic(err.Error())
	}

	albumGateway.DeleteById(id)

	io.WriteString(w, "{}")
}
