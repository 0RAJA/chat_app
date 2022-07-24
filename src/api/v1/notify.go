package v1

import (
	"github.com/0RAJA/Rutils/pkg/app"
	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/0RAJA/chat_app/src/logic"
	"github.com/0RAJA/chat_app/src/model/request"
	"github.com/gin-gonic/gin"
)

type notify struct {
}

func (notify)CreateNotify(c *gin.Context)  {
	rly := app.NewResponse(c)
	params := request.CreateNotify{}
	if err := c.ShouldBindQuery(&params);err!=nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}

	result,mErr := logic.Group.Notify.CreateNotify(c,&params)

	rly.Reply(mErr,result)
}

func (notify)UpdateNotify(c *gin.Context){
	rly := app.NewResponse(c)
	params := request.UpdateNotify{}
	if err := c.ShouldBindQuery(&params);err!=nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}

	result,mErr := logic.Group.Notify.UpdateNotify(c,&params)

	rly.Reply(mErr,result)
}

func (notify)GetNotifyByID(c *gin.Context)  {
	rly := app.NewResponse(c)
	params := request.UpdateNotify{}
	if err := c.ShouldBindQuery(&params);err!=nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	result,mErr := logic.Group.Notify.UpdateNotify(c,&params)

	rly.Reply(mErr,result)
}

