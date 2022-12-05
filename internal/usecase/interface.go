package usecase

import "maps-house/internal/entity"

type UseCase interface {
	GetBeatmapBySetId(setId int) (*entity.Beatmap, error)
}

type DbRepository interface {
	GetBeatmapBySetId(setId int) (*entity.Beatmap, error)
}
