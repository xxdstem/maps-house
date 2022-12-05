package api

import (
	"mapsHouse/internal/usecase"
	"mapsHouse/pkg/logger"
	"strconv"

	"github.com/bytedance/sonic"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

const (
	beatmapGet = "/api/beatmap/{beatmapId}"
)

type handler struct {
	log *logger.Logger
	uc  usecase.UseCase
}

func New(l *logger.Logger, uc usecase.UseCase) Handler {
	return &handler{log: l, uc: uc}
}

func (h *handler) Register(router *router.Router) {
	router.GET(beatmapGet, h.GetBeatmap)
}

func (h *handler) GetBeatmap(ctx *fasthttp.RequestCtx) {
	setIdstr := ctx.UserValue("beatmapId").(string)
	setId, _ := strconv.Atoi(setIdstr)
	result, err := h.uc.GetBeatmapBySetId(setId)
	if err != nil {
		h.log.Error(err)
		ctx.WriteString(err.Error())
		return
	}
	rawBytes, err := sonic.Marshal(result)
	if err != nil {
		h.log.Error(err)
		ctx.WriteString(err.Error())
		return
	}
	ctx.Write(rawBytes)
}
