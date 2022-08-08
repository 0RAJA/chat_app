package chat

import (
	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/0RAJA/chat_app/src/logic"
	"github.com/0RAJA/chat_app/src/model"
	chat2 "github.com/0RAJA/chat_app/src/model/chat"
	"github.com/0RAJA/chat_app/src/model/common"
	socketio "github.com/googollee/go-socket.io"
)

// 用于处理客户端发送的event
type client struct {
}

func (client) OnConnect(s socketio.Conn) error {
	token, err := MustAccount(s.RemoteHeader())
	if err != nil {
		return err
	}
	// TODO: 通知其他账户
	s.SetContext(token)
	return nil
}

func (client) OnError(s socketio.Conn, e error) {

}

func (client) OnDisconnect(s socketio.Conn) {
	// TODO: 通知其他账户
}

func (client) SendMsg(s socketio.Conn, msg string) string {
	token, merr := CheckConnCtxToken(s.Context())
	if merr != nil {
		return common.NewState(merr).JsonStr()
	}
	params := &chat2.ClientSendMsgParams{}
	if err := common.Decode(msg, params); err != nil {
		return common.NewState(errcode.ErrParamsNotValid.WithDetails(err.Error())).JsonStr()
	}
	c, cancel := DefaultContextWithTimeOut()
	defer cancel()
	result, merr := logic.Group.Chat.ClientSendMsg(c, &model.ClientSendMsgParams{
		RelationID: params.RelationID,
		AccountID:  token.Content.ID,
		MsgContent: params.MsgContent,
		MsgExtend:  params.MsgExtend,
		RlyMsgID:   params.RlyMsgID,
	})
	return common.NewState(merr, result).JsonStr()
}

func (client) ReadMsg(s socketio.Conn, msg string) string {
	token, merr := CheckConnCtxToken(s.Context())
	if merr != nil {
		return common.NewState(merr).JsonStr()
	}
	params := &chat2.ClientReadMsgParams{}
	if err := common.Decode(msg, params); err != nil {
		return common.NewState(errcode.ErrParamsNotValid.WithDetails(err.Error())).JsonStr()
	}
	c, cancel := DefaultContextWithTimeOut()
	defer cancel()
	merr = logic.Group.Chat.ClientReadMsg(c, &model.ClientReadMsgParams{
		MsgID:     params.MsgID,
		AccountID: token.Content.ID,
	})
	return common.NewState(merr).JsonStr()
}
