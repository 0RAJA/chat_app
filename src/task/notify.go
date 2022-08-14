package task

import (
	"github.com/0RAJA/Rutils/pkg/utils"
	"github.com/0RAJA/chat_app/src/dao"
	"github.com/0RAJA/chat_app/src/global"
	"github.com/0RAJA/chat_app/src/model"
	"github.com/0RAJA/chat_app/src/model/chat"
	"github.com/0RAJA/chat_app/src/model/chat/server"
)

func CreateNotify(accessToken string, accountID, relationID int64, msgContent string, msgExtend *model.MsgExtend) func() {
	ctx, cancel := global.DefaultContextWithTimeOut()
	defer cancel()
	members, err := dao.Group.DB.GetGroupMembers(ctx, relationID)
	if err != nil {
		global.Logger.Error(err.Error())
	}
	return func() {
		global.ChatMap.SendMany(members, chat.ServerCreateNotify, server.CreateNotify{
			EnToken:    utils.EncodeMD5(accessToken),
			AccountID:  accountID,
			RelationID: relationID,
			MsgContent: msgContent,
			MsgExtent:  msgExtend,
		})
	}
}
func UpdateNotify(accessToken string, accountID, relationID int64, msgContent string, msgExtend *model.MsgExtend) func() {
	ctx, cancel := global.DefaultContextWithTimeOut()
	defer cancel()
	members, err := dao.Group.DB.GetGroupMembers(ctx, relationID)
	if err != nil {
		global.Logger.Error(err.Error())
	}
	return func() {
		global.ChatMap.SendMany(members, chat.ServerUpdateNotify, server.CreateNotify{
			EnToken:    utils.EncodeMD5(accessToken),
			AccountID:  accountID,
			RelationID: relationID,
			MsgContent: msgContent,
			MsgExtent:  msgExtend,
		})
	}
}
