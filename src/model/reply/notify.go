package reply

import (
	"time"

	"github.com/0RAJA/chat_app/src/model"
)

type GroupNotify struct {
	ID         int64            `json:"id"`
	RelationID int64            `json:"relation_id"`
	MsgContent string           `json:"msg_content"`
	MsgExpand  *model.MsgExtend `json:"msg_expand"`
	AccountID  int64            `json:"account_id"`
	CreateAt   time.Time        `json:"create_at"`
	ReadIds    []int64          `json:"read_ids"`
}

type UpdateNotify struct {
}

type GetNotify struct {
	ID         int64            `json:"id"`
	RelationID int64            `json:"relation_id"`
	MsgContent string           `json:"msg_content"`
	MsgExpand  *model.MsgExtend `json:"msg_expand"`
	AccountID  int64            `json:"account_id"`
	CreateAt   time.Time        `json:"create_at"`
	ReadIds    []int64          `json:"read_ids"`
}
