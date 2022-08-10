package task

import (
	"github.com/0RAJA/Rutils/pkg/utils"
	"github.com/0RAJA/chat_app/src/global"
	"github.com/0RAJA/chat_app/src/model/chat"
	"github.com/0RAJA/chat_app/src/model/chat/server"
)

// UpdateAccount 更新账号信息通知
// 参数: 账号ID，更新后的账号信息
func UpdateAccount(accessToken string, accountID int64, name, gender, signature string) func() {
	return func() {
		global.ChatMap.Send(accountID, chat.ServerUpdateAccount, server.UpdateAccount{
			EnToken:   utils.EncodeMD5(accessToken),
			Name:      name,
			Signature: signature,
			Gender:    gender,
		})
	}
}
