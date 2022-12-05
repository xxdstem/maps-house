package usecase

import (
	"mapsHouse/internal/entity"
	"mapsHouse/pkg/logger"
)

type _usecase struct {
	logger *logger.Logger
	db     DbRepository
}

func New(l *logger.Logger, db DbRepository) UseCase {
	return &_usecase{logger: l, db: db}
}

func (uc *_usecase) GetBeatmapBySetId(setId int) (*entity.Beatmap, error) {
	bm, err := uc.db.GetBeatmapBySetId(setId)
	if err != nil {
		uc.logger.Error(err)
		return nil, err
	}
	return bm, nil
}
