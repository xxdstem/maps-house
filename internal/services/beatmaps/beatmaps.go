package beatmaps

import (
	"errors"
	"fmt"
	"io/ioutil"
	"maps-house/internal/entity"
	"maps-house/pkg/logger"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// errors
var (
	ErrorNotFoundDb   = errors.New("not found in db")
	ErrorNotFoundFile = errors.New("not found file")
)

type service struct {
	PriorityPath string
	MainPath     string
}

var log *logger.Logger

func NewService(l *logger.Logger, prior string, main string) *service {
	log = l
	return &service{PriorityPath: prior, MainPath: main}
}

func (this *service) CheckBeatmapAvailability(bm *entity.BeatmapMeta) error {
	if bm == nil {
		return ErrorNotFoundDb
	}
	if bm.Downloaded == false {
		return ErrorNotFoundFile
	}
	filePath := this.setIdToPath(bm.BeatmapsetID)
	if stat, err := os.Stat(filePath); err == nil {
		if stat.Size() < 100*1024 {
			os.Remove(filePath)
			return ErrorNotFoundFile
		}
		return nil
	}
	return ErrorNotFoundFile
}

func (this *service) SaveBeatmapFile(setId int, chimu bool) error {
	user := "Karanos"
	pw := "3ab61ccb3678229a797bf4e48fb96f90"

	filePath := this.setIdToPath(setId)
	var url string
	if chimu {
		url = fmt.Sprintf("https://chimu.moe/d/%dn", setId)
	} else {
		url = fmt.Sprintf("https://osu.ppy.sh/d/%dn?u=%s&h=%s", setId, user, pw)
	}
	// Create the file
	dir := filepath.Dir(filePath)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}
	_, err = os.Stat(filePath)
	if !os.IsNotExist(err) {
		os.Remove(filePath)
	}
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
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	// Writer the body to file
	if len(body) < 100*1024 {
		if chimu {
			return errors.New("fuck")
		}
		return this.SaveBeatmapFile(setId, true)
	}

	_, err = out.Write(body)
	if err != nil {
		return err
	}
	return nil
}

func (this *service) RemoveBeatmapFile(setId int) error {
	filePath := this.setIdToPath(setId)
	// Create the file
	_, err := os.Stat(filePath)
	if !os.IsNotExist(err) {
		return os.Remove(filePath)
	}
	return nil
}

func (this *service) ServeBeatmap(setId int) ([]byte, error) {
	filePath := this.setIdToPath(setId)
	file, err := os.Open(filePath)
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)
	_, err = file.Read(buffer)
	file.Close()
	return buffer, err
}

func (this *service) setIdToPath(setId int) string {
	return filepath.Join(this.MainPath, strconv.Itoa(setId), "map.osz")
}

func (*service) CheckUpdateConditions(bm *entity.BeatmapMeta) bool {
	return bm.Beatmaps[0].Ranked != 2 &&
		time.Since(time.Unix(bm.ApiUpdate, 0))/24 > 3
}
