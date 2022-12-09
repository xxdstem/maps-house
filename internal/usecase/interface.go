package usecase

import "maps-house/internal/entity"

type BeatmapsService interface {
	CheckBeatmapAvailability(setId int) error
}

type OsuApiService interface {
	GetBeatmapData(setId int) ([]entity.BeatmapDTO, error)
}

type DbRepository interface {
	InsertBeatmapSet(meta *entity.BeatmapMeta) error
	GetBeatmapsBySetId(setId int) (*entity.BeatmapMeta, error)
}
