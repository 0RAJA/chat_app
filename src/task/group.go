package task

import (
	"github.com/0RAJA/Rutils/pkg/utils"
	"github.com/0RAJA/chat_app/src/dao"
	"github.com/0RAJA/chat_app/src/global"
	"github.com/0RAJA/chat_app/src/model/chat"
	"github.com/0RAJA/chat_app/src/model/chat/server"
)

func TransferGroup(accessToken string, accountID, relationID int64) func() {
	ctx, cancel := global.DefaultContextWithTimeOut()
	defer cancel()
	members, err := dao.Group.DB.GetGroupMembers(ctx, relationID)
	if err != nil {
		global.Logger.Error(err.Error())
	}
	return func() {
		global.ChatMap.SendMany(members, chat.ServerGroupTransferred, server.TransferGroup{
			EnToken:   utils.EncodeMD5(accessToken),
			AccountID: accountID,
		})
	}
}

func DissolveGroup(accessToken string, relationID int64) func() {
	ctx, cancel := global.DefaultContextWithTimeOut()
	defer cancel()
	members, err := dao.Group.DB.GetGroupMembers(ctx, relationID)
	if err != nil {
		global.Logger.Error(err.Error())
	}
	return func() {
		global.ChatMap.SendMany(members, chat.ServerGroupDissolved, server.DissolveGroup{
			EnToken:    utils.EncodeMD5(accessToken),
			RelationID: relationID,
		})
	}
}

func InviteAccount(accessToken string, relationID int64, accountID int64) func() {
	ctx, cancel := global.DefaultContextWithTimeOut()
	defer cancel()
	members, err := dao.Group.DB.GetGroupMembers(ctx, relationID)
	if err != nil {
		global.Logger.Error(err.Error())
	}
	return func() {
		global.ChatMap.SendMany(members, chat.ServerInviteAccount, server.InviteAccount{
			EnToken:   utils.EncodeMD5(accessToken),
			AccountID: accountID,
		})
	}
}

func QuitGroup(accessToken string, relationID int64, accountID int64) func() {
	ctx, cancel := global.DefaultContextWithTimeOut()
	defer cancel()
	members, err := dao.Group.DB.GetGroupMembers(ctx, relationID)
	if err != nil {
		global.Logger.Error(err.Error())
	}
	return func() {
		global.ChatMap.SendMany(members, chat.ServerQuitGroup, server.QuitGroup{
			EnToken:   utils.EncodeMD5(accessToken),
			AccountID: accountID,
		})
	}
}
