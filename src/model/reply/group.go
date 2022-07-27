package reply

type CreateGroup struct {
	Name        string `json:"name"`
	AccountID   int64  `json:"account_id"`
	RelationID   int64  `json:"relation_id"`
	Description string `json:"description"`
	Avatar      string `json:"avatar"`
}
