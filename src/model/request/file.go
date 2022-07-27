package request

import (
	"mime/multipart"
)

type PublishFile struct {
	File       *multipart.FileHeader `form:"file"  binding:"required"`
	RelationID int64                 `form:"relation_id"  binding:"required"`
	AccountID  int64                 `form:"account_id"  binding:"required"`
}

type DeleteFile struct {
	FileID int64 `json:"file_id" form:"file_id" binding:"required"`
}

type GetRelationFile struct {
	RelationID int64 `json:"relation_id"  form:"relation_id" binding:"required"`
}

type UploadAvatar struct {
	File       *multipart.FileHeader `form:"file"  binding:"required"`
	AccountID  int64                 `form:"account_id"  binding:"required"`
}