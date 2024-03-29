package routing

import (
	v1 "github.com/0RAJA/chat_app/src/api/v1"
	mid "github.com/0RAJA/chat_app/src/middleware"
	"github.com/gin-gonic/gin"
)

type file struct{}

func (file) Init(router *gin.RouterGroup) {
	fg := router.Group("file", mid.MustAccount())
	{
		//fg.POST("publish", v1.Group.File.Publish) //测试用
		//fg.POST("delete", v1.Group.File.DeleteFile) //测试用
		fg.POST("getall", v1.Group.File.GetRelationFile)
		fg.POST("avatar", v1.Group.File.UploadAvatar)
		fg.POST("details", v1.Group.File.GetFileDetailsByID)
	}
}
