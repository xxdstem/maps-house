package api

import (
	"maps-house/internal/entity"
	"maps-house/pkg/logger"
	"strconv"

	"github.com/bytedance/sonic"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

const (
	beatmapGet = "/api/beatmap/{ID}"
)

type Handler interface {
	Register(router *router.Router)
}

type UseCase interface {
	GetBeatmapBySetId(setId int) (*entity.BeatmapsResultDTO, error)
}

type handler struct {
	uc UseCase
}

var log *logger.Logger

func New(uc UseCase, l *logger.Logger) Handler {
	log = l
	return &handler{uc: uc}
}

func (h *handler) Register(router *router.Router) {
	router.GET(beatmapGet, h.GetBeatmap)
}

func (h *handler) GetBeatmap(ctx *fasthttp.RequestCtx) {
	setIdstr := ctx.UserValue("ID").(string)
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
		return
	}
	ctx.Write(rawBytes)
}
