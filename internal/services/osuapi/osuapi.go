package osuapi

import (
	"fmt"
	"github.com/bytedance/sonic"
	"io/ioutil"
	"maps-house/internal/entity"
	"maps-house/pkg/logger"
	"net/http"
)

type Service interface {
	GetBeatmapData(setId int) (*entity.BeatmapDTO, error)
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

func (s *service) GetBeatmapData(setId int) (*entity.BeatmapDTO, error) {
	var result entity.BeatmapDTO

	resp, err := http.Get(fmt.Sprintf(ApiGetBeatmaps+apiKey, setId))
	if err != nil {
		log.Error(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
	}
	err = sonic.Unmarshal(body, &result)

	return &result, err
}
