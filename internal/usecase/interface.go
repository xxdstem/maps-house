package usecase

import "maps-house/internal/entity"

type DbRepository interface {
	GetBeatmapsBySetId(setId int) (*entity.BeatmapMeta, error)
}
