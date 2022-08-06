package request

import (
	"mime/multipart"

	"github.com/0RAJA/chat_app/src/model/common"
)

type GetMsgsByRelationIDAndTime struct {
	RelationID int64 `form:"relation_id" binding:"required,gte=1"` // 关系ID
	LastTime   int32 `form:"last_time" binding:"required,gte=1"`   // 拉取消息的最晚时间(精确到秒)
	common.Pager
}

type GetPinMsgsByRelationID struct {
	RelationID int64 `form:"relation_id" binding:"required,gte=1"` // 关系ID
	common.Pager
}

type GetRlyMsgsInfoByMsgID struct {
	RelationID int64 `form:"relation_id" binding:"required,gte=1"` // 关系ID
	RlyMsgID   int64 `from:"rly_msg_id" binding:"required,gte=1"`  // 回复消息ID
	common.Pager
}

type GetTopMsgByRelationID struct {
	RelationID int64 `form:"relation_id" binding:"required,gte=1"` // 关系ID
}

type UpdateMsgPin struct {
	ID         int64 `json:"id" binding:"required,gte=1"`          // 消息ID
	RelationID int64 `json:"relation_id" binding:"required,gte=1"` // 关系ID
	IsPin      *bool `json:"is_pin" binding:"required"`            // 是否pin
}

type UpdateMsgRevoke struct {
	ID       int64 `json:"id" binding:"required,gte=1"`  // 消息ID
	IsRevoke *bool `json:"is_revoke" binding:"required"` // 是否撤回
}

type UpdateMsgTop struct {
	ID         int64 `json:"id" binding:"required,gte=1"`          // 消息ID
	RelationID int64 `json:"relation_id" binding:"required,gte=1"` // 关系ID
	IsTop      *bool `json:"is_top" binding:"required"`            // 是否置顶
}

type RevokeMsg struct {
	ID int64 `json:"id" binding:"required"` // 消息ID
}

type CreateFileMsg struct {
	RelationID int64          `form:"relation_id" binding:"required,gte=1"` // 关系ID
	File       multipart.File `form:"file" binding:"required"`              // 文件
	RlyMsgID   int64          `form:"rly_msg_id"`                           // 回复消息ID
}
