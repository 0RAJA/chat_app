package client

import (
	"time"

	"github.com/0RAJA/chat_app/src/model"
)

// chat 中 client 端有关消息请求的结构

type HandleSendMsgParams struct {
	RelationID int64            `json:"relation_id" validate:"required,gte=1"` // 关系ID
	MsgContent string           `json:"msg_content" validate:"validate"`       // 消息内容
	MsgExtend  *model.MsgExtend `json:"msg_extend"`                            // 消息扩展信息
	RlyMsgID   int64            `json:"rly_msg_id"`                            // 回复消息ID (如果是回复消息，则此字段大于0)
}

type HandleSendMsgRly struct {
	MsgID    int64     `json:"msg_id"`    // 消息ID
	CreateAt time.Time `json:"create_at"` // 创建时间
}

type HandleReadMsgParams struct {
	MsgID int64 `json:"msg_id" validate:"required,gte=1"` // 消息ID
}
