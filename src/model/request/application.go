package request

import (
	"github.com/0RAJA/chat_app/src/model/common"
)

type CreateApplication struct {
	AccountID int64  `json:"account_id" binding:"required,gte=1"`  // 目标账号ID
	ApplyMsg  string `json:"apply_msg" binding:"required,lte=200"` // 申请信息
}

type DeleteApplication struct {
	AccountID int64 `json:"account_id" binding:"required,gte=1"` // 目标账号ID
}

type AcceptApplication struct {
	AccountID int64 `json:"account_id" binding:"required,gte=1"` // 目标账号ID
}

type RefuseApplication struct {
	AccountID int64  `json:"account_id" binding:"required,gte=1"`   // 目标账号ID
	RefuseMsg string `json:"refuse_msg" binding:"required,lte=200"` // 拒绝信息
}

type ListApplications struct {
	common.Pager
}
