package routing

import (
	v1 "github.com/0RAJA/chat_app/src/api/v1"
	"github.com/gin-gonic/gin"
)

type mGroup struct {

}

func (mGroup) Init(router *gin.RouterGroup) {
	gg := router.Group("group")
	{
		gg.POST("create", v1.Group.MGroup.CreateGroup)
	}

}