package main

import (
	"github.com/maadiii/hertzwrapper/example/handler"
	"github.com/maadiii/hertzwrapper/server"
)

func main() {
	server.Handle(handler.Users)

	hertz := server.Hertz(
		true,
		server.WithHostPorts(":8080"),
	)
	hertz.LoadHTMLGlob("example/views/index.html")

	hertz.Spin()
}
