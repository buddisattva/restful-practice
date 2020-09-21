package main

import (
	"controller"

	"github.com/zenazn/goji"
)

const (
	albumRoutePrefix = "/v1/albums"
)

func main() {
	album := new(controller.Album)
	// GET /v1/albums
	goji.Get(albumRoutePrefix, album.Index)
	// GET /v1/albums/{{id}}
	goji.Get(albumRoutePrefix+"/:id", album.Show)

	misc := new(controller.Misc)
	// FIXME: unused
	goji.Get("/", misc.Root)
	// 404 handler
	goji.NotFound(misc.NotFound)

	// Call Serve() at the bottom of your main() function, and it'll take
	// care of everything else for you, including binding to a socket (with
	// automatic support for systemd and Einhorn) and supporting graceful
	// shutdown on SIGINT. Serve() is appropriate for both development and
	// production.
	goji.Serve()
}
