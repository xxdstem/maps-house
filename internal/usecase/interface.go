package usecase

import "mapsHouse/internal/entity"

type UseCase interface {
	GetBeatmapBySetId(setId int) (*entity.Beatmap, error)
}

type DbRepository interface {
	GetBeatmapBySetId(setId int) (*entity.Beatmap, error)
}
