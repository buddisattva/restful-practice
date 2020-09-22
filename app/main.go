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
	// POST /v1/albums
	goji.Post(albumRoutePrefix, albumController.Store)
	// GET /v1/albums
	goji.Get(albumRoutePrefix, albumController.Index)
	// GET /v1/albums/{{id}}
	goji.Get(albumRoutePrefix+"/:id", albumController.Show)
	// DELETE /v1/albums/{{id}}
	goji.Delete(albumRoutePrefix+"/:id", albumController.Destroy)

	// add "Content-Type: application/json" into responded headers
	goji.Use(ApplicationJSON)

	misc := new(controller.Misc)
	// 404 handler
	goji.NotFound(misc.NotFound)

	goji.Serve()
}
