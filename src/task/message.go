package task

import (
	"github.com/0RAJA/chat_app/src/dao"
	"github.com/0RAJA/chat_app/src/global"
	"github.com/0RAJA/chat_app/src/model/chat"
	"github.com/0RAJA/chat_app/src/model/chat/server"
	"github.com/0RAJA/chat_app/src/model/reply"
)

// 有关消息的推送任务

// PublishMsg 推送消息事件和执行拓展内容
// 参数: 消息和回复消息
func PublishMsg(msgInfo reply.MsgInfo, rlyMsg *reply.RlyMsg) func() {
	return func() {
		ctx, cancel := global.DefaultContextWithTimeOut()
		defer cancel()
		accountIDs, err := dao.Group.Redis.GetAccountsByRelationID(ctx, msgInfo.RelationID)
		if err != nil {
			global.Logger.Error(err.Error())
			return
		}
		global.ChatMap.SendMany(accountIDs, chat.EventServerSendMsg, server.SendMsg{
			MsgInfo: msgInfo,
			RlyMsg:  rlyMsg,
		})
	}
}

// PublishReadMsg 推送阅读消息事件
// 参数: 读者ID，消息发起者ID，消息ID
func PublishReadMsg(readerAccountID, accountID, msgID int64) func() {
	return func() {
		global.ChatMap.Send(accountID, chat.EventServerReadMsg, server.ReadMsg{
			MsgID:     msgID,
			AccountID: readerAccountID,
		})
	}
}
