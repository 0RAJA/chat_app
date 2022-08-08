package routing

import (
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
)

type ws struct {
}

func (ws) Init(router *gin.RouterGroup) {
	server := socketio.NewServer(nil)
	wg := router.Group("socket.io")
	{
		wg.GET("*any", gin.WrapH(server))
		wg.POST("*any", gin.WrapH(server))
	}
}
