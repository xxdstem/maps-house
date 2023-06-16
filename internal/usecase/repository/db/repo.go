package db

import (
	"maps-house/internal/entity"

	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repo{db: db}
}

func (r *repo) GetBeatmapsBySetId(setId int) (*entity.BeatmapMeta, error) {

	var result *entity.BeatmapMeta
	err := r.db.Model(&entity.BeatmapMeta{}).Preload("Beatmaps").Where(&entity.BeatmapMeta{ID: setId}).Find(&result).Error
	if err != nil {
		return nil, err
	}
	if len(result.Beatmaps) == 0 {
		result = nil
	}
	return result, nil
}

func (r *repo) InsertBeatmapSet(meta *entity.BeatmapMeta) error {
	return r.db.Create(&meta).Error
}

func (r *repo) SetDownloadedStatus(setId int, state bool) error {
	return r.db.Save(&entity.BeatmapMeta{ID: setId, Downloaded: state}).Error
}
