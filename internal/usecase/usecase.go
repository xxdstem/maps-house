package usecase

import (
	"errors"
	"fmt"
	"maps-house/internal/entity"
	"maps-house/pkg/logger"
	"os"
	"path/filepath"
)

var log *logger.Logger

type usecase struct {
	db    DbRepository
	paths paths
}

type paths struct {
	PriorityDir string
	MainDir     string
}

func New(l *logger.Logger, db DbRepository, prior string, main string) *usecase {
	log = l
	return &usecase{db: db, paths: paths{PriorityDir: prior, MainDir: main}}
}

func (uc *usecase) GetBeatmapBySetId(setId int) (*entity.BeatmapsDto, error) {
	bm, err := uc.db.GetBeatmapsBySetId(setId)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &entity.BeatmapsDto{Result: bm}, nil
}

func (uc *usecase) CheckBeatmapAvailability(setId int) error {
	bm, err := uc.db.GetBeatmapsBySetId(setId)
	if err != nil {
		return err
	}
	if bm == nil {
		return errors.New("beatmap not found in db?")
	}
	filePath := filepath.Join(uc.paths.PriorityDir, string(setId), fmt.Sprintf("%d/map.osz", setId))
	if _, err = os.Stat(filePath); err == nil {
		return nil
	}

}

func (uc *usecase) DownloadMap(setId int) (*entity.BeatmapFile, error) {
	return nil, nil
}
