package http

import "maps-house/internal/entity"

type UseCase interface {
	GetBeatmapBySetId(setId int) (*entity.BeatmapsResult, error)
	DownloadMap(setId int) (*entity.BeatmapFile, error)
	CheckBeatmapAvailability(setId int) error
	CacheMapFromBancho(setId int) error
}
