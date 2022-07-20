package routing

import (
	v1 "github.com/0RAJA/chat_app/src/api/v1"
	mid "github.com/0RAJA/chat_app/src/middleware"
	"github.com/gin-gonic/gin"
)

type user struct {
}

func (user) Init(router *gin.RouterGroup) {
	ug := router.Group("user")
	{
		ug.POST("register", v1.Group.User.Register)
		ug.POST("login", v1.Group.User.Login)
		updateGroup := ug.Group("update").Use(mid.MustUser())
		{
			updateGroup.PUT("email", v1.Group.User.UpdateUserEmail)
			updateGroup.PUT("pwd", v1.Group.User.UpdateUserPassword)
			updateGroup.DELETE("delete", mid.MustUser(), v1.Group.User.DeleteUser)
		}
	}
}
