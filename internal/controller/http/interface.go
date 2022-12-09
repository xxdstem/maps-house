package http

import "maps-house/internal/entity"

type UseCase interface {
	GetBeatmapBySetId(setId int) (*entity.BeatmapsResultDTO, error)
	DownloadMap(setId int) (*entity.BeatmapFile, error)
}
