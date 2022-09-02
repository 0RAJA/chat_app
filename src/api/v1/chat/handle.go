package chat

import (
	"log"
	"time"

	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/0RAJA/chat_app/src/global"
	"github.com/0RAJA/chat_app/src/model/chat/client"
	"github.com/0RAJA/chat_app/src/model/common"
	"github.com/0RAJA/chat_app/src/task"
	socketio "github.com/googollee/go-socket.io"
)

type handle struct {
}

// OnConnect
// 当客户端连接时触发
func (handle) OnConnect(s socketio.Conn) error {
	log.Println("connected:", s.RemoteAddr().String(), s.ID())
	// 一定时间内需要进行AUTH认证，否则断开连接
	time.AfterFunc(global.PbSettings.Server.DefaultContextTimeout, func() {
		if s.Context() == nil {
			_ = s.Close()
		}
	})
	return nil
}

// OnError
// 当发生错误时触发
func (handle) OnError(s socketio.Conn, e error) {
	log.Println("conn err", e)
	if s == nil {
		return
	}
	// 从在线中退出
	global.ChatMap.Leave(s)
	log.Println("disconnected:", s.RemoteAddr().String(), s.ID())
	_ = s.Close()
}

// OnDisconnect
// 当客户端断开连接时触发
func (handle) OnDisconnect(s socketio.Conn, _ string) {
	// 从在线中退出
	global.ChatMap.Leave(s)
	log.Println("disconnected:", s.RemoteAddr().String(), s.ID())
}

// Auth 身份验证
func (handle) Auth(s socketio.Conn, accessToken string) string {
	token, merr := MustAccount(accessToken)
	if merr != nil {
		return common.NewState(merr).JsonStr()
	}
	s.SetContext(token)
	// 加入在线群组
	global.ChatMap.Link(s, token.Content.ID)
	// 通知其他设备
	global.Worker.SendTask(task.AccountLogin(token.AccessToken, s.RemoteAddr().String(), token.Content.ID))
	log.Println("auth accept:", s.RemoteAddr().String())
	return common.NewState(nil).JsonStr()
}

// Test 测试
func (handle) Test(s socketio.Conn, msg string) string {
	_, ok := CheckAuth(s)
	if !ok {
		return ""
	}
	params := &client.TestParams{}
	log.Println(msg)
	if err := common.Decode(msg, params); err != nil {
		return common.NewState(errcode.ErrParamsNotValid.WithDetails(err.Error())).JsonStr()
	}
	result := common.NewState(nil, client.TestRly{
		Name:    params.Name,
		Age:     params.Age,
		Address: s.RemoteAddr().String(),
		ID:      s.ID(),
	}).JsonStr()
	// test
	s.Emit("test", "test")
	return result
}
