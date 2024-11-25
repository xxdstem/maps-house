package db

import (
	"maps-house/internal/entity"
	"time"

	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *repo {
	return &repo{db: db}
}

func (r *repo) GetBeatmapsBySetId(setId int) (*entity.BeatmapMeta, error) {

	var result *entity.BeatmapMeta
	err := r.db.Model(&entity.BeatmapMeta{}).Where(&entity.BeatmapMeta{BeatmapsetID: setId}).Find(&result).Error
	if err != nil {
		return nil, err
	}
	if result.BeatmapsetID == 0 {
		return nil, nil
	}
	if len(result.Beatmaps) == 0 {
		err = r.db.Model(&[]entity.Beatmap{}).Where(&entity.Beatmap{BeatmapsetID: setId}).Find(&result.Beatmaps).Error
		if err != nil || len(result.Beatmaps) == 0 {
			r.DeleteBeatmapSet(setId)
			result = nil
		}
	}
	return result, nil
}

func (r *repo) InsertBeatmapSet(meta *entity.BeatmapMeta) error {
	return r.db.Create(&meta).Error
}

func (r *repo) UpdateBeatmapSet(meta *entity.BeatmapMeta) error {
	// Need testing
	// probably need to remove beatmaps before insertting
	// cuz of some diffs could be removed and "Save" probably will just keep them.

	return r.db.Save(&meta).Error
}

func (r *repo) DeleteBeatmapSet(setId int) error {
	r.db.Delete(&entity.Beatmap{}, setId)
	return r.db.Delete(&entity.BeatmapMeta{}, setId).Error
}

func (r *repo) SetDownloadedStatus(setId int, state bool) error {
	meta := &entity.BeatmapMeta{BeatmapsetID: setId}
	return r.db.Model(&meta).Update("Downloaded", state).Update("LatestFetch", time.Now().Unix()).Error
}
