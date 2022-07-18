package routing

import (
	v1 "github.com/0RAJA/chat_app/src/api/v1"
	"github.com/gin-gonic/gin"
)

type email struct {
}

func (email) Init(router *gin.RouterGroup) {
	eg := router.Group("email")
	{
		eg.GET("exist", v1.Group.Email.ExistEmail)
		eg.POST("send", v1.Group.Email.SendEmail)
	}
}
