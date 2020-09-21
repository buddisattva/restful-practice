package main

import (
	controllers "app/http/controllers"

	"github.com/zenazn/goji"
)

func main() {
	misc := new(controllers.Misc)

	// Add routes to the global handler
	goji.Get("/", misc.Root)

	// Use a custom 404 handler
	goji.NotFound(misc.NotFound)

	// Call Serve() at the bottom of your main() function, and it'll take
	// care of everything else for you, including binding to a socket (with
	// automatic support for systemd and Einhorn) and supporting graceful
	// shutdown on SIGINT. Serve() is appropriate for both development and
	// production.
	goji.Serve()
}
