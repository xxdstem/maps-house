package usecase

import (
	"maps-house/internal/entity"
	"maps-house/pkg/logger"
	"time"
)

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

func (uc *usecase) CacheMapFromBancho(setId int) error {
	maps, err := uc.osuApiService.GetBeatmapData(setId)
	if err != nil {
		return err
	}
	var firstBeatmap entity.BeatmapDTO = maps[0]
	var curTime time.Time = time.Now()
	meta := entity.BeatmapMeta{
		ID:         firstBeatmap.SetID,
		Artist:     firstBeatmap.Artist,
		Title:      firstBeatmap.Title,
		Creator:    firstBeatmap.Creator,
		Tags:       firstBeatmap.Tags,
		Length:     firstBeatmap.Length,
		BPM:        firstBeatmap.BPM,
		LanguageID: firstBeatmap.LanguageID,
		GenreID:    firstBeatmap.GenreID,
		Downloaded: false,
		Beatmaps:   []*entity.Beatmap{},
	}
	for i := range maps {
		var beatmapDto entity.BeatmapDTO = maps[i]
		t, err := time.Parse("2006-01-02 15:04:05", beatmapDto.LastUpdate)
		if err != nil {
			return err
		}
		beatmap := &entity.Beatmap{
			ID:         beatmapDto.ID,
			SetID:      beatmapDto.SetID,
			MD5:        beatmapDto.MD5,
			Version:    beatmapDto.Version,
			HitLength:  beatmapDto.HitLength,
			Ranked:     beatmapDto.Ranked,
			LastUpdate: t.Unix(),
			ApiUpdate:  curTime.Unix(),
			AR:         beatmapDto.AR,
			OD:         beatmapDto.OD,
			HP:         beatmapDto.HP,
			CS:         beatmapDto.CS,
		}
		meta.Beatmaps = append(meta.Beatmaps, beatmap)
	}
	err = uc.db.InsertBeatmapSet(&meta)

	return err
}

func (uc *usecase) DownloadMap(setId int) (*entity.BeatmapFile, error) {
	// Downloading map

	// Пишу пока всё здесь, потом разбросаю код
	if err := uc.beatmapsService.CheckBeatmapAvailability(setId); err != nil {
		err = uc.CacheMapFromBancho(setId)
		return nil, err
	} else {
		beatmap, err := uc.DownloadMap(setId)
		if err != nil {
			return nil, err
		}
		log.Info(beatmap)
	}
	//then
	return uc.ServeBeatmap(setId)
}

func (uc *usecase) ServeBeatmap(setId int) (*entity.BeatmapFile, error) {
	log.Info(setId)
	return nil, nil
}
