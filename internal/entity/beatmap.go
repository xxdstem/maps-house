package entity

type Beatmap struct {
	BeatmapId    int
	BeatmapsetId int
	Version      string `gorm:"type:varchar(96);collate:latin1_swedish_ci"`
	AR           float64
	OD           float64
	HP           float64
	CS           float64
}
