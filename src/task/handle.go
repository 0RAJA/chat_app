package task

import (
	"github.com/0RAJA/Rutils/pkg/utils"
	"github.com/0RAJA/chat_app/src/global"
	"github.com/0RAJA/chat_app/src/model/chat"
	"github.com/0RAJA/chat_app/src/model/chat/server"
)

func AccountLogin(accessToken, address string, accountID int64) func() {
	return func() {
		global.ChatMap.Send(accountID, chat.ServerAccountLogin, server.AccountLogin{
			EnToken: utils.EncodeMD5(accessToken),
			Address: address,
		})
	}
}

func AccountLogout(accessToken, address string, accountID int64) func() {
	return func() {
		global.ChatMap.Send(accountID, chat.ServerAccountLogout, server.AccountLogout{
			EnToken: utils.EncodeMD5(accessToken),
			Address: address,
		})
	}
}
