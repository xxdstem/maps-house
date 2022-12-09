package entity

type BeatmapMeta struct {
	SetID      int        `json:"id"`
	Artist     string     `json:"artist"`
	Title      string     `json:"title"`
	Downloaded bool       `json:"is_downloaded"`
	Beatmaps   []*Beatmap `json:"beatmaps" gorm:"foreignKey:beatmapset_id;references:beatmapset_id"`
}

type Beatmap struct {
	ID      int     `json:"id" gorm:"type:int;primaryKey;"`
	SetID   int     `json:"-" gorm:"column:beatmapset_id;type:int"`
	Version string  `json:"version" gorm:"type:varchar(96);collate:latin1_swedish_ci"`
	AR      float64 `json:"ar"`
	OD      float64 `json:"od"`
	HP      float64 `json:"hp"`
	CS      float64 `json:"cs"`
}

type BeatmapFile struct {
	Body []byte
}
