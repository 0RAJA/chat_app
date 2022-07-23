package logic

import (
	"github.com/0RAJA/chat_app/src/dao"
	"github.com/gin-gonic/gin"
)

type setting struct {
}

func (setting) GetFriendSettings(c *gin.Context, accountID int64) {
	dao.Group.DB.GetFriendSettingsOrderByName(c, accountID)
}
