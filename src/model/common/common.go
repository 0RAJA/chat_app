package common

import (
	"time"
)

// Pager 分页
type Pager struct {
	Page     int32 `json:"page" form:"page"`           // 第几页
	PageSize int32 `json:"page_size" form:"page_size"` // 每页大小
}

// State 状态码
type State struct {
	Code int         `json:"status_code"`    // 状态码，0-成功，其他值-失败
	Msg  string      `json:"status_msg"`     // 返回状态描述
	Data interface{} `json:"data,omitempty"` // 失败时返回空
}

// List 列表
type List struct {
	List  interface{} `json:"list"`  // 数据
	Total int         `json:"total"` // 总数
}

// Token token
type Token struct {
	Token     string    `json:"token"`      // token
	ExpiredAt time.Time `json:"expired_at"` // token过期时间
}
