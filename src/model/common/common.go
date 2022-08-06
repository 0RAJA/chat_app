package common

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/go-playground/validator/v10"
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

// NewState 创建一个标准回复格式
func NewState(merr errcode.Err, datas ...interface{}) *State {
	var data interface{}
	if len(datas) > 0 {
		data = datas[0]
	}
	if merr == nil {
		merr = errcode.StatusOK
	} else {
		data = nil
	}
	return &State{
		Code: merr.ECode(),
		Msg:  merr.Error(),
		Data: data,
	}
}

var validate *validator.Validate
var validateOnce sync.Once

// Decode 将json格式的数据解析到结构体,并进行校验
func Decode(data string, v interface{}) error {
	if err := json.Unmarshal([]byte(data), v); err != nil {
		return err
	}
	validateOnce.Do(func() { validate = validator.New() })
	return validate.Struct(v)
}
