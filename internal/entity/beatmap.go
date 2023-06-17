package entity

type BeatmapMeta struct {
	BeatmapsetID int        `json:"id" gorm:"column:beatmapset_id;primaryKey"`
	Artist       string     `json:"artist" gorm:"column:artist"`
	Title        string     `json:"title" gorm:"column:title"`
	Creator      string     `json:"creator" gorm:"column:creator"`
	Tags         string     `json:"tags" gorm:"column:tags"`
	Length       int        `json:"total_length" gorm:"column:total_length"`
	BPM          float64    `json:"bpm" gorm:"column:bpm"`
	GenreID      int8       `json:"genre_id" gorm:"column:genre_id"`
	LanguageID   int8       `json:"language_id" gorm:"column:language_id"`
	Downloaded   bool       `json:"is_downloaded" gorm:"column:is_downloaded"`
	Beatmaps     []*Beatmap `json:"beatmaps" gorm:"foreignKey:BeatmapsetID;references:BeatmapsetID"`
}

type Beatmap struct {
	ID           int     `json:"id" gorm:"column:beatmap_id;primaryKey;"`
	BeatmapsetID int     `json:"-" gorm:"column:beatmapset_id"`
	MD5          string  `json:"md5" gorm:"column:file_md5"`
	Version      string  `json:"version" gorm:"type:varchar(96);collate:latin1_swedish_ci"`
	HitLength    int     `json:"hit_length" gorm:"column:hit_length"`
	Ranked       int8    `json:"ranked" gorm:"column:ranked"`
	LastUpdate   int64   `json:"last_update" gorm:"column:latest_update;type:int"`
	ApiUpdate    int64   `json:"api_update" gorm:"column:api_update"`
	AR           float64 `json:"ar"  gorm:"column:ar"`
	OD           float64 `json:"od"  gorm:"column:od"`
	HP           float64 `json:"hp"  gorm:"column:hp"`
	CS           float64 `json:"cs"  gorm:"column:cs"`
}

type BeatmapFile struct {
	Body  []byte
	Title string
}
