package db

import (
	"maps-house/internal/entity"
)

type Repository interface {
	GetBeatmapBySetId(setId int) (*entity.Beatmap, error)
}
