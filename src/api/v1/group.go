package v1

import (
	"github.com/0RAJA/Rutils/pkg/app"
	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/0RAJA/chat_app/src/global"
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
// @Param    data           query     request.CreateGroup                   true  "请求信息"
// @Success  200            {object}  common.State{data=reply.CreateGroup}  "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2009:权限不足"
// @Router   /api/group/create [post]
func (mGroup) CreateGroup(c *gin.Context) {

	rly := app.NewResponse(c)
	params := &request.CreateGroup{}
	if err := c.ShouldBindQuery(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	relationID, mErr := logic.Group.MGroup.CreateGroup(c, content.ID, params.Name, params.Description)
	if mErr != nil {
		rly.Reply(mErr)
		return
	}
	url, mErr := logic.Group.File.UploadGroupAvatar(c, nil, relationID)

	rly.Reply(mErr, reply.CreateGroup{
		Name:        params.Name,
		AccountID:   content.ID,
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
// @Param    Authorization  header    string                                  true  "Bearer 账户令牌"
// @Param    data           query     request.TransferGroup                   true  "请求信息"
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
	result, mErr := logic.Group.MGroup.TransferGroup(c, params.RelationID, content.ID, params.ToAccountID)
	rly.Reply(mErr, result)
}

// DissolveGroup
// @Tags     group
// @Summary  解散群
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                                  true  "Bearer 账户令牌"
// @Param    data           query     request.DissolveGroup                   true  "请求信息"
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
// @Param    Authorization  header    string                                true  "Bearer 账户令牌"
// @Param    data           query     request.UpdateGroup                   true  "请求信息"
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
// @Param    Authorization  header    string                                  true  "Bearer 账户令牌"
// @Param    data           query     request.InviteAccount                   true  "请求信息"
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
// @Param    Authorization  header    string                             true  "Bearer 账户令牌"
// @Param    data           query     request.QuitGroup                   true  "请求信息"
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
	result, mErr := logic.Group.MGroup.QuitGroup(c, params.RelationID, content.ID)
	rly.Reply(mErr, result)
}

// GroupList 1
// @Tags     group
// @Summary  获取群列表
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                             true  "Bearer 账户令牌"
// @Success  200            {object}  common.State{data=reply.GetGroup}  "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2009:权限不足"
// @Router   /api/group/list [get]
func (mGroup) GroupList(c *gin.Context) {
	rly := app.NewResponse(c)
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	result, mErr := logic.Group.MGroup.GroupList(c, content.ID)
	rly.Reply(mErr, result)
}

// GetGroupByName
// @Tags     group
// @Summary  通过群名称模糊查找群
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                                    true  "Bearer 账户令牌"
// @Param    data           query     request.GetGroupByName             true  "请求信息"
// @Success  200            {object}  common.State{data=reply.GetGroup}  "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2009:权限不足 7003:非群员"
// @Router   /api/group/name [post]
func (mGroup) GetGroupByName(c *gin.Context) {
	rly := app.NewResponse(c)
	params := request.GetGroupByName{}
	if err := c.ShouldBindQuery(&params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	limit, offset := global.Page.GetPageSizeAndOffset(c.Request)
	result, mErr := logic.Group.MGroup.GetGroupByName(c, content.ID, limit, offset, params.Name)
	rly.Reply(mErr, result)
}

// GetGroupMembers
// @Tags     group
// @Summary  查看所有群员
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                              true  "Bearer 账户令牌"
// @Param    data           query     request.GetGroupMembers                   true  "请求信息"
// @Success  200            {object}  common.State{data=reply.GetGroupMembers}  "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2009:权限不足"
// @Router   /api/group/members [get]
func (mGroup) GetGroupMembers(c *gin.Context) {
	rly := app.NewResponse(c)
	params := request.GetGroupMembers{}
	if err := c.ShouldBindQuery(&params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	result, mErr := logic.Group.MGroup.GeGroupMembers(c, content.ID, params.RelationID)
	rly.Reply(mErr, result)
}
