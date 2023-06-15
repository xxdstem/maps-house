package http

import (
	"maps-house/internal/controller/http/api"
	"maps-house/internal/controller/http/beatmaps"
	"maps-house/internal/entity"
	"maps-house/pkg/logger"

	"github.com/fasthttp/router"
)

type UseCase interface {
	GetBeatmapBySetId(setId int) (*entity.BeatmapsResultDTO, error)
	DownloadMap(setId int) (*entity.BeatmapFile, error)
}

func RegisterApi(r *router.Router, l *logger.Logger, uc UseCase) {
	apiHandler := api.New(uc, l)
	apiHandler.Register(r)
}

func RegisterMain(r *router.Router, l *logger.Logger, uc UseCase) {
	mainHandler := beatmaps.New(uc, l)
	mainHandler.Register(r)
}
