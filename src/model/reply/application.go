package reply

import (
	"time"
)

type ApplicationInfo struct {
	Account1ID int64     `json:"account1_id"` // 申请者账号ID
	Account2ID int64     `json:"account2_id"` // 目标账号ID
	ApplyMsg   string    `json:"apply_msg"`   // 申请信息
	RefuseMsg  string    `json:"refuse_msg"`  // 拒绝信息
	Status     string    `json:"status"`      // 状态 [已申请,已拒绝,已同意]
	CreateAt   time.Time `json:"create_at"`   // 创建时间
	UpdateAt   time.Time `json:"update_at"`   // 更新时间
	Name       string    `json:"name"`        // 对方账号名称
	Avatar     string    `json:"avatar"`      // 对方头像
}

type ListApplications struct {
	List  []*ApplicationInfo `json:"list"`
	Total int64              `json:"total"`
}
