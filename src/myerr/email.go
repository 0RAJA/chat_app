package myerr

import (
	"github.com/0RAJA/Rutils/pkg/app/errcode"
)

var (
	EmailSendMany     = errcode.NewErr(3001, "邮件发送频繁，请稍后再试")
	EmailCodeNotValid = errcode.NewErr(3002, "邮箱验证码校验失败")
	EmailSame         = errcode.NewErr(3003, "邮箱重复")
	EmailExists       = errcode.NewErr(3004, "邮箱已经注册")
)
