package reply

import (
	"github.com/0RAJA/chat_app/src/model"
)

type GetFriends struct {
	List  []*model.SettingFriend `json:"list"`
	Total int64                  `json:"total"`
}

type GetPins struct {
	List  []*model.SettingPin `json:"list"`
	Total int64               `json:"total"`
}

type GetShows struct {
	List  []*model.Setting `json:"list"`
	Total int64            `json:"total"`
}

type GetFriendsByName struct {
	List  []*model.SettingFriend `json:"list"`
	Total int64                  `json:"total"`
}
