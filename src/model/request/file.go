package request

import "mime/multipart"

type PublishFile struct {
	File       *multipart.FileHeader `form:"file"  binding:"required"`
	RelationID string                `form:"relation_id"  binding:"required"`
	AccountID  string                `form:"account_id"  binding:"required"`
}
