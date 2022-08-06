package v1

import (
	"time"

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

type message struct {
}

// GetMsgsByRelationIDAndTime
// @Tags     message
// @Summary  通过最晚时间戳获取指定关系的信息，获取的消息按照发布时间先后排序
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                                               true  "Bearer 账户令牌"
// @Param    data           query     request.GetMsgsByRelationIDAndTime                   true  "请求信息"
// @Success  200            {object}  common.State{data=reply.GetMsgsByRelationIDAndTime}  "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2009:权限不足 2010:账号不存在"
// @Router   /api/msg/list/time [get]
func (message) GetMsgsByRelationIDAndTime(c *gin.Context) {
	rly := app.NewResponse(c)
	params := request.GetMsgsByRelationIDAndTime{}
	if err := c.ShouldBindQuery(&params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	limit, offset := global.Page.GetPageSizeAndOffset(c.Request)
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	result, err := logic.Group.Message.GetMsgsByRelationIDAndTime(c, model.GetMsgsByRelationIDAndTimeParams{
		AccountID:  content.ID,
		RelationID: params.RelationID,
		LastTime:   time.Unix(int64(params.LastTime), 0),
		Limit:      limit,
		Offset:     offset,
	})
	rly.ReplyList(err, result.Total, result.List)
}

// GetPinMsgsByRelationID
// @Tags     message
// @Summary  获取指定关系的pin消息，按照时间pin时间倒序排序
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                                           true  "Bearer 账户令牌"
// @Param    data           query     request.GetPinMsgsByRelationID                   true  "请求信息"
// @Success  200            {object}  common.State{data=reply.GetPinMsgsByRelationID}  "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2009:权限不足 2010:账号不存在"
// @Router   /api/msg/list/pin [get]
func (message) GetPinMsgsByRelationID(c *gin.Context) {
	rly := app.NewResponse(c)
	params := request.GetPinMsgsByRelationID{}
	if err := c.ShouldBindQuery(&params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	limit, offset := global.Page.GetPageSizeAndOffset(c.Request)
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	result, err := logic.Group.Message.GetPinMsgsByRelationID(c, model.GetPinMsgsByRelationIDParams{
		AccountID:  content.ID,
		RelationID: params.RelationID,
		Limit:      limit,
		Offset:     offset,
	})
	rly.ReplyList(err, result.Total, result.List)
}

// GetRlyMsgsInfoByMsgID
// @Tags     message
// @Summary  获取指定关系的所有回复消息，按照时间回复时间先后排序
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                                          true  "Bearer 账户令牌"
// @Param    data           query     request.GetRlyMsgsInfoByMsgID                   true  "请求信息"
// @Success  200            {object}  common.State{data=reply.GetRlyMsgsInfoByMsgID}  "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2009:权限不足 2010:账号不存在"
// @Router   /api/msg/list/rly [get]
func (message) GetRlyMsgsInfoByMsgID(c *gin.Context) {
	rly := app.NewResponse(c)
	params := request.GetRlyMsgsInfoByMsgID{}
	if err := c.ShouldBindQuery(&params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	limit, offset := global.Page.GetPageSizeAndOffset(c.Request)
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	result, err := logic.Group.Message.GetRlyMsgsInfoByMsgID(c, model.GetRlyMsgsInfoByMsgIDParams{
		AccountID:  content.ID,
		RelationID: params.RelationID,
		RlyMsgID:   params.RlyMsgID,
		Limit:      limit,
		Offset:     offset,
	})
	rly.ReplyList(err, result.Total, result.List)
}

// GetTopMsgByRelationID
// @Tags     message
// @Summary  获取指定关系的置顶消息，如果不存在则为null
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                                          true  "Bearer 账户令牌"
// @Param    data           query     request.GetTopMsgByRelationID                   true  "请求信息"
// @Success  200            {object}  common.State{data=reply.GetTopMsgByRelationID}  "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2009:权限不足 2010:账号不存在"
// @Router   /api/msg/info/top [get]
func (message) GetTopMsgByRelationID(c *gin.Context) {
	rly := app.NewResponse(c)
	params := request.GetTopMsgByRelationID{}
	if err := c.ShouldBindQuery(&params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	result, err := logic.Group.Message.GetTopMsgByRelationID(c, model.GetTopMsgByRelationIDParams{
		AccountID:  content.ID,
		RelationID: params.RelationID,
	})
	rly.Reply(err, result)
}

// UpdateMsgPin
// @Tags     message
// @Summary  设置消息的pin状态
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                true  "Bearer 账户令牌"
// @Param    data           body      request.UpdateMsgPin  true  "设置信息"
// @Success  200            {object}  common.State{}        "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2009:权限不足 2010:账号不存在 5001:消息不存在"
// @Router   /api/msg/update/pin [put]
func (message) UpdateMsgPin(c *gin.Context) {
	rly := app.NewResponse(c)
	params := request.UpdateMsgPin{}
	if err := c.ShouldBindJSON(&params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	err := logic.Group.Message.UpdateMsgPin(c, model.UpdateMsgPinParams{
		AccountID:  content.ID,
		RelationID: params.RelationID,
		MsgID:      params.ID,
		IsPin:      *params.IsPin,
	})
	rly.Reply(err)
}

// UpdateMsgTop
// @Tags     message
// @Summary  设置消息的置顶状态
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                true  "Bearer 账户令牌"
// @Param    data           body      request.UpdateMsgTop  true  "设置信息"
// @Success  200            {object}  common.State{}        "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2009:权限不足 2010:账号不存在 5001:消息不存在"
// @Router   /api/msg/update/top [put]
func (message) UpdateMsgTop(c *gin.Context) {
	rly := app.NewResponse(c)
	params := request.UpdateMsgTop{}
	if err := c.ShouldBindJSON(&params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	err := logic.Group.Message.UpdateMsgTop(c, model.UpdateMsgTopParams{
		AccountID:  content.ID,
		RelationID: params.RelationID,
		MsgID:      params.ID,
		IsTop:      *params.IsTop,
	})
	rly.Reply(err)
}

// RevokeMsg
// @Tags     message
// @Summary  撤回消息
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string             true  "Bearer 账户令牌"
// @Param    data           body      request.RevokeMsg  true  "信息id"
// @Success  200            {object}  common.State{}     "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2009:权限不足 2010:账号不存在 5001:消息不存在 5002:消息已经撤销"
// @Router   /api/msg/update/revoke [put]
func (message) RevokeMsg(c *gin.Context) {
	rly := app.NewResponse(c)
	params := request.RevokeMsg{}
	if err := c.ShouldBindJSON(&params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	err := logic.Group.Message.RevokeMsg(c, model.RevokeMsgParams{
		AccountID: content.ID,
		MsgID:     params.ID,
	})
	rly.Reply(err)
}

// CreateFileMsg
// @Tags      upload
// @Summary   发布文件消息
// @Security  BasicAuth
// @accept    multipart/form-data
// @Produce   application/json
// @Param     Authorization  header    string  true   "Bearer 账户令牌"
// @Param     file           formData  file    true   "文件"
// @Param     relation_id    body      int64   true   "关系id"
// @Param     rly_msg_id     body      int64   false  "回复消息id"
// @Success   200            {object}  common.State{reply.CreateFileMsg}
// @Router    /api/msg/file [post]
func (message) CreateFileMsg(c *gin.Context) {
	rly := app.NewResponse(c)
	params := request.CreateFileMsg{}
	if err := c.ShouldBind(&params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	result, err := logic.Group.Message.CreateFileMsg(c, model.CreateFileMsgParams{
		AccountID:  content.ID,
		RelationID: params.RelationID,
		RlyMsgID:   params.RlyMsgID,
		File:       params.File,
	})
	rly.Reply(err, result)
}
