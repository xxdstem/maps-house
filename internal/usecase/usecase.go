package usecase

import (
	"errors"
	"maps-house/internal/entity"
	"maps-house/pkg/logger"
	"os"
	"path/filepath"
	"strconv"
)

type usecase struct {
	db    DbRepository
	paths paths
}

type paths struct {
	PriorityDir string
	MainDir     string
}

var log *logger.Logger

// errors
var (
	ERROR_NOT_FOUND_DB = errors.New("not found in db?")
	ERROR_NOT_FOUND    = errors.New("not found")
)

func New(l *logger.Logger, db DbRepository, prior string, main string) *usecase {
	log = l
	return &usecase{db: db, paths: paths{PriorityDir: prior, MainDir: main}}
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
	bm, err := uc.db.GetBeatmapsBySetId(setId)
	if err != nil {
		return err
	}
	if bm == nil {
		return ERROR_NOT_FOUND_DB
	}
	filePath := filepath.Join(uc.paths.PriorityDir, strconv.Itoa(setId), "map.osz")
	if _, err = os.Stat(filePath); err == nil {
		return nil
	}
	return ERROR_NOT_FOUND
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
