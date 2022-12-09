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
	ID       string `json:"beatmap_id" gorm:"column:beatmap_id"`
	SetID    string `json:"beatmapset_id" gorm:"column:beatmapset_id"`
	MD5      string `json:"file_md5" gorm:"column:file_md5"`
	Version  string `json:"version"`
	AR       string `json:"diff_approach"`
	OD       string `json:"diff_overall"`
	CS       string `json:"diff_size"`
	HP       string `json:"diff_drain"`
	Length   string `json:"hit_length" gorm:"column:hit_length"`
	Ranked   string `json:"approved" gorm:"column:ranked"`
	MaxCombo string `json:"max_combo" gorm:"column:max_combo"`
	Mode     string `json:"mode"`
}

type BeatmapsResultDTO struct {
	Result *BeatmapMeta `json:"result"`
}
