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

type notify struct {
}

// CreateNotify
// @Tags     notify
// @Summary  创建群通知
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                                 true  "Bearer 账户令牌"
// @Param    data           query     request.CreateNotify                   true  "请求信息"
// @Success  200            {object}  common.State{data=reply.GroupNotify}  "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2009:权限不足 7003:非群成员"
// @Router   /api/notify/create [post]
func (notify) CreateNotify(c *gin.Context) {
	rly := app.NewResponse(c)
	params := request.CreateNotify{}
	if err := c.ShouldBindQuery(&params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	result, mErr := logic.Group.Notify.CreateNotify(c, &params)

	rly.Reply(mErr, result)
}

// UpdateNotify
// @Tags     notify
// @Summary  更新群通知
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                                 true  "Bearer 账户令牌"
// @Param    data           query     request.UpdateNotify                   true  "请求信息"
// @Success  200            {object}  common.State{data=reply.UpdateNotify}  "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2009:权限不足 7003:非群成员"
// @Router   /api/notify/update [post]
func (notify) UpdateNotify(c *gin.Context) {
	rly := app.NewResponse(c)
	params := request.UpdateNotify{}
	if err := c.ShouldBindQuery(&params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	result, mErr := logic.Group.Notify.UpdateNotify(c, &params)

	rly.Reply(mErr, result)
}

// GetNotifyByID
// @Tags     notify
// @Summary  更新群通知
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                                 true  "Bearer 账户令牌"
// @Param    data           query     request.GetNotifyByID                 true  "请求信息"
// @Success  200            {object}  common.State{data=[]reply.GetNotify}  "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2009:权限不足 7003:非群成员"
// @Router   /api/notify/ [get]
func (notify) GetNotifyByID(c *gin.Context) {
	rly := app.NewResponse(c)
	params := request.GetNotifyByID{}
	if err := c.ShouldBindQuery(&params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	result, mErr := logic.Group.Notify.GetNotifyByID(c, params.RelationID, content.ID)

	rly.ReplyList(mErr, int64(len(result)), result)
}
