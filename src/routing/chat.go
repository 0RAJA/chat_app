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
	{
		server.OnConnect("/", v1.Group.Chat.Handle.OnConnect)
		server.OnError("/", v1.Group.Chat.Handle.OnError)
		server.OnDisconnect("/", v1.Group.Chat.Handle.OnDisconnect)
	}
	chatMessage(server)
	wg := router.Group("socket.io")
	{
		wg.GET("*any", gin.WrapH(server))
		wg.POST("*any", gin.WrapH(server))
	}
}

func chatMessage(server *socketio.Server) {
	server.OnEvent("/", chat2.ClientSendMsg, v1.Group.Chat.Message.SendMsg)
	server.OnEvent("/", chat2.ClientReadMsg, v1.Group.Chat.Message.ReadMsg)
}
