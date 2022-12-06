package entity

type BeatmapMeta struct {
	BeatmapsetId int        `json:"id"`
	Artist       string     `json:"artist"`
	Title        string     `json:"title"`
	Downloaded   bool       `json:"is_downloaded"`
	Beatmaps     []*Beatmap `json:"beatmaps" gorm:"foreignKey:beatmapset_id;references:beatmapset_id"`
}

type Beatmap struct {
	BeatmapId    int     `json:"id" gorm:"type:int;primaryKey;"`
	BeatmapsetId int     `json:"-" gorm:"colulmn:beatmapset_id;type:int"`
	Version      string  `json:"version" gorm:"type:varchar(96);collate:latin1_swedish_ci"`
	AR           float64 `json:"ar"`
	OD           float64 `json:"od"`
	HP           float64 `json:"hp"`
	CS           float64 `json:"cs"`
}

type BeatmapFile struct {
	Body []byte
}

type BeatmapsDto struct {
	Result *BeatmapMeta
}
