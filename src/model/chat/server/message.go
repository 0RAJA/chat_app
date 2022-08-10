package server

import (
	"github.com/0RAJA/chat_app/src/model/reply"
)

// chat 中 server 端有关消息请求的结构

type SendMsg struct {
	EnToken string `json:"en_token"` // 加密后的Token
	reply.MsgInfoWithRly
}

type ReadMsg struct {
	EnToken  string  `json:"en_token"`  // 加密后的Token
	MsgIDs   []int64 `json:"msg_ids"`   // 已读消息IDs
	ReaderID int64   `json:"reader_id"` // 读者账号ID
}
