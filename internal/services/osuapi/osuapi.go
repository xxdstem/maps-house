package osuapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"maps-house/internal/entity"
	"maps-house/pkg/logger"
	"net/http"
)

type Service interface {
	GetBeatmapData(setId int) ([]entity.BeatmapDTO, error)
}

// Probably we need here something?
type service struct {
}

var log *logger.Logger

var apiKey string

var (
	ApiGetBeatmaps = "https://osu.ppy.sh/api/get_beatmaps?s=%d&k="
)

func NewService(l *logger.Logger, apikey string) Service {
	apiKey = apikey
	log = l
	return &service{}
}

func (s *service) GetBeatmapData(setId int) ([]entity.BeatmapDTO, error) {
	var result []entity.BeatmapDTO
	var finalUrl string = fmt.Sprintf(ApiGetBeatmaps+apiKey, setId)
	resp, err := http.Get(finalUrl)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("ioutil err", err.Error())
		return nil, err
	}
	log.Info("fixing shit!")
	//body = jsonhelper.FixJsonNewLines(body)
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return result, err
}
