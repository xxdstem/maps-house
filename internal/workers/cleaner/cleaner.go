package cleaner

import (
	"maps-house/internal/entity"
	"maps-house/pkg/logger"

	"time"

	"gorm.io/gorm"
)

type BeatmapsService interface {
	RemoveBeatmapFile(setId int) error
}

type cleaner struct {
	service BeatmapsService
}

var db *gorm.DB
var log *logger.Logger

func New(l *logger.Logger, _db *gorm.DB, _service BeatmapsService) *cleaner {
	db = _db
	log = l
	cleaner := &cleaner{service: _service}
	go cleaner.check()

	return cleaner
}

func (w *cleaner) check() {
	for {
		var result []*entity.BeatmapMeta
		err := db.Model(&entity.BeatmapMeta{}).Where("api_update < ?", time.Now().Unix()-86400*190).Find(&result).Error
		if err != nil {
			log.Error(err)
		} else {
			log.Info("first", len(result))
			for _, v := range result {
				err = db.Model(&[]entity.Beatmap{}).Where(&entity.Beatmap{BeatmapsetID: v.BeatmapsetID}).Find(&v.Beatmaps).Error
				if err != nil {
					continue
				}
				if len(v.Beatmaps) > 0 && v.Beatmaps[0].Ranked != 2 && v.Beatmaps[0].Ranked != 5 {
					setId := v.BeatmapsetID
					w.service.RemoveBeatmapFile(setId)
					db.Delete(&entity.Beatmap{}, setId)
					db.Delete(&entity.BeatmapMeta{}, setId)
				}
			}
		}
		time.Sleep(24 * time.Hour)
	}
}
