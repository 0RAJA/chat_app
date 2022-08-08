package chat

import (
	"context"
	"database/sql"

	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/0RAJA/chat_app/src/dao"
	db "github.com/0RAJA/chat_app/src/dao/postgres/sqlc"
	"github.com/0RAJA/chat_app/src/global"
	"github.com/0RAJA/chat_app/src/logic"
	"github.com/0RAJA/chat_app/src/model"
	"github.com/0RAJA/chat_app/src/model/chat/client"
	"github.com/0RAJA/chat_app/src/model/reply"
	"github.com/0RAJA/chat_app/src/myerr"
	"github.com/0RAJA/chat_app/src/task"
)

type message struct {
}

func (message) SendMsg(c context.Context, params *model.HandleSendMsg) (*client.HandleSendMsgRly, errcode.Err) {
	// 判断权限
	ok, merr := logic.ExistsSetting(context.Background(), params.AccountID, params.RelationID)
	if merr != nil {
		return nil, merr
	}
	if !ok {
		return nil, myerr.AuthPermissionsInsufficient
	}
	var rlyMsgID int64
	var rlyMsg *reply.RlyMsg
	// 判断回复消息
	if params.RlyMsgID > 0 {
		msgInfo, merr := logic.GetMsgInfoByID(c, params.RlyMsgID)
		if merr != nil {
			return nil, merr
		}
		// 不能回复别的群的消息
		if msgInfo.RelationID != params.RelationID {
			return nil, myerr.RlyMsgNotOneRelation
		}
		// 不能回复已经撤回的消息
		if msgInfo.IsRevoke {
			return nil, myerr.RlyMsgHasRevoked
		}
		rlyMsgID = params.RlyMsgID
		rlyMsgExtend, err := model.JsonToExpand(msgInfo.MsgExtend)
		if err != nil {
			global.Logger.Error(err.Error())
			return nil, errcode.ErrServer
		}
		rlyMsg = &reply.RlyMsg{
			MsgID:      msgInfo.ID,
			MsgType:    msgInfo.MsgType,
			MsgContent: msgInfo.MsgContent,
			MsgExtend:  rlyMsgExtend,
			IsRevoke:   msgInfo.IsRevoke,
		}
	}
	msgExtend, err := model.ExpandToJson(params.MsgExtend)
	if err != nil {
		global.Logger.Error(err.Error())
		return nil, errcode.ErrServer
	}
	result, err := dao.Group.DB.CreateMsg(context.Background(), &db.CreateMsgParams{
		NotifyType: db.MsgnotifytypeCommon,
		MsgType:    string(model.MsgTypeText),
		MsgContent: params.MsgContent,
		MsgExtend:  msgExtend,
		AccountID:  sql.NullInt64{Int64: params.AccountID, Valid: true},
		RlyMsgID:   sql.NullInt64{Int64: rlyMsgID, Valid: rlyMsgID > 0},
		RelationID: params.RelationID,
	})
	if err != nil {
		global.Logger.Error(err.Error())
		return nil, errcode.ErrServer
	}
	// 推送消息
	global.Worker.SendTask(task.PublishMsg(reply.MsgInfo{
		ID:         result.ID,
		NotifyType: string(db.MsgnotifytypeCommon),
		MsgType:    string(model.MsgTypeText),
		MsgContent: params.MsgContent,
		Extend:     params.MsgExtend,
		AccountID:  params.AccountID,
		RelationID: params.RelationID,
		CreateAt:   result.CreateAt,
	}, rlyMsg))
	return &client.HandleSendMsgRly{MsgID: result.ID}, nil
}

func (message) ReadMsg(c context.Context, params *model.HandleReadMsg) errcode.Err {
	msgInfo, merr := logic.GetMsgInfoByID(c, params.MsgID)
	if merr != nil {
		return merr
	}
	// 判断权限
	ok, merr := logic.ExistsSetting(context.Background(), params.AccountID, msgInfo.RelationID)
	if merr != nil {
		return merr
	}
	if !ok {
		return myerr.AuthPermissionsInsufficient
	}
	ok, err := dao.Group.DB.HasReadMsg(c, &db.HasReadMsgParams{
		MsgID:     params.MsgID,
		AccountID: params.AccountID,
	})
	if err != nil {
		global.Logger.Error(err.Error())
		return errcode.ErrServer
	}
	// 不能重复读
	if ok {
		return myerr.MsgAlreadyRead
	}
	if err := dao.Group.DB.UpdateMsgReads(c, &db.UpdateMsgReadsParams{
		ID:        params.MsgID,
		Accountid: params.AccountID,
	}); err != nil {
		global.Logger.Error(err.Error())
		return errcode.ErrServer
	}
	// 推送消息已经被读取
	global.Worker.SendTask(task.PublishReadMsg(params.AccountID, msgInfo.AccountID.Int64, params.MsgID))
	return nil
}
