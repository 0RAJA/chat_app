package logic

import (
	"errors"

	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/0RAJA/Rutils/pkg/utils"
	"github.com/0RAJA/chat_app/src/dao"
	"github.com/0RAJA/chat_app/src/global"
	mid "github.com/0RAJA/chat_app/src/middleware"
	"github.com/0RAJA/chat_app/src/myerr"
	"github.com/0RAJA/chat_app/src/pkg/mark"
	"github.com/gin-gonic/gin"
)

type email struct {
}

// ExistEmail 是否存在email
func (email) ExistEmail(c *gin.Context, emailStr string) (bool, errcode.Err) {
	ok, err := dao.Group.Redis.ExistEmail(c, emailStr)
	if err == nil {
		return ok, nil
	}
	global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
	ok, err = dao.Group.DB.ExistEmail(c, emailStr)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return false, errcode.ErrServer
	}
	return ok, nil
}

// CheckEmailNotExists 判断邮箱是否已经注册
func CheckEmailNotExists(c *gin.Context, emailStr string) errcode.Err {
	ok, err := email{}.ExistEmail(c, emailStr)
	if err != nil {
		return err
	}
	if ok {
		return myerr.EmailExists
	}
	return nil
}

// SendEmail 发送邮件
func (email) SendEmail(c *gin.Context, emailStr string) errcode.Err {
	// 判断是否已经注册邮箱
	if err := CheckEmailNotExists(c, emailStr); err != nil {
		return err
	}
	// 判断发送频率
	if global.EmailMark.CheckUserExist(emailStr) {
		return myerr.EmailSendMany
	}
	// 异步发送邮件
	global.Worker.SendTask(func() {
		code := utils.RandomString(global.PbSettings.Rule.CodeLength)
		if err := global.EmailMark.SendMail(emailStr, code); err != nil && !errors.Is(err, mark.ErrSendTooMany) {
			global.Logger.Error(err.Error())
		}
	})
	return nil
}
