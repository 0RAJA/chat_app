package request

type CreateNotify struct {
	RelationID int64  `json:"relation_id" form:"relation_id" binding:"required"`
	MsgContent string `json:"msg_content" form:"msg_content" binding:"required"`
	MsgExpand  string `json:"msg_expand"  form:"msg_expand" binding:"required"`
	AccountID  int64  `json:"account_id" form:"account_id" binding:"required"`
}
type UpdateNotify struct {
	ID         int64  `json:"id" form:"id" binding:"required"`
	RelationID int64  `json:"relation_id" form:"relation_id" binding:"required"`
	MsgContent string `json:"msg_content" form:"msg_content" binding:"required"`
	MsgExpand  string ` json:"msg_expand"  form:"msg_expand" binding:"required"`
	AccountID  int64  `json:"account_id" form:"account_id" binding:"required"`
}
type GetNotifyByID struct {
	RelationID int64 `json:"relation_id" form:"relation_id"`
}
