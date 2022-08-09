package routing

import (
	v1 "github.com/0RAJA/chat_app/src/api/v1"
	mid "github.com/0RAJA/chat_app/src/middleware"
	"github.com/gin-gonic/gin"
)

type setting struct {
}

func (setting) Init(router *gin.RouterGroup) {
	sg := router.Group("setting", mid.MustAccount())
	{
		sg.GET("pins", v1.Group.Setting.GetPins)
		sg.GET("shows", v1.Group.Setting.GetShows)
		updateGroup := sg.Group("update")
		{
			updateGroup.PUT("nick_name", v1.Group.Setting.UpdateNickName)
			updateGroup.PUT("pin", v1.Group.Setting.UpdateSettingPin)
			updateGroup.PUT("disturb", v1.Group.Setting.UpdateSettingDisturb)
			updateGroup.PUT("show", v1.Group.Setting.UpdateSettingShow)
		}
		friendGroup := sg.Group("friend")
		{
			friendGroup.GET("list", v1.Group.Setting.GetFriends)
			friendGroup.DELETE("delete", v1.Group.Setting.DeleteFriend)
			friendGroup.GET("name", v1.Group.Setting.GetFriendsByName)
		}
	}
}
