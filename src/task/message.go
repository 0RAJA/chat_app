package task

import (
	"github.com/0RAJA/Rutils/pkg/utils"
	"github.com/0RAJA/chat_app/src/dao"
	"github.com/0RAJA/chat_app/src/global"
	"github.com/0RAJA/chat_app/src/model/chat"
	"github.com/0RAJA/chat_app/src/model/chat/server"
	"github.com/0RAJA/chat_app/src/model/reply"
)

// 有关消息的推送任务

// PublishMsg 推送消息事件和执行拓展内容
// 参数: 消息和回复消息
func PublishMsg(accessToken string, msgInfo reply.MsgInfo, rlyMsg *reply.RlyMsg) func() {
	return func() {
		ctx, cancel := global.DefaultContextWithTimeOut()
		defer cancel()
		accountIDs, err := dao.Group.Redis.GetAccountsByRelationID(ctx, msgInfo.RelationID)
		if err != nil {
			global.Logger.Error(err.Error())
			return
		}
		global.ChatMap.SendMany(accountIDs, chat.ServerSendMsg, server.SendMsg{
			EnToken: utils.EncodeMD5(accessToken),
			MsgInfo: reply.MsgInfoWithRly{
				MsgInfo: msgInfo,
				RlyMsg:  rlyMsg,
			},
		})
	}
}

// ReadMsg 推送阅读消息事件
// 参数: 读者ID，消息Map(accountID:[]msgID),所有msgIDs
func ReadMsg(accessToken string, readerID int64, msgMap map[int64][]int64, allMsgIDs []int64) func() {
	return func() {
		if len(msgMap) == 0 {
			return
		}
		enToken := utils.EncodeMD5(accessToken)
		// 给发送消息者推送
		for accountID, msgIDs := range msgMap {
			global.ChatMap.Send(accountID, chat.ServerReadMsg, server.ReadMsg{
				EnToken:  enToken,
				MsgIDs:   msgIDs,
				ReaderID: readerID,
			})
		}
		// 给自己的其他设备同步
		global.ChatMap.Send(readerID, chat.ServerSendMsg, server.ReadMsg{
			EnToken:  enToken,
			MsgIDs:   allMsgIDs,
			ReaderID: readerID,
		})
	}
}

func UpdateMsgState(accessToken string, relationID, msgID int64, msgType server.MsgType, state bool) func() {
	return func() {
		ctx, cancel := global.DefaultContextWithTimeOut()
		defer cancel()
		accountIDs, err := dao.Group.Redis.GetAccountsByRelationID(ctx, relationID)
		if err != nil {
			global.Logger.Error(err.Error())
			return
		}
		global.ChatMap.SendMany(accountIDs, chat.ServerUpdateMsgState, server.UpdateMsgState{
			EnToken: utils.EncodeMD5(accessToken),
			MsgType: msgType,
			MsgID:   msgID,
			State:   state,
		})
	}
}
