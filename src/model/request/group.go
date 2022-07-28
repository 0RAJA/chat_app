package request

type CreateGroup struct {
	Name        string `json:"name" form:"name" binding:"required"`
	AccountID   int64  `json:"account_id" form:"account_id" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
}

type TransferGroup struct {
	RelationID    int64 `json:"relation_id" form:"relation_id" binding:"required"`
	FromAccountID int64 `json:"from_account_id" form:"from_account_id" binding:"required"`
	ToAccountID   int64 `json:"to_account_id" form:"to_account_id" binding:"required"`
}
type DissolveGroup struct {
	RelationID int64 `json:"relation_id" form:"relation_id" binding:"required"`
}

type UpdateGroup struct {
	RelationID  int64  `json:"relation_id" form:"relation_id" binding:"required"`
	Name        string `json:"name" form:"name" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
}