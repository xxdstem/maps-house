package entity

type BeatmapMeta struct {
	BeatmapsetId int
	Artist       string
	Title        string

	Beatmaps []*Beatmap `gorm:"foreignKey:beatmapset_id;references:beatmapset_id"`
}

type Beatmap struct {
	BeatmapId    int    `gorm:"type:int;primaryKey;"`
	BeatmapsetId int    `json:"-" gorm:"colulmn:beatmapset_id;type:int"`
	Version      string `gorm:"type:varchar(96);collate:latin1_swedish_ci"`
	AR           float64
	OD           float64
	HP           float64
	CS           float64
}

