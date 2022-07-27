package routing

import (
	v1 "github.com/0RAJA/chat_app/src/api/v1"
	"github.com/gin-gonic/gin"
)

type file struct{}

func (file) Init(router *gin.RouterGroup) {
	fg := router.Group("file")
	{
		fg.POST("publish", v1.Group.File.Publish)
		fg.POST("getall", v1.Group.File.GetRelationFile)
		fg.POST("avatar", v1.Group.File.UploadAvatar)
	}
}
