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

func (uc *usecase) GetBeatmapBySetId(setId int) (*entity.BeatmapsResultDTO, error) {
	bm, err := uc.db.GetBeatmapsBySetId(setId)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &entity.BeatmapsResultDTO{Result: bm}, nil
}

func (uc *usecase) CheckBeatmapAvailability(setId int) error {
	return uc.beatmapsService.CheckBeatmapAvailability(setId)
}

func (uc *usecase) CacheMapFromBancho(setId int) error {
	a, err := uc.osuApiService.GetBeatmapData(setId)
	if err != nil {
		return err
	}
	var b entity.BeatmapDTO = a[0]
	meta := entity.BeatmapMeta{
		SetID:  b.SetID,
		Artist: b.Artist,
		Title:  a[0].Title,
	}
	return err
}

func (uc *usecase) DownloadMap(setId int) (*entity.BeatmapFile, error) {
	// Downloading map

	// Пишу пока всё здесь, потом разбросаю код

	//then
	return uc.ServeBeatmap(setId)
}

func (uc *usecase) ServeBeatmap(setId int) (*entity.BeatmapFile, error) {
	log.Info(setId)
	return nil, nil
}
