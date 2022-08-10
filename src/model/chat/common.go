package chat

// chat 中事件的关键字,以Server开头的事件为服务端发送的事件，以Client开头的事件为客户端发送的事件

// 服务端事件
const (
	ServerSendMsg       = "send_msg"       // 推送消息
	ServerReadMsg       = "read_msg"       // 已读消息
	ServerAccountLogin  = "account_login"  // 账户上线
	ServerAccountLogout = "account_logout" // 账户离线
	ServerUpdateAccount = "update_account" // 更新账户信息
	ServerApplication   = "application"    // 好友申请
)

// 客户端事件
const (
	ClientSendMsg = "send_msg" // 发送消息
	ClientReadMsg = "read_msg" // 已读消息
)
