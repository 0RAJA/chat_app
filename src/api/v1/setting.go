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

type setting struct {
}

// GetFriends
// @Tags     setting
// @Summary  获取当前账户的好友列表
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                               true  "Bearer 账户令牌"
// @Success  200            {object}  common.State{data=reply.GetFriends}  "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2010:账号不存在"
// @Router   /api/setting/friend/list [get]
func (setting) GetFriends(c *gin.Context) {
	rly := app.NewResponse(c)
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	result, err := logic.Group.Setting.GetFriends(c, content.ID)
	rly.Reply(err, result)
}

// GetPins
// @Tags     setting
// @Summary  获取当前账户pin的好友和群组列表
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                            true  "Bearer 账户令牌"
// @Success  200            {object}  common.State{data=reply.GetPins}  "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2010:账号不存在 2009:权限不足 3002:申请不存在 3003:申请不合法"
// @Router   /api/setting/pins [get]
func (setting) GetPins(c *gin.Context) {
	rly := app.NewResponse(c)
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	result, err := logic.Group.Setting.GetPins(c, content.ID)
	rly.Reply(err, result)
}

// GetShows
// @Tags     setting
// @Summary  获取当前账户首页显示的好友和群组列表(TODO: 待完善)
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                             true  "Bearer 账户令牌"
// @Success  200            {object}  common.State{data=reply.GetShows}  "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2010:账号不存在 2009:权限不足 3002:申请不存在 3003:申请不合法"
// @Router   /api/setting/shows [get]
func (setting) GetShows(c *gin.Context) {
	rly := app.NewResponse(c)
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	result, err := logic.Group.Setting.GetShows(c, content.ID)
	rly.Reply(err, result)
}

// DeleteFriend
// @Tags     setting
// @Summary  删除好友关系(双向删除)
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                true  "Bearer 账户令牌"
// @Param    data           body      request.DeleteFriend  true  "关系ID"
// @Success  200            {object}  common.State{}        "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2009:权限不足 2010:账号不存在 4002:关系不存在"
// @Router   /api/setting/friend/delete [delete]
func (setting) DeleteFriend(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.DeleteFriend{}
	if err := c.ShouldBindJSON(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	err := logic.Group.Setting.DeleteFriend(c, content.ID, params.RelationID)
	rly.Reply(err)
}

// UpdateNickName
// @Tags     setting
// @Summary  更新好友备注或群昵称
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                  true  "Bearer 账户令牌"
// @Param    data           body      request.UpdateNickName  true  "关系ID，备注或群昵称"
// @Success  200            {object}  common.State{}          "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2010:账号不存在 4002:关系不存在"
// @Router   /api/setting/update/nick_name [put]
func (setting) UpdateNickName(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.UpdateNickName{}
	if err := c.ShouldBindJSON(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	err := logic.Group.Setting.UpdateNickName(c, content.ID, params.RelationID, params.NickName)
	rly.Reply(err)
}

// UpdateSettingPin
// @Tags     setting
// @Summary  更新好友或群组pin选项
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                    true  "Bearer 账户令牌"
// @Param    data           body      request.UpdateSettingPin  true  "关系ID，pin状态"
// @Success  200            {object}  common.State{}            "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2010:账号不存在 4002:关系不存在"
// @Router   /api/setting/update/pin [put]
func (setting) UpdateSettingPin(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.UpdateSettingPin{}
	if err := c.ShouldBindJSON(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	err := logic.Group.Setting.UpdateSettingPin(c, content.ID, params.RelationID, *params.IsPin)
	rly.Reply(err)
}

// UpdateSettingDisturb
// @Tags     setting
// @Summary  更改好友或群组免打扰选项
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                        true  "Bearer 账户令牌"
// @Param    data           body      request.UpdateSettingDisturb  true  "关系ID，免打扰状态"
// @Success  200            {object}  common.State{}                "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2010:账号不存在 4002:关系不存在"
// @Router   /api/setting/update/disturb [put]
func (setting) UpdateSettingDisturb(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.UpdateSettingDisturb{}
	if err := c.ShouldBindJSON(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	err := logic.Group.Setting.UpdateSettingDisturb(c, content.ID, params.RelationID, *params.IsNotDisturb)
	rly.Reply(err)
}

// GetFriendsByName
// @Tags     setting
// @Summary  通过姓名模糊查询好友或群组
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                                     true  "Bearer 账户令牌"
// @Param    data           query     request.GetFriendsByName                   true  "关系ID，免打扰状态"
// @Success  200            {object}  common.State{data=reply.GetFriendsByName}  "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2010:账号不存在"
// @Router   /api/setting/friend/name [get]
func (setting) GetFriendsByName(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.GetFriendsByName{}
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
	result, err := logic.Group.Setting.GetFriendsByName(c, content.ID, params.Name, limit, offset)
	rly.Reply(err, result)
}
