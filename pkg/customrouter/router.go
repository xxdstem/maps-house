package customrouter

import (
	"fmt"
	"time"

	"github.com/fasthttp/router"
	"github.com/savsgio/gotils/strconv"
	"github.com/valyala/fasthttp"
)

type myRouter struct {
	r *router.Router
}

func (c myRouter) Handler(ctx *fasthttp.RequestCtx) {
	begin := time.Now()
	header := &ctx.Response.Header
	method := strconv.B2S(ctx.Request.Header.Method())
	if method == fasthttp.MethodOptions {
		if strconv.B2S(header.Peek("Access-Control-Request-Method")) != "" {
			header.Set("Access-Control-Allow-Methods", strconv.B2S(header.Peek("Allow")))

		}
		// Adjust status code to 204
		ctx.Error(fasthttp.StatusMessage(fasthttp.StatusNoContent), fasthttp.StatusNoContent)
	}
	header.Set("Access-Control-Allow-Origin", "*")
	header.Set("Content-Type", "application/json")
	c.r.Handler(ctx)
	fmt.Printf("> Request end - time took: %s\n", time.Since(begin).String())
}

func NewRouter(r *router.Router) *myRouter {
	return &myRouter{r: r}
}
