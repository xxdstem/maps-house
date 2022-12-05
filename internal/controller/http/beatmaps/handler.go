package beatmaps

import (
	"maps-house/pkg/logger"
	"strconv"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

const (
	downloadRoute = "/d/{ID}"
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
	setIdstr := ctx.UserValue("ID").(string)
	setId, _ := strconv.Atoi(setIdstr)
	if err := h.uc.CheckBeatmapAvailability(setId); err != nil {
		ctx.WriteString(err.Error())
		return
	}
	
}
