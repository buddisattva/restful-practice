package main

import (
	"controller"

	"github.com/zenazn/goji"
)

const (
	albumRoutePrefix = "/v1/albums"
)

func main() {
	albumController := new(controller.AlbumController)
	// GET /v1/albums
	goji.Get(albumRoutePrefix, albumController.Index)
	// GET /v1/albums/{{id}}
	goji.Get(albumRoutePrefix+"/:id", albumController.Show)

	misc := new(controller.Misc)
	// 404 handler
	goji.NotFound(misc.NotFound)

	goji.Serve()
}
