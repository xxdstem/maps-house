package usecase

import "maps-house/internal/entity"

type UseCase interface {
	GetBeatmapBySetId(setId int) (*entity.BeatmapsDto, error)
}

type DbRepository interface {
	GetBeatmapsBySetId(setId int) (*entity.BeatmapMeta, error)
}
