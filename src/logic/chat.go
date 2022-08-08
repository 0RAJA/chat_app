package logic

import (
	"context"
	"database/sql"

	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/0RAJA/chat_app/src/dao"
	db "github.com/0RAJA/chat_app/src/dao/postgres/sqlc"
	"github.com/0RAJA/chat_app/src/global"
	"github.com/0RAJA/chat_app/src/model"
	chat2 "github.com/0RAJA/chat_app/src/model/chat"
	"github.com/0RAJA/chat_app/src/myerr"
)

// TODO: 聊天逻辑
type chat struct {
}

func (chat) ClientSendMsg(c context.Context, params *model.ClientSendMsgParams) (*chat2.ClientSendMsgRly, errcode.Err) {
	// 判断权限
	ok, merr := ExistsSetting(context.Background(), params.AccountID, params.RelationID)
	if merr != nil {
		return nil, merr
	}
	if !ok {
		return nil, myerr.AuthPermissionsInsufficient
	}
	var rlyMsgID int64
	// 判断回复消息
	if params.RlyMsgID > 0 {
		msgInfo, merr := GetMsgInfoByID(c, params.RlyMsgID)
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
	// TODO: 拓展消息通知
	// TODO: 消息推送
	return &chat2.ClientSendMsgRly{MsgID: result.ID}, nil
}

func (chat) ClientReadMsg(c context.Context, params *model.ClientReadMsgParams) errcode.Err {
	return nil
}
