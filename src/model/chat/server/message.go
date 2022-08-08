package server

import (
	"github.com/0RAJA/chat_app/src/model/reply"
)

type SendMsg reply.MsgInfoWithRly

type ReadMsg struct {
	MsgID     int64 `json:"msg_id"`     // 消息ID
	AccountID int64 `json:"account_id"` // 读者账号ID
}
