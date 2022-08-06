package model

import (
	"mime/multipart"
	"time"
)

type MsgType string

const (
	MsgTypeText MsgType = "text"
	MsgTypeFile MsgType = "file"
)

type Remind struct {
	Idx       int64 `json:"idx" binding:"required,gte=1"`        // 第几个@
	AccountID int64 `json:"account_id" binding:"required,gte=1"` // 被@的账号ID
}

// MsgExtend 消息扩展信息 可能为null
type MsgExtend struct {
	Remind []Remind `json:"remind"` // @的描述信息
}

type GetMsgsByRelationIDAndTimeParams struct {
	AccountID, RelationID int64
	LastTime              time.Time
	Limit, Offset         int32
}

type GetPinMsgsByRelationIDParams struct {
	AccountID, RelationID int64
	Limit, Offset         int32
}

type GetRlyMsgsInfoByMsgIDParams struct {
	AccountID, RelationID, RlyMsgID int64
	Limit, Offset                   int32
}

type GetTopMsgByRelationIDParams struct {
	AccountID, RelationID int64
}

type UpdateMsgPinParams struct {
	AccountID, RelationID, MsgID int64
	IsPin                        bool
}

type UpdateMsgTopParams struct {
	AccountID, RelationID, MsgID int64
	IsTop                        bool
}

type RevokeMsgParams struct {
	AccountID, MsgID int64
}

type CreateFileMsgParams struct {
	AccountID, RelationID, RlyMsgID int64
	File                            multipart.File
}

type CreateMsgParams struct {
	AccountID, RelationID, FileID, RlyMsgID int64
	NotifyType, MsgType, MsgContent         string
	MsgExtend                               *MsgExtend
}
