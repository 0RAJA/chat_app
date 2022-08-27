package routing

import (
	v1 "github.com/0RAJA/chat_app/src/api/v1"
	chat2 "github.com/0RAJA/chat_app/src/model/chat"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
)

type ws struct {
}

func (ws) Init(router *gin.Engine) *socketio.Server {
	server := socketio.NewServer(nil)
	{
		server.OnConnect("/", v1.Group.Chat.Handle.OnConnect)
		server.OnError("/", v1.Group.Chat.Handle.OnError)
		server.OnDisconnect("/", v1.Group.Chat.Handle.OnDisconnect)
	}
	chatHandle(server)
	router.GET("socket.io/*any", gin.WrapH(server))
	router.POST("socket.io/*any", gin.WrapH(server))
	return server
}

func chatHandle(server *socketio.Server) {
	event := "/chat"
	server.OnEvent(event, chat2.ClientSendMsg, v1.Group.Chat.Message.SendMsg)
	server.OnEvent(event, chat2.ClientReadMsg, v1.Group.Chat.Message.ReadMsg)
	server.OnEvent(event, chat2.ClientTest, v1.Group.Chat.Handle.Test)
	server.OnEvent(event, chat2.ClientAuth, v1.Group.Chat.Handle.Auth)
}
