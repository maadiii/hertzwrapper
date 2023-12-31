package server

import (
	"context"
	"net"

	"github.com/cloudwego/hertz/pkg/app"
)

type Context struct {
	context.Context //nolint
	request         *app.RequestContext
}

func (ctx *Context) Host() []byte {
	return ctx.request.Host()
}

func (ctx *Context) RemoteAddr() net.Addr {
	return ctx.request.RemoteAddr()
}

func (ctx *Context) Path() []byte {
	return ctx.request.URI().Path()
}

func (ctx *Context) FullPath() string {
	return ctx.request.FullPath()
}
