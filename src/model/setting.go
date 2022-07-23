package model

import (
	"time"
)

type SettingFriendInfo struct {
	AccountID int64  `json:"account_id"` // 好友ID
	Name      string `json:"name"`       // 好友姓名
	Avatar    string `json:"avatar"`     //	好友头像
}

type SettingGroupInfo struct {
}

type SettingInfo struct {
	RelationID   int64     `json:"relation_id"`    // 关系ID
	RelationType string    `json:"relation_type"`  // 关系类型['group','friend']
	NickName     string    `json:"nick_name"`      // 昵称(群组时为在群中昵称，好友时为好友昵称, 空表示未设置)
	IsNotDisturb bool      `json:"is_not_disturb"` // 是否免打扰
	IsPin        bool      `json:"is_pin"`         // 是否pin
	PinTime      time.Time `json:"pin_time"`       // pin时间
	IsShow       bool      `json:"is_show"`        // 是否显示
	LastShow     time.Time `json:"last_show"`      // 最后显示时间
}

type SettingPinInfo struct {
	RelationID   int64     `json:"relation_id"`   // 关系ID
	RelationType string    `json:"relation_type"` // 关系类型['group','friend']
	NickName     string    `json:"nick_name"`     // 昵称(群组时为在群中昵称，好友时为好友昵称, 空表示未设置)
	PinTime      time.Time `json:"pin_time"`      // pin时间
}

type SettingPin struct {
	SettingPinInfo
	GroupInfo  *SettingGroupInfo  `json:"group_info"`  // 群组信息
	FriendInfo *SettingFriendInfo `json:"friend_info"` // 好友信息
}

type setting struct {
	SettingInfo
	GroupInfo  *SettingGroupInfo  `json:"group_info"`  // 群组信息
	FriendInfo *SettingFriendInfo `json:"friend_info"` // 好友信息
}
