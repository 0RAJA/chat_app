package reply

import (
	"time"

	"github.com/0RAJA/chat_app/src/model/common"
)

type CreateAccount struct {
	ID           int64        `json:"id"`            // 账号ID
	Name         string       `json:"name"`          // 名称
	Avatar       string       `json:"avatar"`        // 头像
	AccountToken common.Token `json:"account_Token"` // 账号Token
}

type GetAccountToken struct {
	AccountToken common.Token `json:"account_token"` // 账号Token
}

type GetAccountByID struct {
	ID        int64     `json:"id"`        // 账号ID
	Name      string    `json:"name"`      // 名称
	Avatar    string    `json:"avatar"`    // 头像
	Gender    string    `json:"gender"`    // 性别
	Signature string    `json:"signature"` // 个性签名
	CreateAt  time.Time `json:"create_at"` // 创建时间
}

type AccountInfo struct {
	ID     int64  `json:"account_id"` // 账号ID
	Name   string `json:"name"`       // 名称
	Avatar string `json:"avatar"`     // 头像
}

type GetAccountsByUserID struct {
	List  []*AccountInfo `json:"list"`  // 账号列表
	Total int64          `json:"total"` // 总数
}

type AccountFriendInfo struct {
	AccountInfo
	RelationID int64 // 好友关系ID，0表示没有好友关系
}

type GetAccountsByName struct {
	List  []*AccountFriendInfo `json:"list"`  // 账号列表
	Total int64                `json:"total"` // 总数
}
