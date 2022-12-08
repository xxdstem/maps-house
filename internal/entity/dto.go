package entity

type BeatmapMetaDTO struct {
	ID          int     `json:"beatmapset_id" gorm:"column:beatmapset_id"`
	Artist      string  `json:"artist"`
	Title       string  `json:"title"`
	Creator     string  `json:"creator"`
	Tags        string  `json:"tags"`
	TotalLength int     `json:"total_length" gorm:"column:beatmapset_id"`
	BPM         float64 `json:"bpm"`
	Genre       int     `json:"genre_id"`
	Language    int     `json:"language_id"`
	//Ranked      int8    `json:"approved" gorm:"column:ranked"`
}

type BeatmapDTO struct {
	ID       int     `json:"beatmap_id,string" gorm:"column:beatmap_id"`
	SetID    int     `json:"beatmapset_id,string" gorm:"column:beatmapset_id"`
	MD5      string  `json:"file_md5" gorm:"column:file_md5"`
	Version  string  `json:"version"`
	AR       float64 `json:"diff_approach,string"`
	OD       float64 `json:"diff_overall,string"`
	CS       float64 `json:"diff_size,string"`
	HP       float64 `json:"diff_drain,string"`
	Length   int     `json:"hit_length,string" gorm:"column:hit_length"`
	Ranked   int8    `json:"approved,string" gorm:"column:ranked"`
	MaxCombo int     `json:"max_combo,string" gorm:"column:max_combo"`
	Mode     int8    `json:"mode,string"`
}
