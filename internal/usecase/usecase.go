package usecase

import (
	"maps-house/internal/entity"
	"maps-house/internal/services/beatmaps"
	"maps-house/pkg/logger"
)

type usecase struct {
	osuApiService   OsuApiService
	beatmapsService BeatmapsService
	db              beatmaps.DbRepository
}

var log *logger.Logger

func New(l *logger.Logger, db beatmaps.DbRepository, osuApi OsuApiService, beatmaps BeatmapsService) *usecase {
	log = l
	return &usecase{db: db, osuApiService: osuApi, beatmapsService: beatmaps}
}

func (uc *usecase) GetBeatmapBySetId(setId int) (*entity.BeatmapsResult, error) {
	bm, err := uc.db.GetBeatmapsBySetId(setId)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &entity.BeatmapsResult{Result: bm}, nil
}

func (uc *usecase) CheckBeatmapAvailability(setId int) error {
	return uc.beatmapsService.CheckBeatmapAvailability(setId)
}

func (uc *usecase) DownloadMap(setId int) (*entity.BeatmapFile, error) {
	// Downloading map

	// Пишу пока всё здесь, потом разбросаю код

	//then
	return uc.ServeBeatmap(setId)
}

func (uc *usecase) ServeBeatmap(setId int) (*entity.BeatmapFile, error) {
	return nil, nil
}
