package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/0RAJA/chat_app/src/global"
	"github.com/0RAJA/chat_app/src/routing/router"
	"github.com/0RAJA/chat_app/src/setting"
	"github.com/gin-gonic/gin"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
}

// @title        chat
// @version      1.0
// @description  在线聊天系统

// @license.name  raja,chong
// @license.url

// @host      chat.humraja.xyz
// @BasePath  /api

// @securityDefinitions.basic  BasicAuth
func main() {
	setting.AllInit()
	if global.PbSettings.Server.RunMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := router.NewRouter() // 注册路由
	s := &http.Server{
		Addr:           global.PbSettings.Server.Address,
		Handler:        r,
		ReadTimeout:    global.PbSettings.Server.ReadTimeout,
		WriteTimeout:   global.PbSettings.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	global.Logger.Info("Server started!")
	fmt.Println("AppName:", global.PbSettings.App.Name, "Version:", global.PbSettings.App.Version, "Address:", global.PbSettings.Server.Address, "RunMode:", global.PbSettings.Server.RunMode)
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			global.Logger.Info(err.Error())
		}
	}()

	gracefulExit(s) // 优雅退出
	global.Logger.Info("Server exited!")
}

// 优雅退出
func gracefulExit(s *http.Server) {
	// 退出通知
	quit := make(chan os.Signal, 1)
	// 等待退出通知
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	global.Logger.Info("ShutDown Server...")
	// 给几秒完成剩余任务
	ctx, cancel := context.WithTimeout(context.Background(), global.PbSettings.Server.DefaultContextTimeout)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil { // 优雅退出
		global.Logger.Info("Server forced to ShutDown,Err:" + err.Error())
	}
}
