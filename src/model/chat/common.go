package chat

// chat 中事件的关键字,以Server开头的事件为服务端发送的事件，以Client开头的事件为客户端发送的事件

const (
	ServerSendMsg      = "send_msg"
	ServerReadMsg      = "read_msg"
	ServerAccountLogin = "account_login"
)

const (
	ClientSendMsg = "send_msg"
	ClientReadMsg = "read_msg"
)
