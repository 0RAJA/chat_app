package reply

import "github.com/0RAJA/chat_app/src/model"

type CreateGroup struct {
	Name        string `json:"name"`
	AccountID   int64  `json:"account_id"`
	RelationID  int64  `json:"relation_id"`
	Description string `json:"description"`
	Avatar      string `json:"avatar"`
}
type TransferGroup struct {
}
type DissolveGroup struct {
}
type UpdateGroup struct {
	Name        string `json:"name" form:"name" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
}
type InviteAccount struct {
	InviteMember []int64 `json:"invite_member"`
}
type QuitGroup struct {
}
type GetGroup struct {
	List  []model.SettingGroup
	Total int64
}
type GetGroupMembers struct {
	ID       int64  `json:"account_id"` // 账号ID
	Name     string `json:"name"`       // 名称
	Avatar   string `json:"avatar"`     // 头像
	NickName string `json:"nick_name"`
	IsLeader bool   `json:"is_leader"`
}
