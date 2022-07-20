package request

import (
	"github.com/0RAJA/chat_app/src/model/common"
)

type CreateAccount struct {
	Name   string `json:"name" binding:"required,gte=1,lte=20"`    // 名称
	Avatar string `json:"avatar" binding:"required,gte=1,lte=100"` // 头像
}

type GetAccountToken struct {
	AccountID int64 `form:"account_id" binding:"required,gte=1"` // 账号ID
}

type DeleteAccount struct {
	AccountID int64 `json:"account_id" binding:"required,gte=1"` // 账号ID
}

type UpdateAccount struct {
	Name      string `json:"name" binding:"required,gte=1,lte=20"`       // 名称
	Avatar    string `json:"avatar" binding:"required,gte=1,lte=100"`    // 头像
	Gender    string `json:"gender" binding:"required,oneof=男 女 未知"`     // 性别
	Signature string `json:"signature" binding:"required,gte=1,lte=200"` // 个性签名
}

type GetAccountByID struct {
	AccountID int64 `form:"account_id" binding:"required,gte=1"` // 账号ID
}

type GetAccountsByName struct {
	Name string `form:"name" binding:"required,gte=1,lte=20"` // 搜索名称
	common.Pager
}
