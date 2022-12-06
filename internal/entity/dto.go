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
	ID       int     `json:"beatmap_id" gorm:"column:beatmap_id"`
	SetID    int     `json:"beatmapse_id" gorm:"column:beatmapse_id"`
	MD5      string  `json:"file_md5" gorm:"column:file_md5"`
	Version  string  `json:"version"`
	AR       float64 `json:"ar"`
	OD       float64 `json:"od"`
	CS       float64 `json:"cs"`
	HP       float64 `json:"hp"`
	Length   int     `json:"hit_length" gorm:"column:hit_length"`
	Ranked   int8    `json:"approved" gorm:"column:ranked"`
	MaxCombo int     `json:"max_combo" gorm:"column:max_combo"`
	Mode     int     `json:"mode"`
}
