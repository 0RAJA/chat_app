package model

type ClientSendMsgParams struct {
	RelationID int64      // 关系ID
	AccountID  int64      // 账户ID
	MsgContent string     // 消息内容
	MsgExtend  *MsgExtend // 消息扩展信息
	RlyMsgID   int64      // 回复消息ID (如果是回复消息，则此字段大于0)
}

type ClientReadMsgParams struct {
	MsgID     int64 // 消息ID
	AccountID int64 // 账户ID
}
