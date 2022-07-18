package v1

import (
	"github.com/0RAJA/Rutils/pkg/app"
	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/0RAJA/chat_app/src/logic"
	"github.com/0RAJA/chat_app/src/model/request"
	"github.com/gin-gonic/gin"
)

type email struct {
}

// ExistEmail 是否已经存在该email
// @Tags     email
// @Summary  是否已经存在该email
// @accept   application/json
// @Produce  application/json
// @Param    data  query     request.ExistEmail                   true  "email"
// @Success  200   {object}  common.State{data=reply.ExistEmail}  "是否存在该email"
// @Router   /api/email/exist [get]
func (email) ExistEmail(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.ExistEmail{}
	if err := c.ShouldBindQuery(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	result, err := logic.Group.Email.ExistEmail(c, params.Email)
	rly.Reply(err, result)
}

// SendEmail 发送邮件
// @Tags     email
// @Summary  发送邮件
// @accept   application/json
// @Produce  application/json
// @Param    data  body      request.SendEmail  true  "email"
// @Success  200   {object}  common.State{}     "发送情况"
// @Router   /api/email/send [post]
func (email) SendEmail(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.SendEmail{}
	if err := c.ShouldBindJSON(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	err := logic.Group.Email.SendEmail(c, params.Email)
	rly.Reply(err)
}
