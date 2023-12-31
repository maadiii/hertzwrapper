package main

import (
	"github.com/maadiii/hertzwrapper/server"
)

func main() {
	server.Handle(Users)

	server.Run(
		server.WithAddress(":8080"),
	)
}

// @action /users/:id [POST] 200 application/json
func Users(ctx *server.Context, in *RequestGet) (out *ResponseGet, err error) {
	out = &ResponseGet{
		ID:     in.ID,
		Name:   "maadi",
		Family: "azizi",
	}

	return
}

type RequestGet struct {
	ID int `path:"id" json:"id" form:"id" query:"id" cookie:"id" header:"id"`
}

type ResponseGet struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Family string `json:"family"`
}
