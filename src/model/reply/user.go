package reply

import (
	"time"

	"github.com/0RAJA/chat_app/src/model/common"
)

// UserInfo 用户信息
type UserInfo struct {
	ID       int64     `json:"id"`        // user id
	Email    string    `json:"email"`     // 邮箱
	CreateAt time.Time `json:"create_at"` // 创建时间
}

type Register struct {
	UserInfo  UserInfo     `json:"user_info"`  // 用户信息
	UserToken common.Token `json:"user_token"` // 用户令牌
}

type Login struct {
	UserInfo  UserInfo     `json:"user_info"`  // 用户信息
	UserToken common.Token `json:"user_token"` // 用户令牌
}
