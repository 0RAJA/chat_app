package server

type DeleteRelation struct {
	EnToken    string `json:"en_token"`
	RelationID int64  `json:"relation_id"`
}

type UpdateNickName struct {
	EnToken    string `json:"en_token"`
	RelationID int64  `json:"relation_id"`
	NickName   string `json:"nick_name"`
}

type SettingType string

const (
	SettingPin        = "pin"
	SettingShow       = "show"
	SettingNotDisturb = "not_disturb"
)

type UpdateSettingState struct {
	EnToken    string      `json:"en_token"`    // 加密token
	Type       SettingType `json:"type"`        // 通知类型 [pin,show,disturb]
	RelationID int64       `json:"relation_id"` // 关系ID
	State      bool        `json:"state"`       // 状态设置
}
