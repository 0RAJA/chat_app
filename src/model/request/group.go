package request

import "github.com/0RAJA/chat_app/src/model/common"

type CreateGroup struct {
	Name        string `json:"name" form:"name" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
}

type TransferGroup struct {
	RelationID  int64 `json:"relation_id" form:"relation_id" binding:"required"`
	ToAccountID int64 `json:"to_account_id" form:"to_account_id" binding:"required"`
}
type DissolveGroup struct {
	RelationID int64 `json:"relation_id" form:"relation_id" binding:"required"`
}

type UpdateGroup struct {
	RelationID  int64  `json:"relation_id" form:"relation_id" binding:"required"`
	Name        string `json:"name" form:"name" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
}

type InviteAccount struct {
	AccountID  int64 `json:"account_id,omitempty" form:"account_id" binding:"required"`
	RelationID int64 `json:"relation_id,omitempty" form:"relation_id" binding:"required"`
}

type QuitGroup struct {
	AccountID  int64 `json:"account_id,omitempty" form:"account_id" binding:"required"`
	RelationID int64 `json:"relation_id,omitempty" form:"relation_id" binding:"required"`
}
type GetGroupByName struct {
	Name string `json:"name" form:"name" binding:"required"`
	common.Pager
}
