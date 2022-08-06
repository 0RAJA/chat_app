package reply

import (
	"github.com/jackc/pgtype"
	"time"
)

type GroupNotify struct {
	ID         int64       `json:"id"`
	RelationID int64       `json:"relation_id"`
	MsgContent string      `json:"msg_content"`
	MsgExpand  pgtype.JSON `json:"msg_expand" swaggerignore:"true"`
	AccountID  int64       `json:"account_id"`
	CreateAt   time.Time   `json:"create_at"`
	ReadIds    []int64     `json:"read_ids"`
}

type UpdateNotify struct {
}

type GetNotify struct {
	ID         int64       `json:"id"`
	RelationID int64       `json:"relation_id"`
	MsgContent string      `json:"msg_content"`
	MsgExpand  pgtype.JSON `json:"msg_expand" swaggerignore:"true"`
	AccountID  int64       `json:"account_id"`
	CreateAt   time.Time   `json:"create_at"`
	ReadIds    []int64     `json:"read_ids"`
}
