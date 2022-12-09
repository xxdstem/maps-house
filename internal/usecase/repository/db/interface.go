package db

import (
	"maps-house/internal/entity"
)

type Repository interface {
	GetBeatmapsBySetId(setId int) (*entity.BeatmapMeta, error)
	InsertBeatmapSet(meta *entity.BeatmapMeta) error
}
