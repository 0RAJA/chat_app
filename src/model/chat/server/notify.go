package server

import "github.com/0RAJA/chat_app/src/model"

type CreateNotify struct {
	EnToken    string           `json:"en_token,omitempty"`
	RelationID int64            `json:"relation_id,omitempty"`
	AccountID  int64            `json:"account_id,omitempty"`
	MsgContent string           `json:"msg_content,omitempty"`
	MsgExtent  *model.MsgExtend `json:"msg_extent"`
}
