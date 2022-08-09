package server

import (
	"github.com/0RAJA/chat_app/src/model/reply"
)

type SendMsg struct {
	EnToken string `json:"en_token"` // 加密后的Token
	reply.MsgInfoWithRly
}

type ReadMsg struct {
	EnToken   string `json:"en_token"`   // 加密后的Token
	MsgID     int64  `json:"msg_id"`     // 消息ID
	AccountID int64  `json:"account_id"` // 读者账号ID
}
