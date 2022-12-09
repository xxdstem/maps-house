package entity

type BeatmapMeta struct {
	ID         int        `json:"id" gorm:"column:beatmapset_id"`
	Artist     string     `json:"artist"`
	Title      string     `json:"title"`
	Creator    string     `json:"creator"`
	Tags       string     `json:"tags"`
	Length     int        `json:"total_length" gorm:"column:total_length"`
	BPM        float64    `json:"bpm"`
	GenreID    int        `json:"genre_id"`
	LanguageID int        `json:"language_id"`
	Downloaded bool       `json:"is_downloaded"`
	Beatmaps   []*Beatmap `json:"beatmaps" gorm:"foreignKey:beatmapset_id;references:beatmapset_id"`
}

type Beatmap struct {
	ID         int     `json:"id" gorm:"column:beatmap_id;type:int;primaryKey;"`
	SetID      int     `json:"-" gorm:"column:beatmapset_id;type:int"`
	MD5        string  `json:"md5" gorm:"column:file_md5"`
	Version    string  `json:"version" gorm:"type:varchar(96);collate:latin1_swedish_ci"`
	HitLength  int     `json:"hit_length" gorm:"column:hit_length"`
	Ranked     int8    `json:"ranked"`
	LastUpdate int64   `json:"last_update" gorm:"column:latest_update;type:int"`
	ApiUpdate  int64   `json:"api_update" gorm:"column:api_update"`
	AR         float64 `json:"ar"`
	OD         float64 `json:"od"`
	HP         float64 `json:"hp"`
	CS         float64 `json:"cs"`
}

type BeatmapFile struct {
	Body []byte
}
