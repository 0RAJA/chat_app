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
// @Summary  上传文件
// @accept   application/json
// @Param    data           body      request.PublishFile  true  "文件信息"
// @Success  200            {object}  common.State{data=reply.PublishFile}             "5001:存储失败"
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
	result, mErr := logic.Group.File.PublishFile(c, params, fileType, params.File)

	rly.Reply(mErr, result)
}

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

func (file) GetRelationFile(c *gin.Context) {
	rly := app.NewResponse(c)
	params := request.GetRelationFile{}
	if err := c.ShouldBindQuery(&params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	result,mErr := logic.Group.File.GetRelationFile(c,params.RelationID)

	rly.ReplyList(mErr,int64(len(result)),result)
}
