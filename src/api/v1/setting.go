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

// TODO:增加api文档

type setting struct {
}

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

func (setting) Delete(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.Delete{}
	if err := c.ShouldBindJSON(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	err := logic.Group.Setting.Delete(c, content.ID, params.RelationID)
	rly.Reply(err)
}

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

func (setting) GetFriendsByName(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.GetFriendSettingsByName{}
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
