package server

import (
	"github.com/0RAJA/chat_app/src/model/reply"
)

// chat 中 server 端有关消息请求的结构

type SendMsg struct {
	EnToken string               `json:"en_token"` // 加密后的Token
	MsgInfo reply.MsgInfoWithRly `json:"msg_info"`
}

type ReadMsg struct {
	EnToken  string  `json:"en_token"`  // 加密后的Token
	MsgIDs   []int64 `json:"msg_ids"`   // 已读消息IDs
	ReaderID int64   `json:"reader_id"` // 读者账号ID
}

type MsgType string

const (
	MsgPin    MsgType = "pin"    // 置顶消息
	MsgTop    MsgType = "top"    // 置顶消息
	MsgRevoke MsgType = "revoke" // 撤回消息
)

type UpdateMsgState struct {
	EnToken string  `json:"en_token"` // 加密后的Token
	MsgType MsgType `json:"type"`     // 消息类型 [pin,top,revoke]
	MsgID   int64   `json:"msg_id"`   // 消息ID
	State   bool    `json:"state"`    // 状态设置
}
