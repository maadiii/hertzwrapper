package server

import (
	"context"
	"net"

	"github.com/cloudwego/hertz/pkg/app"
)

type Context struct {
	context.Context //nolint
	Request         *app.RequestContext
}

func (ctx *Context) Host() []byte {
	return ctx.Request.Host()
}

func (ctx *Context) RemoteAddr() net.Addr {
	return ctx.Request.RemoteAddr()
}

func (ctx *Context) Path() []byte {
	return ctx.Request.URI().Path()
}

func (ctx *Context) FullPath() string {
	return ctx.Request.FullPath()
}

func (ctx *Context) Set(key string, value any) {
	ctx.Request.Set(key, value)
}

func (ctx *Context) Value(key any) any {
	return ctx.Request.Value(key)
}

func (ctx *Context) SetIdentity(identity Identity) {
	ctx.Request.Set(identityKey, identity)
}

func (ctx *Context) Identity() Identity {
	return ctx.Request.Value(identityKey).(Identity)
}

type Identity map[string]any

const identityKey = "identity"
