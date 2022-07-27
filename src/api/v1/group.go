package v1

import (
	"github.com/0RAJA/Rutils/pkg/app"
	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/0RAJA/chat_app/src/logic"
	"github.com/0RAJA/chat_app/src/model/reply"
	"github.com/0RAJA/chat_app/src/model/request"
	"github.com/gin-gonic/gin"
)

type mGroup struct {
}

func (mGroup) CreateGroup(c *gin.Context) {
	rly := app.NewResponse(c)
	params := request.CreateGroup{}
	if err := c.ShouldBind(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	relationID, mErr := logic.Group.MGroup.CreateGroup(c, params.AccountID, params.Name, params.Description)
	if mErr != nil {
		rly.Reply(mErr)
		return
	}
	mErr,url := logic.Group.File.UploadGroupAvatar(c, nil,relationID)

	rly.Reply(mErr, reply.CreateGroup{
		Name:        params.Name,
		AccountID:   params.AccountID,
		RelationID:  relationID,
		Description: params.Description,
		Avatar:      url,
	})
}
