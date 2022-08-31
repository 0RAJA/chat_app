package model

import (
	"mime/multipart"
	"time"

	"github.com/jackc/pgtype"
)

type MsgType string

const (
	MsgTypeText MsgType = "text"
	MsgTypeFile MsgType = "file"
)

// ExpandToJson MsgExtend 转 pgtype.Json 可以存nil
// 参数: 消息扩展信息
// 返回: pgtype.Json 对象
func ExpandToJson(extend *MsgExtend) (pgtype.JSON, error) {
	data := pgtype.JSON{}
	err := data.Set(extend)
	return data, err
}

// JsonToExpand pgtype.Json 转 MsgExtend,
// 参数: pgtype.Json 对象(如果存的json为nil或未定义则返回nil)
// 返回: 解析后的消息扩展信息(可能为nil)
func JsonToExpand(data pgtype.JSON) (*MsgExtend, error) {
	if data.Status != pgtype.Present {
		return nil, nil
	}
	var extend = &MsgExtend{}
	err := data.AssignTo(&extend)
	return extend, err
}

type Remind struct {
	Idx       int64 `json:"idx" binding:"required,gte=1" validate:"required,gte=1"`        // 第几个@
	AccountID int64 `json:"account_id" binding:"required,gte=1" validate:"required,gte=1"` // 被@的账号ID
}

// MsgExtend 消息扩展信息 可能为null
type MsgExtend struct {
	Remind []Remind `json:"remind"` // @的描述信息
}

type GetMsgsByRelationIDAndTime struct {
	AccountID, RelationID int64
	LastTime              time.Time
	Limit, Offset         int32
}

type GetPinMsgsByRelationID struct {
	AccountID, RelationID int64
	Limit, Offset         int32
}

type GetRlyMsgsInfoByMsgID struct {
	AccountID, RelationID, RlyMsgID int64
	Limit, Offset                   int32
}

type GetTopMsgByRelationID struct {
	AccountID, RelationID int64
}

type UpdateMsgPin struct {
	AccountID, RelationID, MsgID int64
	IsPin                        bool
}

type UpdateMsgTop struct {
	AccountID, RelationID, MsgID int64
	IsTop                        bool
}

type RevokeMsg struct {
	AccountID, MsgID int64
}

type CreateFileMsg struct {
	AccountID, RelationID, RlyMsgID int64
	File                            *multipart.FileHeader
}

type CreateMsg struct {
	AccountID, RelationID, FileID, RlyMsgID int64
	NotifyType, MsgType, MsgContent         string
	MsgExtend                               *MsgExtend
}

type GetMsgsByContent struct {
	RelationID, AccountID int64
	Limit, Offset         int32
	Content               string
}

type FeedMsgsByAccountIDAndTime struct {
	AccountID     int64
	Limit, Offset int32
	LastTime      time.Time
}
