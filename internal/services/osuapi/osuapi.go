package osuapi

import (
	"github.com/bytedance/sonic"
	"io/ioutil"
	"log"
	"maps-house/internal/entity"
	"net/http"
)

type Service interface {
	GetBeatmapData(setId int) (*entity.BeatmapDTO, error)
}

// Probably we need here something?
type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) GetBeatmapData(setId int) (*entity.BeatmapDTO, error) {
	var result entity.BeatmapDTO
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	err = sonic.Unmarshal(body, &result)

	return &result, err
}
