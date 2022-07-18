package middleware

import (
	"strings"

	"github.com/0RAJA/Rutils/pkg/app"
	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/0RAJA/chat_app/src/dao"
	"github.com/0RAJA/chat_app/src/global"
	"github.com/0RAJA/chat_app/src/model"
	"github.com/0RAJA/chat_app/src/myerr"
	"github.com/gin-gonic/gin"
)

// Auth 鉴权中间件,用于解析并写入token
func Auth() func(c *gin.Context) {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader(global.PvSettings.Token.AuthorizationKey)
		if len(authorizationHeader) == 0 {
			c.Next()
			return
		}
		fields := strings.SplitN(authorizationHeader, " ", 2)
		if len(fields) != 2 || strings.ToLower(fields[0]) != global.PvSettings.Token.AuthorizationType {
			c.Next()
			return
		}
		accessToken := fields[1]
		payload, err := global.Maker.VerifyToken(accessToken)
		if err != nil {
			c.Next()
			return
		}
		var content model.Content
		if err := content.Unmarshal(payload.Content); err != nil {
			c.Next()
			return
		}
		c.Set(global.PvSettings.Token.AuthorizationKey, &content)
		c.Next()
	}
}

// MustUser 必须是用户
func MustUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		rly := app.NewResponse(c)
		val, ok := c.Get(global.PvSettings.Token.AuthorizationKey)
		if !ok {
			rly.Reply(myerr.AuthNotExist)
			c.Abort()
			return
		}
		data := val.(*model.Content)
		if data.Type != model.UserToken {
			rly.Reply(myerr.AuthenticationFailed)
			c.Abort()
			return
		}
		ok, err := dao.Group.DB.ExistsUserByID(c, data.ID)
		if err != nil {
			global.Logger.Error(err.Error(), ErrLogMsg(c)...)
			rly.Reply(errcode.ErrServer)
			c.Abort()
			return
		}
		if !ok {
			rly.Reply(myerr.UserNotFound)
			c.Abort()
			return
		}
		c.Next()
	}
}

// MustAccount 必须是账号
func MustAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		rly := app.NewResponse(c)
		val, ok := c.Get(global.PvSettings.Token.AuthorizationKey)
		if !ok {
			rly.Reply(myerr.AuthNotExist)
			c.Abort()
			return
		}
		data := val.(*model.Content)
		if data.Type != model.UserToken {
			rly.Reply(myerr.AuthenticationFailed)
			c.Abort()
			return
		}
		ok, err := dao.Group.DB.ExistsAccountByID(c, data.ID)
		if err != nil {
			global.Logger.Error(err.Error(), ErrLogMsg(c)...)
			rly.Reply(errcode.ErrServer)
			c.Abort()
			return
		}
		if !ok {
			rly.Reply(myerr.AccountNotFound)
			c.Abort()
			return
		}
		c.Next()
	}
}

// GetTokenContent 获取token内容
func GetTokenContent(c *gin.Context) (*model.Content, bool) {
	val, ok := c.Get(global.PvSettings.Token.AuthorizationKey)
	if !ok {
		return nil, false
	}
	return val.(*model.Content), true
}
