package reply

import (
	"time"
)

type PublishFile struct {
	ID       int64     `json:"id"`
	FileType string    `json:"file_type"`
	FileSize int64     `json:"file_size"`
	Url      string    `json:"url"`
	CreateAt time.Time `json:"create_at"`
}

type DeleteFile struct {
}
type File struct {
	FileID    int64    `json:"file_id"`
	FileName  string    `json:"file_name"`
	FileType  string    `json:"file_type"`
	FileSize  int64     `json:"file_size"`
	Url       string    `json:"url"`
	AccountID int64     `json:"account_id"`
	CreateAt  time.Time `json:"create_at"`
}
type GetRelationFile struct {
	FileList []*File `json:"file_list"`
}
