package chat

import (
	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/0RAJA/chat_app/src/chat"
	"github.com/0RAJA/chat_app/src/global"
	"github.com/0RAJA/chat_app/src/model"
	chat2 "github.com/0RAJA/chat_app/src/model/chat/client"
	"github.com/0RAJA/chat_app/src/model/common"
	socketio "github.com/googollee/go-socket.io"
)

// 用于处理客户端发送的event
type message struct {
}

func (message) SendMsg(s socketio.Conn, msg string) string {
	token, merr := CheckConnCtxToken(s.Context())
	if merr != nil {
		return common.NewState(merr).JsonStr()
	}
	params := &chat2.HandleSendMsgParams{}
	if err := common.Decode(msg, params); err != nil {
		return common.NewState(errcode.ErrParamsNotValid.WithDetails(err.Error())).JsonStr()
	}
	c, cancel := global.DefaultContextWithTimeOut()
	defer cancel()
	result, merr := chat.Group.Message.SendMsg(c, &model.HandleSendMsg{
		RelationID: params.RelationID,
		AccountID:  token.Content.ID,
		MsgContent: params.MsgContent,
		MsgExtend:  params.MsgExtend,
		RlyMsgID:   params.RlyMsgID,
	})
	return common.NewState(merr, result).JsonStr()
}

func (message) ReadMsg(s socketio.Conn, msg string) string {
	token, merr := CheckConnCtxToken(s.Context())
	if merr != nil {
		return common.NewState(merr).JsonStr()
	}
	params := &chat2.HandleReadMsgParams{}
	if err := common.Decode(msg, params); err != nil {
		return common.NewState(errcode.ErrParamsNotValid.WithDetails(err.Error())).JsonStr()
	}
	c, cancel := global.DefaultContextWithTimeOut()
	defer cancel()
	merr = chat.Group.Message.ReadMsg(c, &model.HandleReadMsg{
		MsgID:     params.MsgID,
		AccountID: token.Content.ID,
	})
	return common.NewState(merr).JsonStr()
}
