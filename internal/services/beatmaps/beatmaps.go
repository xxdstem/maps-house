package beatmaps

import (
	"errors"
	"maps-house/pkg/logger"
	"os"
	"path/filepath"
	"strconv"
)

type Service interface {
	CheckBeatmapAvailability(setId int) error
}

type service struct {
	log          *logger.Logger
	db           DbRepository
	PriorityPath string
	MainPath     string
}

// errors
var (
	ErrorNotFoundDb   = errors.New("not found in db")
	ErrorNotFoundFile = errors.New("not found file")
)

func NewService(log *logger.Logger, db DbRepository, prior string, main string) Service {
	return &service{log: log, db: db, PriorityPath: prior, MainPath: main}
}

func (service *service) CheckBeatmapAvailability(setId int) error {
	bm, err := service.db.GetBeatmapsBySetId(setId)
	if err != nil {
		return err
	}
	if bm == nil {
		return ErrorNotFoundDb
	}
	filePath := filepath.Join(service.PriorityPath, strconv.Itoa(setId), "map.osz")
	if _, err = os.Stat(filePath); err == nil {
		return nil
	}
	return ErrorNotFoundFile
}
