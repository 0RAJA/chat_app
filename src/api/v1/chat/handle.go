package chat

import (
	"errors"

	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/0RAJA/chat_app/src/global"
	"github.com/0RAJA/chat_app/src/model"
	"github.com/0RAJA/chat_app/src/model/common"
	"github.com/0RAJA/chat_app/src/task"
	socketio "github.com/googollee/go-socket.io"
)

type handle struct {
}

// OnConnect
// 当客户端连接时触发
func (handle) OnConnect(s socketio.Conn) error {
	token, err := MustAccount(s.RemoteHeader())
	if err != nil {
		return err
	}
	s.SetContext(token)
	// 加入在线群组
	global.ChatMap.Link(s, token.Content.ID)
	// 通知其他设备
	global.Worker.SendTask(task.AccountLogin(token.AccessToken, s.RemoteAddr().String(), token.Content.ID))
	return nil
}

// OnError
// 当发生错误时触发
func (handle) OnError(s socketio.Conn, e error) {
	if s == nil {
		return
	}
	var merr errcode.Err
	if !errors.As(e, &merr) {
		global.Logger.Error(e.Error())
		merr = errcode.ErrServer
	}
	// TODO: 发送错误描述
	s.Emit("error", common.NewState(merr))
}

// OnDisconnect
// 当客户端断开连接时触发
func (handle) OnDisconnect(s socketio.Conn, _ string) {
	token, ok := s.Context().(*model.Token)
	if !ok {
		return
	}
	// TODO: 通知其他账户
	// 从在线中退出
	global.ChatMap.Leave(s, token.Content.ID)
}
