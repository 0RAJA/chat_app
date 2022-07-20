package v1

import (
	"fmt"
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

func (file) Publish(c *gin.Context)  {
	rly := app.NewResponse(c)
	params := request.PublishFile{}
	if err := c.ShouldBind(&params);err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	f,err := params.File.Open()
	if err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	fSrc ,err := ioutil.ReadAll(f)
	fileType := gtype.GetFileType(fSrc[:10])
	if err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	fmt.Println(fileType,params.File.Size,params.File.Filename)
	result,mErr := logic.Group.File.PublishFile(c,params,"img",params.File)

	rly.Reply(mErr,result)
}