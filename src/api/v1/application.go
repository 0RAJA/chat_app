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

type application struct {
}

// Create
// @Tags     application
// @Summary  创建好友申请
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                     true  "Bearer 账户令牌"
// @Param    data           body      request.CreateApplication  true  "申请信息"
// @Success  200            {object}  common.State{}             "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 3001:申请已经存在 3003:申请不合法"
// @Router   /api/application/create [post]
func (application) Create(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.CreateApplication{}
	if err := c.ShouldBindJSON(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	err := logic.Group.Application.Create(c, content.ID, params.AccountID, params.ApplyMsg)
	rly.Reply(err)
}

// Delete
// @Tags     application
// @Summary  申请者删除好友申请
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                     true  "Bearer 账户令牌"
// @Param    data           body      request.DeleteApplication  true  "需要删除的申请"
// @Success  200            {object}  common.State{}             "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2009:权限不足 3002:申请不存在 3003:申请不合法"
// @Router   /api/application/delete [delete]
func (application) Delete(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.DeleteApplication{}
	if err := c.ShouldBindJSON(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	err := logic.Group.Application.Delete(c, content.ID, params.AccountID)
	rly.Reply(err)
}

// Accept
// @Tags     application
// @Summary  被申请者同意好友申请
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                     true  "Bearer 账户令牌"
// @Param    data           body      request.AcceptApplication  true  "需要同意的申请"
// @Success  200            {object}  common.State{}             "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 3002:申请不存在 3004:重复操作申请"
// @Router   /api/application/accept [put]
func (application) Accept(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.AcceptApplication{}
	if err := c.ShouldBindJSON(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	err := logic.Group.Application.Accept(c, content.ID, params.AccountID)
	rly.Reply(err)
}

// Refuse
// @Tags     application
// @Summary  被申请者拒绝好友申请
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                     true  "Bearer 账户令牌"
// @Param    data           body      request.AcceptApplication  true  "需要拒绝的申请"
// @Success  200            {object}  common.State{}             "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 3002:申请不存在 3004:重复操作申请"
// @Router   /api/application/refuse [put]
func (application) Refuse(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.RefuseApplication{}
	if err := c.ShouldBindJSON(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	err := logic.Group.Application.Refuse(c, content.ID, params.AccountID, params.RefuseMsg)
	rly.Reply(err)
}

// List
// @Tags     application
// @Summary  账户查看和自身相关的好友申请(不论是申请者还是被申请者)
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                                     true  "Bearer 账户令牌"
// @Param    data           query     request.ListApplications                   true  "分页参数"
// @Success  200            {object}  common.State{data=reply.ListApplications}  "1003:系统错误 2007:身份不存在 2008:身份验证失败"
// @Router   /api/application/list [get]
func (application) List(c *gin.Context) {
	rly := app.NewResponse(c)
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	limit, offset := global.Page.GetPageSizeAndOffset(c.Request)
	result, err := logic.Group.Application.List(c, content.ID, limit, offset)
	rly.Reply(err, result)
}
