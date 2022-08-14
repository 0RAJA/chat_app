package server

type TransferGroup struct {
	EnToken   string `json:"en_token"`
	AccountID int64  `json:"account_id"`
}
type DissolveGroup struct {
	EnToken    string `json:"en_token"`
	RelationID int64  `json:"relation_id"`
}
type InviteAccount struct {
	EnToken   string `json:"en_token"`
	AccountID int64  `json:"account_id"`
}
type QuitGroup struct {
	EnToken   string `json:"en_token"`
	AccountID int64  `json:"account_id"`
}
