package handler

import (
	"github.com/maadiii/hertzwrapper/server"
)

// @action /users/:id [GET] 200 index.html
func Users(_ *server.Context, in *RequestGet) (out *ResponseGet, err error) {
	out = &ResponseGet{
		ID:     in.ID,
		Name:   "Maadi",
		Family: "Azizi",
	}

	return
}

type RequestGet struct {
	ID int `path:"id" json:"id" form:"id" query:"id" cookie:"id" header:"id"` //nolint
}

type ResponseGet struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Family string `json:"family"`
}
