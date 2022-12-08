package usecase

import "maps-house/internal/entity"

type BeatmapsService interface {
	CheckBeatmapAvailability(setId int) error
}

type OsuApiService interface {
	GetBeatmapData(setId int) ([]entity.BeatmapDTO, error)
}
