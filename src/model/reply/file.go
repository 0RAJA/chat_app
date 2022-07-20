package reply

import "time"

type PublishFile struct {
	FileType string `json:"file_type"`
	FileSize int64 `json:"file_size"`
	Url      string `json:"url"`
	CreateAt time.Time `json:"create_at"`
}
