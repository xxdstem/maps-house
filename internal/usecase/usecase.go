package usecase

import (
	"errors"
	"fmt"
	"maps-house/internal/entity"
	"maps-house/pkg/logger"
	"time"
)

type BeatmapsService interface {
	CheckBeatmapAvailability(bm *entity.BeatmapMeta) error
	SaveBeatmapFile(setId int, chimu bool) error
	ServeBeatmap(setId int) ([]byte, error)
	RemoveBeatmapFile(setId int) error
	CheckUpdateConditions(bm *entity.BeatmapMeta) bool
}

type OsuApiService interface {
	GetBeatmapData(setId int) ([]entity.BeatmapDTO, error)
}

type DbRepository interface {
	InsertBeatmapSet(meta *entity.BeatmapMeta) error
	UpdateBeatmapSet(meta *entity.BeatmapMeta) error
	GetBeatmapsBySetId(setId int) (*entity.BeatmapMeta, error)
	DeleteBeatmapSet(setId int) error
	SetDownloadedStatus(setId int, state bool) error
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

func (uc *usecase) FetchBeatmapMeta(setId int) (*entity.BeatmapMeta, error) {
	maps, err := uc.osuApiService.GetBeatmapData(setId)
	if err != nil {
		return nil, err
	}
	if len(maps) == 0 {
		return nil, errors.New("no beatmaps?")
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
		ApiUpdate:    curTime.Unix(),
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
			AR:           beatmapDto.AR,
			OD:           beatmapDto.OD,
			HP:           beatmapDto.HP,
			CS:           beatmapDto.CS,
		}
		if meta.LastUpdate < t.Unix() {
			meta.LastUpdate = t.Unix()
		}
		meta.Beatmaps = append(meta.Beatmaps, beatmap)
	}
	return &meta, err
}

func (uc *usecase) DownloadMap(setId int) (*entity.BeatmapFile, error) {
	// Downloading map
	bm, err := uc.db.GetBeatmapsBySetId(setId)
	if err != nil {
		return nil, err
	}
	err = uc.beatmapsService.CheckBeatmapAvailability(bm)
	bmInDB := bm != nil
	if err != nil {
		log.Error("[", setId, "] Beatmap availability check not passed")
		bm, err := uc.FetchBeatmapMeta(setId)
		if err != nil {
			log.Error("[", setId, "] Beatmap available, cannot fetch beatmap meta")
			return nil, err
		}
		if !bmInDB {
			err = uc.db.InsertBeatmapSet(bm)
			if err != nil {
				return nil, err
			}
		}

		err = uc.beatmapsService.SaveBeatmapFile(setId, false)
		if err != nil {
			log.Error("[", setId, "] Cannot save beatmap file")
			return nil, err
		}
		log.Info("[", setId, "] Beatmap successfully downloaded")
	} else if uc.beatmapsService.CheckUpdateConditions(bm) {
		apiData, err := uc.FetchBeatmapMeta(setId)
		if err != nil {
			//remove map?
			log.Info("[", setId, "] Cannot fetch beatmap meta, gonna remove map")
			err = uc.beatmapsService.RemoveBeatmapFile(setId)
			uc.db.DeleteBeatmapSet(setId)
			return nil, err
		}
		if bm.LastUpdate < apiData.LastUpdate {
			// re-download
			log.Info("[", setId, "] Beatmap LastUpdate changed, re-downloading")

			uc.db.UpdateBeatmapSet(apiData)

			err = uc.beatmapsService.SaveBeatmapFile(setId, false)
			if err != nil {
				return nil, err
			}
		}
		// re-check if beatmap not updated
	}
	uc.db.SetDownloadedStatus(setId, true)
	//then
	return uc.ServeBeatmap(setId, bm)
}

func (uc *usecase) ServeBeatmap(setId int, beatmap *entity.BeatmapMeta) (*entity.BeatmapFile, error) {
	if beatmap == nil {
		return nil, errors.New("haven't got beatmap in serve?")
	}
	body, err := uc.beatmapsService.ServeBeatmap(setId)
	if err != nil {
		return nil, err
	}
	songTitle := fmt.Sprintf("%s - %s", beatmap.Artist, beatmap.Title)
	return &entity.BeatmapFile{Body: body, Title: songTitle}, nil
}
