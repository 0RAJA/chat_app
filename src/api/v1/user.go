package v1

import (
	"github.com/0RAJA/Rutils/pkg/app"
	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/0RAJA/chat_app/src/logic"
	mid "github.com/0RAJA/chat_app/src/middleware"
	"github.com/0RAJA/chat_app/src/model"
	"github.com/0RAJA/chat_app/src/model/request"
	"github.com/0RAJA/chat_app/src/myerr"
	"github.com/gin-gonic/gin"
)

type user struct {
}

// Register 用户注册
// @Tags     user
// @Summary  用户注册
// @accept   application/json
// @Produce  application/json
// @Param    data  body      request.Register                   true  "用户注册信息"
// @Success  200   {object}  common.State{data=reply.Register}  "注册后user_token和用户信息"
// @Router   /api/user/register [post]
func (user) Register(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.Register{}
	if err := c.ShouldBindJSON(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	result, err := logic.Group.User.Register(c, params.Email, params.Password, params.Code)
	rly.Reply(err, result)
}

// Login 用户登录
// @Tags     user
// @Summary  用户登陆
// @accept   application/json
// @Produce  application/json
// @Param    data  body      request.Login                   true  "用户登陆信息"
// @Success  200   {object}  common.State{data=reply.Login}  "登陆后user_token和用户信息"
// @Router   /api/user/login [post]
func (user) Login(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.Login{}
	if err := c.ShouldBindJSON(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	result, err := logic.Group.User.Login(c, params.Email, params.Password)
	rly.Reply(err, result)
}

// UpdateUserEmail 更新用户邮箱
// @Tags      user
// @Summary   更新用户邮箱
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     Authorization  header    string                   true  "Bearer 用户令牌"
// @Param     data           body      request.UpdateUserEmail  true  "新邮箱和验证码"
// @Success   200            {object}  common.State{}
// @Router    /api/user/update/email [put]
func (user) UpdateUserEmail(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.UpdateUserEmail{}
	if err := c.ShouldBindJSON(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.UserToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	err := logic.Group.User.UpdateUserEmail(c, content.ID, params.Email, params.Code)
	rly.Reply(err)
}

// UpdateUserPassword 更新用户密码
// @Tags      user
// @Summary   更新用户密码
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     Authorization  header    string                      true  "Bearer 用户令牌"
// @Param     data           body      request.UpdateUserPassword  true  "旧密码和新密码"
// @Success   200            {object}  common.State{}
// @Router    /api/user/update/pwd [put]
func (user) UpdateUserPassword(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.UpdateUserPassword{}
	if err := c.ShouldBindJSON(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.UserToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	err := logic.Group.User.UpdateUserPassword(c, content.ID, params.OldPassword, params.NewPassword)
	rly.Reply(err)
}

// DeleteUser 删除当前用户
// @Tags      user
// @Summary   删除当前用户
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     Authorization  header    string  true  "Bearer 用户令牌"
// @Success   200            {object}  common.State{}
// @Router    /api/user/delete [delete]
func (user) DeleteUser(c *gin.Context) {
	rly := app.NewResponse(c)
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.UserToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	err := logic.Group.User.DeleteUser(c, content.ID)
	rly.Reply(err)
}
