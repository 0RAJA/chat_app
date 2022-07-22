package routing

import (
	v1 "github.com/0RAJA/chat_app/src/api/v1"
	mid "github.com/0RAJA/chat_app/src/middleware"
	"github.com/gin-gonic/gin"
)

type application struct {
}

func (application) Init(router *gin.RouterGroup) {
	ag := router.Group("application", mid.MustAccount())
	{
		ag.POST("create", v1.Group.Application.Create)
		ag.DELETE("delete", v1.Group.Application.Delete)
		ag.PUT("accept", v1.Group.Application.Accept)
		ag.PUT("refuse", v1.Group.Application.Refuse)
		ag.GET("list", v1.Group.Application.List)
	}
}
