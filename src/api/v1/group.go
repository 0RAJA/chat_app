package v1

import (
	"github.com/0RAJA/Rutils/pkg/app"
	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/0RAJA/chat_app/src/logic"
	mid "github.com/0RAJA/chat_app/src/middleware"
	"github.com/0RAJA/chat_app/src/model"
	"github.com/0RAJA/chat_app/src/model/reply"
	"github.com/0RAJA/chat_app/src/model/request"
	"github.com/0RAJA/chat_app/src/myerr"
	"github.com/gin-gonic/gin"
)

type mGroup struct {
}

// CreateGroup
// @Tags     group
// @Summary  创建群
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                                true  "Bearer 账户令牌"
// @Param    data           query     request.CreateGroup                 true  "请求信息"
// @Success  200            {object}  common.State{data=reply.CreateGroup}  "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2009:权限不足"
// @Router   /api/group/create [post]
func (mGroup) CreateGroup(c *gin.Context) {

	rly := app.NewResponse(c)
	params := &request.CreateGroup{}
	if err := c.ShouldBindQuery(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}

	relationID, mErr := logic.Group.MGroup.CreateGroup(c, params.AccountID, params.Name, params.Description)
	if mErr != nil {
		rly.Reply(mErr)
		return
	}
	url, mErr := logic.Group.File.UploadGroupAvatar(c, nil, relationID)

	rly.Reply(mErr, reply.CreateGroup{
		Name:        params.Name,
		AccountID:   params.AccountID,
		RelationID:  relationID,
		Description: params.Description,
		Avatar:      url.Url,
	})
}

// TransferGroup
// @Tags     group
// @Summary  转让群
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                                 true  "Bearer 账户令牌"
// @Param    data           query     request.TransferGroup                 true  "请求信息"
// @Success  200            {object}  common.State{data=reply.TransferGroup}  "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2009:权限不足 7001:非群主 7003:非群成员"
// @Router   /api/group/transfer [post]
func (mGroup) TransferGroup(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.TransferGroup{}
	if err := c.ShouldBindQuery(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	result, mErr := logic.Group.MGroup.TransferGroup(c, params.RelationID, params.FromAccountID, params.ToAccountID)
	rly.Reply(mErr, result)
}

// DissolveGroup
// @Tags     group
// @Summary  解散群
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                                 true  "Bearer 账户令牌"
// @Param    data           query     request.DissolveGroup                 true  "请求信息"
// @Success  200            {object}  common.State{data=reply.DissolveGroup}  "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2009:权限不足 7001:非群主"
// @Router   /api/group/dissolve [post]
func (mGroup) DissolveGroup(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.DissolveGroup{}
	if err := c.ShouldBindQuery(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	result, mErr := logic.Group.MGroup.DissolveGroup(c, params.RelationID, content.ID)
	rly.Reply(mErr, result)
}

// UpdateGroup
// @Tags     group
// @Summary  更新群信息
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                                 true  "Bearer 账户令牌"
// @Param    data           query     request.UpdateGroup               true  "请求信息"
// @Success  200            {object}  common.State{data=reply.UpdateGroup}  "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2009:权限不足 7003:非群员"
// @Router   /api/group/update [post]
func (mGroup) UpdateGroup(c *gin.Context) {
	rly := app.NewResponse(c)
	params := request.UpdateGroup{}
	if err := c.ShouldBindQuery(&params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	result, mErr := logic.Group.MGroup.UpdateGroup(c, params, content.ID)
	rly.Reply(mErr, result)
}

// InviteAccount
// @Tags     group
// @Summary  邀请人进群
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                                 true  "Bearer 账户令牌"
// @Param    data           query     request.InviteAccount                true  "请求信息"
// @Success  200            {object}  common.State{data=reply.InviteAccount}  "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2009:权限不足 7003:非群员"
// @Router   /api/group/invite [post]
func (mGroup) InviteAccount(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.InviteAccount{}
	if err := c.ShouldBindQuery(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	result, mErr := logic.Group.MGroup.InviteAccount(c, params.RelationID, params.AccountID, content.ID)
	rly.Reply(mErr, result)
}

// QuitGroup
// @Tags     group
// @Summary  退群
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                                 true  "Bearer 账户令牌"
// @Param    data           query     request.QuitGroup                true  "请求信息"
// @Success  200            {object}  common.State{data=reply.QuitGroup}  "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2009:权限不足 7002:是群主 7003:非群员"
// @Router   /api/group/quit [post]
func (mGroup) QuitGroup(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.QuitGroup{}
	if err := c.ShouldBindQuery(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	result, mErr := logic.Group.MGroup.QuitGroup(c, params.RelationID, params.AccountID)
	rly.Reply(mErr, result)
}
