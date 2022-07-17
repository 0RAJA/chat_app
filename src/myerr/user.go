package myerr

import (
	"github.com/0RAJA/Rutils/pkg/app/errcode"
)

var (
	UserNotFound     = errcode.NewErr(1001, "用户不存在")
	PasswordNotValid = errcode.NewErr(1002, "密码错误")
)
