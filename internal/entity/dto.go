package entity

type BeatmapDTO struct {
	ID         int     `json:"beatmap_id,string" gorm:"column:beatmap_id"`
	SetID      int     `json:"beatmapset_id,string" gorm:"column:beatmapset_id"`
	Artist     string  `json:"artist"`
	Title      string  `json:"title"`
	Creator    string  `json:"creator"`
	Tags       string  `json:"tags"`
	MD5        string  `json:"file_md5" gorm:"column:file_md5"`
	Version    string  `json:"version"`
	BPM        float64 `json:"bpm,string"`
	AR         float64 `json:"diff_approach,string"`
	OD         float64 `json:"diff_overall,string"`
	CS         float64 `json:"diff_size,string"`
	HP         float64 `json:"diff_drain,string"`
	LastUpdate string  `json:"last_update"`
	Length     int     `json:"total_length,string"`

	HitLength  int  `json:"hit_length,string" gorm:"column:hit_length"`
	Ranked     int8 `json:"approved,string" gorm:"column:ranked"`
	MaxCombo   int  `json:"max_combo,string" gorm:"column:max_combo"`
	Mode       int8 `json:"mode,string"`
	LanguageID int8 `json:"language_id,string"`
	GenreID    int8 `json:"genre_id,string"`
}

type BeatmapsResultDTO struct {
	Result *BeatmapMeta `json:"result"`
}
