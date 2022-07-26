package routing

import (
	v1 "github.com/0RAJA/chat_app/src/api/v1"
	mid "github.com/0RAJA/chat_app/src/middleware"
	"github.com/gin-gonic/gin"
)

type account struct {
}

func (account) Init(router *gin.RouterGroup) {
	ag := router.Group("account")
	{
		userGroup := ag.Group("").Use(mid.MustUser())
		{
			userGroup.POST("create", v1.Group.Account.CreateAccount)
			userGroup.GET("token", v1.Group.Account.GetAccountToken)
			userGroup.DELETE("delete", v1.Group.Account.DeleteAccount)
			userGroup.GET("infos/user", v1.Group.Account.GetAccountsByUserID)
		}
		accountGroup := ag.Group("").Use(mid.MustAccount())
		{
			accountGroup.PUT("update", v1.Group.Account.UpdateAccount)
			accountGroup.GET("infos/name", v1.Group.Account.GetAccountsByName)
			accountGroup.GET("info", v1.Group.Account.GetAccountByID)
		}
	}
}
