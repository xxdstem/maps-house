package beatmaps

import (
	"errors"
	"fmt"
	"io"
	"maps-house/internal/entity"
	"maps-house/pkg/logger"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

// errors
var (
	ErrorNotFoundDb   = errors.New("not found in db")
	ErrorNotFoundFile = errors.New("not found file")
)

type Service interface {
	CheckBeatmapAvailability(setId int) error
}

type DbRepository interface {
	GetBeatmapsBySetId(setId int) (*entity.BeatmapMeta, error)
	SetDownloadedStatus(setId int, state bool) error
}

type service struct {
	db           DbRepository
	PriorityPath string
	MainPath     string
}

var log *logger.Logger

func NewService(l *logger.Logger, db DbRepository, prior string, main string) Service {
	log = l
	return &service{db: db, PriorityPath: prior, MainPath: main}
}

func (this *service) CheckBeatmapAvailability(setId int) error {
	bm, err := this.db.GetBeatmapsBySetId(setId)
	if err != nil {
		return err
	}
	if bm == nil {
		return ErrorNotFoundDb
	}
	if bm.Downloaded == false {
		return ErrorNotFoundFile
	}
	filePath := this.setIdToPath(setId)
	if _, err = os.Stat(filePath); err == nil {
		return nil
	}
	return ErrorNotFoundFile
}

func (this *service) SaveBeatmapFile(setId int) error {

	filePath := this.setIdToPath(setId)
	url := fmt.Sprintf("https://osu.ppy.sh/d/%dn", setId)
	// Create the file
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	err = this.db.SetDownloadedStatus(setId, true)
	if err != nil {
		log.Error("lol")
		log.Error(err)
	}
	return nil
}

func (this *service) ServeBeatmap(setId int) ([]byte, error) {
	filePath := this.setIdToPath(setId)
	file, err := os.Open(filePath)
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)
	_, err = file.Read(buffer)
	defer file.Close()
	return buffer, err
}

func (this *service) setIdToPath(setId int) string {
	return filepath.Join(this.MainPath, strconv.Itoa(setId), "map.osz")
}
