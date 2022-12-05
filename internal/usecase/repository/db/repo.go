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

func (r *repo) GetBeatmapBySetId(setId int) (*entity.Beatmap, error) {
	var beatmap *entity.Beatmap

	r.db.Where(&entity.Beatmap{BeatmapsetId: setId}).First(&beatmap)
	if beatmap.BeatmapsetId == 0 {
		return nil, ERROR_NOT_FOUND
	}
	return beatmap, nil
}
