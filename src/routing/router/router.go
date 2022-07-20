package router

import (
	"github.com/0RAJA/Rutils/pkg/app"
	_ "github.com/0RAJA/chat_app/docs"
	"github.com/0RAJA/chat_app/src/global"
	mid "github.com/0RAJA/chat_app/src/middleware"
	"github.com/0RAJA/chat_app/src/routing"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(mid.Cors(), mid.GinLogger(), mid.Recovery(true), mid.LogBody(), mid.Auth())
	root := r.Group("api")
	{
		root.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		root.GET("ping", func(c *gin.Context) {
			rly := app.NewResponse(c)
			global.Logger.Info("ping", mid.ErrLogMsg(c)...)
			rly.Reply(nil, "pang")
		})
		rg := routing.Group
		rg.User.Init(root)
		rg.Email.Init(root)
		rg.File.Init(root)
		rg.Account.Init(root)
		rg.Application.Init(root)
	}
	return r
}
