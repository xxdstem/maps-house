package api

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type Handler interface {
	Register(router *router.Router)
	GetBeatmap(ctx *fasthttp.RequestCtx)
}
