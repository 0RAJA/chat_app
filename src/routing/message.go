package routing

import (
	v1 "github.com/0RAJA/chat_app/src/api/v1"
	mid "github.com/0RAJA/chat_app/src/middleware"
	"github.com/gin-gonic/gin"
)

type message struct {
}

func (message) Init(router *gin.RouterGroup) {
	mg := router.Group("msg", mid.MustAccount())
	{
		mg.POST("file", v1.Group.Message.CreateFileMsg)
		list := mg.Group("list")
		{
			list.GET("time", v1.Group.Message.GetMsgsByRelationIDAndTime)
			list.GET("feed", v1.Group.Message.FeedMsgsByAccountIDAndTime)
			list.GET("pin", v1.Group.Message.GetPinMsgsByRelationID)
			list.GET("rly", v1.Group.Message.GetRlyMsgsInfoByMsgID)
			list.GET("content", v1.Group.Message.GetMsgsByContent)
		}
		info := mg.Group("info")
		{
			info.GET("top", v1.Group.Message.GetTopMsgByRelationID)
		}
		update := mg.Group("update")
		{
			update.PUT("pin", v1.Group.Message.UpdateMsgPin)
			update.PUT("top", v1.Group.Message.UpdateMsgTop)
			update.PUT("revoke", v1.Group.Message.RevokeMsg)
		}
	}
}
