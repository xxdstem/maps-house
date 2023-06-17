package beatmaps

import (
	"fmt"
	"maps-house/internal/entity"
	"maps-house/pkg/logger"
	"strconv"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

const (
	downloadRoute = "/d/{ID}"
)

type Handler interface {
	Register(router *router.Router)
}

type UseCase interface {
	DownloadMap(setId int) (*entity.BeatmapFile, error)
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
	router.GET(downloadRoute, h.DownloadMap)
}

func (h *handler) DownloadMap(ctx *fasthttp.RequestCtx) {
	setIdstr := ctx.UserValue("ID").(string)
	setId, _ := strconv.Atoi(setIdstr)
	beatmapFile, err := h.uc.DownloadMap(setId)
	if err != nil {
		ctx.WriteString(err.Error())
		return
	}

	ctx.Write(beatmapFile.Body)
	ctx.Response.Header.Add("Content-type", "application/octet-stream")
	ctx.Response.Header.Add("Content-length", fmt.Sprintf("%d", len(beatmapFile.Body)))
	ctx.Response.Header.Add("Content-Description", "File Transfer")
	ctx.Response.Header.Add("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.osz\"", beatmapFile.Title))
}
