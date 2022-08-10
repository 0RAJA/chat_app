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
		rlyInfo, merr := logic.GetMsgInfoByID(c, params.RlyMsgID)
		if merr != nil {
			return nil, merr
		}
		// 不能回复别的群的消息
		if rlyInfo.RelationID != params.RelationID {
			return nil, myerr.RlyMsgNotOneRelation
		}
		// 不能回复已经撤回的消息
		if rlyInfo.IsRevoke {
			return nil, myerr.RlyMsgHasRevoked
		}
		rlyMsgID = params.RlyMsgID
		rlyMsgExtend, err := model.JsonToExpand(rlyInfo.MsgExtend)
		if err != nil {
			global.Logger.Error(err.Error())
			return nil, errcode.ErrServer
		}
		rlyMsg = &reply.RlyMsg{
			MsgID:      rlyInfo.ID,
			MsgType:    rlyInfo.MsgType,
			MsgContent: rlyInfo.MsgContent,
			MsgExtend:  rlyMsgExtend,
			IsRevoke:   rlyInfo.IsRevoke,
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
	global.Worker.SendTask(task.PublishMsg(
		params.AccessToken,
		reply.MsgInfo{
			ID:         result.ID,
			NotifyType: string(db.MsgnotifytypeCommon),
			MsgType:    string(model.MsgTypeText),
			MsgContent: params.MsgContent,
			Extend:     params.MsgExtend,
			AccountID:  params.AccountID,
			RelationID: params.RelationID,
			CreateAt:   result.CreateAt,
		}, rlyMsg))
	return &client.HandleSendMsgRly{MsgID: result.ID, CreateAt: result.CreateAt}, nil
}

func (message) ReadMsg(c context.Context, params *model.HandleReadMsg) errcode.Err {
	// 判断权限
	ok, merr := logic.ExistsSetting(context.Background(), params.ReaderID, params.RelationID)
	if merr != nil {
		return merr
	}
	if !ok {
		return myerr.AuthPermissionsInsufficient
	}
	msgIDs, err := dao.Group.DB.UpdateMsgReads(c, &db.UpdateMsgReadsParams{
		Msgids:     params.MsgIDs,
		RelationID: params.RelationID,
		Accountid:  params.ReaderID,
	})
	if err != nil {
		global.Logger.Error(err.Error())
		return errcode.ErrServer
	}
	// 推送消息已经被读取
	global.Worker.SendTask(task.ReadMsg(
		params.AccessToken,
		params.RelationID,
		params.ReaderID,
		msgIDs,
	))
	return nil
}
