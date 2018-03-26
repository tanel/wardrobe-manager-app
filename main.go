package main

import (
	"github.com/tanel/wardrobe-organizer/router"
	"github.com/tanel/webapp/server"
)

func main() {
	server.Serve("wardrobe", router.New())
}
