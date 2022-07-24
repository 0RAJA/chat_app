package request

import "time"

type CreateNotify struct {
	RelationID int64   `json:"relation_id" form:"relation_id" binding:"required"`
	MsgContent string  `json:"msg_content" form:"msg_content" binding:"required"`
	MsgExpand  string  `json:"msg_expand"  form:"mag_expand" binding:"required"`
	AccountID  int64   `json:"account_id" form:"account_id" binding:"required"`
}
type UpdateNotify struct {
	RelationID int64   `json:"relation_id" form:"relation_id" binding:"required"`
	MsgContent string  `json:"msg_content" form:"msg_content" binding:"required"`
	MsgExpand  string  `json:"msg_expand"  form:"mag_expand" binding:"required"`
	AccountID  int64   `json:"account_id" form:"account_id" binding:"required"`
	UpdateAt  time.Time `json:"update_at" form:"update_at" binding:"required"`
}
type GetNotifyByID struct {
	RelationID int64 `json:"relation_id" form:"relation_id"`
}