package v1

import (
	"github.com/0RAJA/Rutils/pkg/app"
	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/0RAJA/chat_app/src/logic"
	"github.com/0RAJA/chat_app/src/model/request"
	"github.com/0RAJA/chat_app/src/pkg/gtype"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type file struct {
}

// Publish
// @Tags     file
// @Summary  上传文件(测试用)
// @accept   multipart/form-data
// @Param    file  formData  request.PublishFile   true "文件"
// @Success  200            {object}  common.State{data=reply.PublishFile}             "1001:参数有误 1003:系统错误 8001:存储失败"
// @Router   /api/file/publish [post]
func (file) Publish(c *gin.Context) {
	rly := app.NewResponse(c)
	params := request.PublishFile{}
	if err := c.ShouldBind(&params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	f, err := params.File.Open()
	if err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	fSrc, err := ioutil.ReadAll(f)
	fileType := gtype.GetFileType(fSrc[:10])
	if err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	if fileType != "img" {
		fileType = "file"
	}
	//result, mErr := logic.Group.File.PublishFile(c, params, fileType)

	rly.Reply(nil)
}

// DeleteFile
// @Tags     file
// @Summary  删除文件(测试用)
// @accept   application/json
// @Param    data           body      request.DeleteFile  true  "文件ID"
// @Success  200            {object}  common.State{data=reply.DeleteFile}             "1001:参数有误 1003:系统错误 8002:文件不存在 8003文件删除失败"
// @Router   /api/file/publish [post]
func (file) DeleteFile(c *gin.Context) {
	rly := app.NewResponse(c)
	params := request.DeleteFile{}
	if err := c.ShouldBindQuery(&params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	result, mErr := logic.Group.File.DeleteFile(c, params.FileID)

	rly.Reply(mErr, result)
}

// GetRelationFile
// @Tags     file
// @Summary  获取关系文件列表
// @accept   application/json
// @Param    data           body      request.GetRelationFile  true  "关系ID"
// @Success  200            {object}  common.State{data=[]reply.File}             "1001:参数有误 1003:系统错误 8001:存储失败"
// @Router   /api/file/getall [post]
func (file) GetRelationFile(c *gin.Context) {
	rly := app.NewResponse(c)
	params := request.GetRelationFile{}
	if err := c.ShouldBindQuery(&params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	result, mErr := logic.Group.File.GetRelationFile(c, params.RelationID)

	rly.ReplyList(mErr, int64(len(result)), result)
}

// UploadAvatar
// @Tags     file
// @Summary  更新群头像活用户头像
// @accept   multipart/form-data
// @Param    file formData file true "文件"
// @Param    data           body      request.UploadAvatar  true  "文件及账号信息"
// @Success  200            {object}  common.State{data=reply.PublishFile}             "1001:参数有误 1003:系统错误 8001:存储失败"
// @Router   /api/file/avatar [post]
func (file) UploadAvatar(c *gin.Context) {
	rly := app.NewResponse(c)
	params := request.UploadAvatar{}
	if err := c.ShouldBind(&params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	if params.RelationID == 0 {
		result, mErr := logic.Group.File.UploadAccountAvatar(c, params.AccountID, params.File)
		rly.Reply(mErr, result)
		return
	}
	result, mErr := logic.Group.File.UploadGroupAvatar(c, params.File, params.RelationID)
	rly.Reply(mErr, result)
	return
}
