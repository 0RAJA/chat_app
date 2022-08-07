package request

import (
	"github.com/0RAJA/chat_app/src/model"
)

type ClientSendMsg struct {
	ID         int64            `json:"id" validate:"required,gte=1"`    // 目标ID，群ID或者账户ID
	IsFriend   *bool            `json:"is_friend" validate:"required"`   // 说明此次消息是否为好友消息
	MsgContent string           `json:"msg_content" validate:"validate"` // 消息内容
	MsgExtend  *model.MsgExtend `json:"msg_extend"`                      // 消息扩展信息
	RlyMsgID   int64            `json:"rly_msg_id"`                      // 回复消息ID (如果是回复消息，则此字段大于0)
}

type ClientReadMsg struct {
	MsgID int64 `json:"msg_id" validate:"required,gte=1"` // 消息ID
}
