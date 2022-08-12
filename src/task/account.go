package task

import (
	"github.com/0RAJA/Rutils/pkg/utils"
	"github.com/0RAJA/chat_app/src/dao"
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

// UpdateEmail 更新邮箱通知用户的每个账号
func UpdateEmail(accessToken string, userID int64, email string) func() {
	return func() {
		ctx, cancel := global.DefaultContextWithTimeOut()
		defer cancel()
		accountIDs, err := dao.Group.DB.GetAccountIDsByUserID(ctx, userID)
		if err != nil {
			global.Logger.Error(err.Error())
			return
		}
		global.ChatMap.SendMany(accountIDs, chat.ServerUpdateEmail, server.UpdateEmail{
			EnToken: utils.EncodeMD5(accessToken),
			Email:   email,
		})
	}
}
