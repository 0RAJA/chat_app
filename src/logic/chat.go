package logic

import (
	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/0RAJA/chat_app/src/model"
	"github.com/0RAJA/chat_app/src/model/reply"
)

// TODO: 聊天逻辑
type chat struct {
}

func (chat) ClientSendMsg(params *model.ClientSendMsgParams) (*reply.ClientSendMsg, errcode.Err) {
	// msgExtend, err := model.ExpandToJson(params.MsgExtend)
	// if err != nil {
	// 	global.Logger.Error(err.Error())
	// 	return nil, errcode.ErrServer
	// }
	// ok, merr := ExistsSetting(context.Background(), params.AccountID, params.ID)
	// if merr != nil {
	// 	return nil, merr
	// }
	// if !ok {
	// 	return nil, myerr.AuthPermissionsInsufficient
	// }
	// dao.Group.DB.CreateMsg(context.Background(), &db.CreateMsgParams{
	// 	NotifyType: db.MsgnotifytypeCommon,
	// 	MsgType:    string(model.MsgTypeText),
	// 	MsgContent: params.MsgContent,
	// 	MsgExtend:  msgExtend,
	// 	FileID:     sql.NullInt64{},
	// 	AccountID:  sql.NullInt64{},
	// 	RlyMsgID:   sql.NullInt64{},
	// 	RelationID: 0,
	// })
	return nil, nil
}

func (chat) ClientReadMsg(params *model.ClientReadMsgParams) errcode.Err {
	return nil
}
