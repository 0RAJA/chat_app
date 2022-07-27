package logic

import (
	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/0RAJA/chat_app/src/dao"
	db "github.com/0RAJA/chat_app/src/dao/postgres/sqlc"
	"github.com/gin-gonic/gin"
)

type mGroup struct {

}

func (mGroup)CreateGroup(c *gin.Context,accountID int64,name string,desc string) (relationID int64,mErr errcode.Err) {
	relationID,err := dao.Group.DB.CreateGroupRelation(c,&db.CreateGroupRelationParams{
		Name:        name,
		Description: desc,
		Avatar:      "",
	})
	if err != nil {
		return 0,errcode.ErrServer
	}
	err = dao.Group.DB.CreateSetting(c,&db.CreateSettingParams{
		AccountID:  accountID,
		RelationID: relationID,
		IsLeader:   true,
		IsSelf:     false,
	})
	if err != nil {
		return 0,errcode.ErrServer
	}
	return relationID,nil
}
