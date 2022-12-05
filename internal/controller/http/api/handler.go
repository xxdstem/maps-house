package api

import (
	"maps-house/internal/usecase"
	"maps-house/pkg/logger"
	"strconv"

	"github.com/bytedance/sonic"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

const (
	beatmapGet = "/api/beatmap/{beatmapId}"
)

var log *logger.Logger

type handler struct {
	uc usecase.UseCase
}

func New(l *logger.Logger, uc usecase.UseCase) Handler {
	log = l
	return &handler{uc: uc}
}

func (h *handler) Register(router *router.Router) {
	router.GET(beatmapGet, h.GetBeatmap)
}

func (h *handler) GetBeatmap(ctx *fasthttp.RequestCtx) {
	setIdstr := ctx.UserValue("beatmapId").(string)
	setId, _ := strconv.Atoi(setIdstr)
	result, err := h.uc.GetBeatmapBySetId(setId)
	if err != nil {
		log.Error(err)
		ctx.WriteString(err.Error())
		return
	}
	rawBytes, err := sonic.Marshal(result)
	if err != nil {
		log.Error(err)
		ctx.WriteString(err.Error())
		return
	}
	ctx.Write(rawBytes)
}
