package routing

import (
	v1 "github.com/0RAJA/chat_app/src/api/v1"
	"github.com/gin-gonic/gin"
)

type notify struct {

}

func (notify)Init(router *gin.RouterGroup)  {
	r := router.Group("notify")
	{
		r.POST("create",v1.Group.Notify.CreateNotify)
		r.POST("update",v1.Group.Notify.UpdateNotify)
		r.POST("getnotify",v1.Group.Notify.GetNotifyByID)
	}
}