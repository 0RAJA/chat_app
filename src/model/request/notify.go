package request

import "github.com/0RAJA/chat_app/src/model"

type CreateNotify struct {
	RelationID int64            `json:"relation_id" form:"relation_id" binding:"required"`
	MsgContent string           `json:"msg_content" form:"msg_content" binding:"required"`
	MsgExtend  *model.MsgExtend `json:"msg_expand"  form:"msg_expand"`
}
type UpdateNotify struct {
	ID         int64            `json:"id" form:"id" binding:"required"`
	RelationID int64            `json:"relation_id" form:"relation_id" binding:"required"`
	MsgContent string           `json:"msg_content" form:"msg_content" binding:"required"`
	MsgExtend  *model.MsgExtend ` json:"msg_expand"  form:"msg_expand"`
}
type GetNotifyByID struct {
	RelationID int64 `json:"relation_id" form:"relation_id" binding:"required"`
}
type DeleteNotify struct {
	ID         int64 `json:"id" form:"id" binding:"required"`
	RelationID int64 `json:"relation_id" form:"relation_id" binding:"required"`
}
