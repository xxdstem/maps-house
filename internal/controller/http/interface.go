package http

import "maps-house/internal/entity"

type UseCase interface {
	GetBeatmapBySetId(setId int) (*entity.BeatmapsDto, error)
	DownloadMap(setId int) (*entity.BeatmapFile, error)
	CheckBeatmapAvailability(setId int) error
}
