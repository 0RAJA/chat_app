package chat

import (
	"context"
	"net/http"
	"time"

	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/0RAJA/chat_app/src/dao"
	"github.com/0RAJA/chat_app/src/global"
	mid "github.com/0RAJA/chat_app/src/middleware"
	"github.com/0RAJA/chat_app/src/model"
	"github.com/0RAJA/chat_app/src/model/common"
	"github.com/0RAJA/chat_app/src/model/request"
	"github.com/0RAJA/chat_app/src/myerr"
	socketio "github.com/googollee/go-socket.io"
)

// TODO: 聊天接收层
type chat struct {
}

// MustAccount 从header中获取并解析token并判断是否是账户，返回token
func MustAccount(header http.Header) (*model.Token, errcode.Err) {
	payload, merr := mid.ParseHeader(header)
	if merr != nil {
		return nil, merr
	}
	content := &model.Content{}
	if err := content.Unmarshal(payload.Content); err != nil {
		return nil, myerr.AuthenticationFailed
	}
	if content.Type != model.AccountToken {
		return nil, myerr.AuthenticationFailed
	}
	ok, err := dao.Group.DB.ExistsAccountByID(context.Background(), content.ID)
	if err != nil {
		global.Logger.Error(err.Error())
		return nil, errcode.ErrServer
	}
	if !ok {
		return nil, myerr.UserNotFound
	}
	return &model.Token{
		Payload: payload,
		Content: content,
	}, nil
}

// CheckConnCtxToken 检查连接上下文中的token是否有效
func CheckConnCtxToken(v interface{}) errcode.Err {
	token, ok := v.(*model.Token)
	if !ok {
		return myerr.AuthenticationFailed
	}
	if token.Payload.ExpiredAt.Before(time.Now()) {
		return myerr.AuthOverTime
	}
	return nil
}

func (chat) OnConnect(s socketio.Conn) error {
	token, err := MustAccount(s.RemoteHeader())
	if err != nil {
		return err
	}
	s.SetContext(token)
	return nil
}

func (chat) ClientSendMsg(s socketio.Conn, msg string) string {
	if merr := CheckConnCtxToken(s.Context()); merr != nil {
		return string(common.NewState(merr).MustJson())
	}
	params := &request.ClientSendMsg{}
	if err := common.Decode(msg, params); err != nil {
		return string(common.NewState(errcode.ErrParamsNotValid.WithDetails(err.Error())).MustJson())
	}
	// logic.Group.Chat.ClientSendMsg(params)
	return ""
}
