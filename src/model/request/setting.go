package request

import (
	"github.com/0RAJA/chat_app/src/model/common"
)

type DeleteFriend struct {
	RelationID int64 `json:"relation_id" binding:"required,gte=1"` // 关系ID
}

type UpdateNickName struct {
	RelationID int64  `json:"relation_id" binding:"required,gte=1"`      // 关系ID
	NickName   string `json:"nick_name" binding:"required,gte=1,lte=20"` // 昵称
}

type UpdateSettingPin struct {
	RelationID int64 `json:"relation_id" binding:"required,gte=1"` // 关系ID
	IsPin      *bool `json:"is_pin" binding:"required"`            // 是否pin
}

type UpdateSettingDisturb struct {
	RelationID   int64 `json:"relation_id" binding:"required,gte=1"` // 关系ID
	IsNotDisturb *bool `json:"is_not_disturb" binding:"required"`    // 是否免打扰
}

type UpdateSettingShow struct {
	RelationID int64 `json:"relation_id" binding:"required,gte=1"` // 关系ID
	IsShow     *bool `json:"is_show" binding:"required"`           // 是否展示
}

type GetFriendsByName struct {
	Name string `form:"name" binding:"required,gte=1,lte=20"` // 查询名称
	common.Pager
}
