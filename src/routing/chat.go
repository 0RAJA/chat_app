package routing

import (
	v1 "github.com/0RAJA/chat_app/src/api/v1"
	chat2 "github.com/0RAJA/chat_app/src/model/chat"
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
	server.OnConnect("/", v1.Group.Chat.Handle.OnConnect)
	server.OnError("/", v1.Group.Chat.Handle.OnError)
	server.OnDisconnect("/", v1.Group.Chat.Handle.OnDisconnect)
	chatMessage(server)
}

func chatMessage(server *socketio.Server) {
	server.OnEvent("/", chat2.EventClientSendMsg, v1.Group.Chat.Message.SendMsg)
	server.OnEvent("/", chat2.EventClientReadMsg, v1.Group.Chat.Message.ReadMsg)
}
