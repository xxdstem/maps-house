package db

import (
	"maps-house/internal/entity"

	"errors"

	"gorm.io/gorm"
)

// errors
var (
	ERROR_NOT_FOUND = errors.New("not found")
)

type repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repo{db: db}
}

func (r *repo) GetBeatmapsBySetId(setId int) (*entity.BeatmapMeta, error) {

	var result *entity.BeatmapMeta
	err := r.db.Model(&entity.BeatmapMeta{}).Preload("Beatmaps").Where(&entity.BeatmapMeta{BeatmapsetId: setId}).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
