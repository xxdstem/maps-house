package usecase

import (
	"fmt"
	"maps-house/internal/entity"
	"maps-house/pkg/logger"
	"time"
)

type BeatmapsService interface {
	CheckBeatmapAvailability(setId int) (*entity.BeatmapMeta, error)
	SaveBeatmapFile(setId int) error
	ServeBeatmap(setId int) ([]byte, error)
}

type OsuApiService interface {
	GetBeatmapData(setId int) ([]entity.BeatmapDTO, error)
}

type DbRepository interface {
	InsertBeatmapSet(meta *entity.BeatmapMeta) error
	GetBeatmapsBySetId(setId int) (*entity.BeatmapMeta, error)
}

type usecase struct {
	osuApiService   OsuApiService
	beatmapsService BeatmapsService
	db              DbRepository
}

var log *logger.Logger

func New(l *logger.Logger, db DbRepository, osuApi OsuApiService, beatmaps BeatmapsService) *usecase {
	log = l
	return &usecase{db: db, osuApiService: osuApi, beatmapsService: beatmaps}
}

func (uc *usecase) GetBeatmapBySetId(setId int) (*entity.BeatmapsResultDTO, error) {
	bm, err := uc.db.GetBeatmapsBySetId(setId)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &entity.BeatmapsResultDTO{Result: bm}, nil
}

func (uc *usecase) SaveMapMeta(setId int) (*entity.BeatmapMeta, error) {
	maps, err := uc.osuApiService.GetBeatmapData(setId)
	if err != nil {
		return nil, err
	}
	var firstBeatmap entity.BeatmapDTO = maps[0]
	var curTime time.Time = time.Now()
	meta := entity.BeatmapMeta{
		BeatmapsetID: firstBeatmap.SetID,
		Artist:       firstBeatmap.Artist,
		Title:        firstBeatmap.Title,
		Creator:      firstBeatmap.Creator,
		Tags:         firstBeatmap.Tags,
		Length:       firstBeatmap.Length,
		BPM:          firstBeatmap.BPM,
		LanguageID:   firstBeatmap.LanguageID,
		GenreID:      firstBeatmap.GenreID,
		Downloaded:   false,
		Beatmaps:     []*entity.Beatmap{},
	}
	for i := range maps {
		var beatmapDto entity.BeatmapDTO = maps[i]
		t, err := time.Parse("2006-01-02 15:04:05", beatmapDto.LastUpdate)
		if err != nil {
			return nil, err
		}
		beatmap := &entity.Beatmap{
			ID:           beatmapDto.ID,
			BeatmapsetID: beatmapDto.SetID,
			MD5:          beatmapDto.MD5,
			Version:      beatmapDto.Version,
			HitLength:    beatmapDto.HitLength,
			Ranked:       beatmapDto.Ranked,
			LastUpdate:   t.Unix(),
			ApiUpdate:    curTime.Unix(),
			AR:           beatmapDto.AR,
			OD:           beatmapDto.OD,
			HP:           beatmapDto.HP,
			CS:           beatmapDto.CS,
		}
		meta.Beatmaps = append(meta.Beatmaps, beatmap)
	}
	err = uc.db.InsertBeatmapSet(&meta)
	return &meta, err
}

func (uc *usecase) DownloadMap(setId int) (*entity.BeatmapFile, error) {
	// Downloading map
	var bm *entity.BeatmapMeta
	bm, err := uc.beatmapsService.CheckBeatmapAvailability(setId)
	if err != nil {
		bm, err = uc.SaveMapMeta(setId)
		if err != nil {
			return nil, err
		}
		err = uc.beatmapsService.SaveBeatmapFile(setId)
		if err != nil {
			return nil, err
		}
	}
	//then
	return uc.ServeBeatmap(setId, bm)
}

func (uc *usecase) ServeBeatmap(setId int, beatmap *entity.BeatmapMeta) (*entity.BeatmapFile, error) {
	body, err := uc.beatmapsService.ServeBeatmap(setId)
	if err != nil {
		return nil, err
	}
	songTitle := fmt.Sprintf("%s - %s", beatmap.Artist, beatmap.Title)
	return &entity.BeatmapFile{Body: body, Title: songTitle}, nil
}
