package request

import (
	"github.com/0RAJA/chat_app/src/model/common"
)

type GetMsgsByRelationIDAndTime struct {
	RelationID int64 `form:"relation_id"` // 关系ID
	LastTime   int32 `form:"last_time"`   // 拉取消息的最晚时间(精确到秒)
	common.Pager
}

type GetPinMsgsByRelationID struct {
	RelationID int64 `form:"relation_id"` // 关系ID
	common.Pager
}

type GetRlyMsgsInfoByMsgID struct {
	RelationID int64 `form:"relation_id"` // 关系ID
	RlyMsgID   int64 `from:"rly_msg_id"`  // 回复消息ID
	common.Pager
}

type GetTopMsgByRelationID struct {
	RelationID int64 `form:"relation_id"` // 关系ID
}

type UpdateMsgPin struct {
	ID         int64 `json:"id"`          // 消息ID
	RelationID int64 `json:"relation_id"` // 关系ID
	IsPin      *bool `json:"is_pin"`      // 是否pin
}

type UpdateMsgRevoke struct {
	ID       int64 `json:"id"`        // 消息ID
	IsRevoke *bool `json:"is_revoke"` // 是否撤回
}

type UpdateMsgTop struct {
	ID         int64 `json:"id"`          // 消息ID
	RelationID int64 `json:"relation_id"` // 关系ID
	IsTop      *bool `json:"is_top"`      // 是否置顶
}

type RevokeMsg struct {
	ID int64 `json:"id"` // 消息ID
}
