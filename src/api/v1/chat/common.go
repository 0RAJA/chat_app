package chat

import (
	"context"
	"time"

	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/0RAJA/chat_app/src/dao"
	"github.com/0RAJA/chat_app/src/global"
	mid "github.com/0RAJA/chat_app/src/middleware"
	"github.com/0RAJA/chat_app/src/model"
	"github.com/0RAJA/chat_app/src/model/chat"
	"github.com/0RAJA/chat_app/src/model/common"
	"github.com/0RAJA/chat_app/src/myerr"
	socketio "github.com/googollee/go-socket.io"
)

// MustAccount 解析token并判断是否是账户，返回token
// 参数: accessToken
// 成功: 解析token的content并进行校验返回*model.Token,nil
// 失败: 返回 myerr.AuthenticationFailed,myerr.UserNotFound,errcode.ErrServer
func MustAccount(accessToken string) (*model.Token, errcode.Err) {
	payload, _, merr := mid.ParseHeader(accessToken)
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
		AccessToken: accessToken,
		Payload:     payload,
		Content:     content,
	}, nil
}

// CheckConnCtxToken 检查连接上下文中的token是否有效，有效返回token
// 参数: 连接上下文
// 成功: 上下文中包含 *model.Token 且有效
// 失败: 返回 myerr.AuthenticationFailed,myerr.AuthOverTime
func CheckConnCtxToken(v interface{}) (*model.Token, errcode.Err) {
	token, ok := v.(*model.Token)
	if !ok {
		return nil, myerr.AuthenticationFailed
	}
	if token.Payload.ExpiredAt.Before(time.Now()) {
		return nil, myerr.AuthOverTime
	}
	return token, nil
}

// CheckAuth 检查token是否有效，有效返回token，否则断开链接
func CheckAuth(s socketio.Conn) (*model.Token, bool) {
	token, merr := CheckConnCtxToken(s.Context())
	if merr != nil {
		s.Emit(chat.ServerError, common.NewState(merr))
		_ = s.Close()
		return nil, false
	}
	return token, true
}
