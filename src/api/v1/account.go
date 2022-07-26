package v1

import (
	"github.com/0RAJA/Rutils/pkg/app"
	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/0RAJA/chat_app/src/global"
	"github.com/0RAJA/chat_app/src/logic"
	mid "github.com/0RAJA/chat_app/src/middleware"
	"github.com/0RAJA/chat_app/src/model"
	"github.com/0RAJA/chat_app/src/model/request"
	"github.com/0RAJA/chat_app/src/myerr"
	"github.com/gin-gonic/gin"
)

type account struct {
}

// CreateAccount
// @Tags     account
// @Summary  创建账号
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                                  true  "Bearer 用户令牌"
// @Param    data           body      request.CreateAccount                   true  "创建账号信息"
// @Success  200            {object}  common.State{data=reply.CreateAccount}  "1001:参数有误 1003:系统错误 2008:身份验证失败 2012:账号数量超过限制 2011:账号名已经存在 2007:身份不存在"
// @Router   /api/account/create [post]
func (account) CreateAccount(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.CreateAccount{}
	if err := c.ShouldBindJSON(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.UserToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	result, err := logic.Group.Account.CreateAccount(c, content.ID, params.Name, global.PbSettings.Rule.DefaultAvatarURL, params.Gender, params.Signature)
	rly.Reply(err, result)
}

// GetAccountToken
// @Tags     account
// @Summary  获取账号令牌
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                                    true  "Bearer 用户令牌"
// @Param    data           query     request.GetAccountToken                   true  "账号ID"
// @Success  200            {object}  common.State{data=reply.GetAccountToken}  "1001:参数有误  1003:系统错误 2007:身份不存在 2008:身份验证失败 2009:权限不足 2010:账号不存在"
// @Router   /api/account/token [get]
func (account) GetAccountToken(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.GetAccountToken{}
	if err := c.ShouldBindQuery(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.UserToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	result, err := logic.Group.Account.GetAccountToken(c, content.ID, params.AccountID)
	rly.Reply(err, result)
}

// DeleteAccount
// @Tags     account
// @Summary  删除账户
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                 true  "Bearer 用户令牌"
// @Param    data           body      request.DeleteAccount  true  "账号ID"
// @Success  200            {object}  common.State{}         "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2009:权限不足 2010:账号不存在"
// @Router   /api/account/delete [delete]
func (account) DeleteAccount(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.DeleteAccount{}
	if err := c.ShouldBindJSON(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.UserToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	err := logic.Group.Account.DeleteAccount(c, content.ID, params.AccountID)
	rly.Reply(err)
}

// UpdateAccount
// @Tags     account
// @Summary  更新账户信息
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                                   true  "Bearer 账户令牌"
// @Param    data           body      request.UpdateAccount  true  "账号信息"
// @Success  200            {object}  common.State{}         "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败"
// @Router   /api/account/update [put]
func (account) UpdateAccount(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.UpdateAccount{}
	if err := c.ShouldBindJSON(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	err := logic.Group.Account.UpdateAccount(c, content.ID, params.Name, params.Avatar, params.Gender, params.Signature)
	rly.Reply(err)
}

// GetAccountByID
// @Tags     account
// @Summary  获取账户信息
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                 true  "Bearer 账户令牌"
// @Param    data           query     request.GetAccountByID                   true  "账号信息"
// @Success  200            {object}  common.State{data=reply.GetAccountByID}  "1001:参数有误 1003:系统错误 2009:权限不足 2007:身份不存在 2008:身份验证失败 2010:账号不存在"
// @Router   /api/account/info [get]
func (account) GetAccountByID(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.GetAccountByID{}
	if err := c.ShouldBindQuery(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	result, err := logic.Group.Account.GetAccountByID(c, params.AccountID, content.ID)
	rly.Reply(err, result)
}

// GetAccountsByUserID
// @Tags     account
// @Summary  获取用户的所有账户
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                                        true  "Bearer 用户令牌"
// @Success  200            {object}  common.State{data=reply.GetAccountsByUserID}  "1003:系统错误 2008:身份验证失败 2010:账号不存在"
// @Router   /api/account/infos/user [get]
func (account) GetAccountsByUserID(c *gin.Context) {
	rly := app.NewResponse(c)
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.UserToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	result, err := logic.Group.Account.GetAccountsByUserID(c, content.ID)
	rly.ReplyList(err, result.Total, result.List)
}

// GetAccountsByName
// @Tags     account
// @Summary  通过昵称模糊查找账户
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                                      true  "Bearer 账户令牌"
// @Param    data           query     request.GetAccountsByName                   true  "账号信息"
// @Success  200            {object}  common.State{data=reply.GetAccountsByName}  "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2010:账号不存在"
// @Router   /api/account/infos/name [get]
func (account) GetAccountsByName(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.GetAccountsByName{}
	if err := c.ShouldBindQuery(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	limit, offset := global.Page.GetPageSizeAndOffset(c.Request)
	result, err := logic.Group.Account.GetAccountsByName(c, content.ID, params.Name, limit, offset)
	rly.ReplyList(err, result.Total, result.List)
}
