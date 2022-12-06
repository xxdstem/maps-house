package beatmaps

import (
	"maps-house/internal/entity"

	"github.com/fasthttp/router"
)

type Handler interface {
	Register(router *router.Router)
}

type UseCase interface {
	DownloadMap(setId int) (*entity.BeatmapFile, error)
	CheckBeatmapAvailability(setId int) error
}
