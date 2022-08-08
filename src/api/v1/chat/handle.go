package chat

import (
	"errors"

	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/0RAJA/chat_app/src/global"
	"github.com/0RAJA/chat_app/src/model"
	"github.com/0RAJA/chat_app/src/model/common"
	socketio "github.com/googollee/go-socket.io"
)

type handle struct {
}

func (handle) OnConnect(s socketio.Conn) error {
	token, err := MustAccount(s.RemoteHeader())
	if err != nil {
		return err
	}
	s.SetContext(token)
	// TODO: 通知其他账户
	// 加入在线群组
	global.Worker.SendTask(func() {
		global.ChatMap.Link(s, token.Content.ID)
	})
	return nil
}

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

func (handle) OnDisconnect(s socketio.Conn, _ string) {
	token, ok := s.Context().(*model.Token)
	if !ok {
		return
	}
	// TODO: 通知其他账户
	// 从在线中退出
	global.ChatMap.Leave(s, token.Content.ID)
}
