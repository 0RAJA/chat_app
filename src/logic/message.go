package logic

import (
	"time"

	"github.com/gin-gonic/gin"
)

type message struct {
}

func (message) GetMsgsByRelationIDAndTime(c *gin.Context, accountID int64, lastTime time.Time, limit, offset int32) {

}

func (message) GetPinMsgsByRelationID(c *gin.Context, rlyMsgID, accountID int64, limit, offset int32) {

}
