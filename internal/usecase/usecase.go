package usecase

import (
	"maps-house/internal/entity"
	"maps-house/pkg/logger"
)

var log *logger.Logger

type usecase struct {
	db DbRepository
}

func New(l *logger.Logger, db DbRepository) *usecase {
	log = l
	return &usecase{db: db}
}

func (uc *usecase) GetBeatmapBySetId(setId int) (*entity.BeatmapsDto, error) {
	bm, err := uc.db.GetBeatmapsBySetId(setId)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &entity.BeatmapsDto{Result: bm}, nil
}

func (uc *usecase) DownloadMap(setId int) (*entity.BeatmapFile, error) {
	return nil, nil
}
