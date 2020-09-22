package controller

import (
	"encoding/json"
	"gateway"
	"input"
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

// Store creates a new album
func (a AlbumController) Store(w http.ResponseWriter, r *http.Request) {
	var album input.NewAlbum

	// TODO: validate input format(maybe defined in another package)

	err := json.NewDecoder(r.Body).Decode(&album)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	albumGateway.Store(album)

	RenderJSON(http.StatusCreated, w, []byte{})
}

// Index lists all albums
func (a AlbumController) Index(w http.ResponseWriter, r *http.Request) {
	albums := albumGateway.ListAll()
	json, err := json.Marshal(albums)
	if err != nil {
		panic(err.Error())
	}

	io.WriteString(w, string(json))
}

// Show shows an album having input ID
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

// Update updates properties of an album
func (a AlbumController) Update(c web.C, w http.ResponseWriter, r *http.Request) {
	// TODO: validate input id. input id should be an unsigned int
	id, err := strconv.Atoi(c.URLParams["id"])
	if err != nil {
		panic(err.Error())
	}

	var album input.NewAlbum

	err = json.NewDecoder(r.Body).Decode(&album)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := albumGateway.UpdateById(id, album)
	if res {
		io.WriteString(w, "{}")
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	io.WriteString(w, "{}")
}
