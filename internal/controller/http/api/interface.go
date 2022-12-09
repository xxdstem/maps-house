package api

import (
	"maps-house/internal/entity"

	"github.com/fasthttp/router"
)

type Handler interface {
	Register(router *router.Router)
}

type UseCase interface {
	GetBeatmapBySetId(setId int) (*entity.BeatmapsResultDTO, error)
}
