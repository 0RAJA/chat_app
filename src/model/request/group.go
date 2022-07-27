package request

type CreateGroup struct {
	Name        string               `json:"name" form:"name" binding:"required"`
	AccountID   int64                `json:"account_id" form:"account_id" binding:"required"`
	Description string               `json:"description" form:"description" binding:"required"`
}
