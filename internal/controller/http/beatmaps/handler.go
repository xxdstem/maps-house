package beatmaps

import (
	"maps-house/pkg/logger"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

const (
	downloadRoute = "/d/{beatmapId}"
)

var log *logger.Logger

type handler struct {
	uc UseCase
}

func New(uc UseCase, l *logger.Logger) Handler {
	log = l
	return &handler{uc: uc}
}

func (h *handler) Register(router *router.Router) {
	router.GET(downloadRoute, h.DownloadMap)
}

func (h *handler) DownloadMap(ctx *fasthttp.RequestCtx) {

}
