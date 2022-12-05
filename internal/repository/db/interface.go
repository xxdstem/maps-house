package db

import (
	"mapsHouse/internal/entity"
)

type Repository interface {
	GetBeatmapBySetId(setId int) (*entity.Beatmap, error)
}
