package logic

import (
	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/0RAJA/chat_app/src/dao"
	db "github.com/0RAJA/chat_app/src/dao/postgres/sqlc"
	"github.com/0RAJA/chat_app/src/model/reply"
	"github.com/0RAJA/chat_app/src/model/request"
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
func (mGroup)TransferGroup(c *gin.Context,relationID int64,fID int64,tID int64) (reply.TransferGroup,errcode.Err) {
	err := dao.Group.DB.TransferIsSelfFalse(c,&db.TransferIsSelfFalseParams{
		RelationID: relationID,
		AccountID:  fID,
	})
	if err != nil {
		return reply.TransferGroup{},errcode.ErrServer
	}
	err = dao.Group.DB.TransferIsSelfTrue(c,&db.TransferIsSelfTrueParams{
		RelationID: relationID,
		AccountID:  tID,
	})
	if err != nil {
		return reply.TransferGroup{},errcode.ErrServer
	}
	return reply.TransferGroup{},nil
}
func (mGroup)DissolveGroup(c *gin.Context,relationId int64) (result reply.DissolveGroup,mErr errcode.Err)  {
	err := dao.Group.DB.DissolveGroup(c,relationId)
	if err != nil {
		return result,errcode.ErrServer
	}
	return result,nil
}
func (mGroup)UpdateGroup(c *gin.Context,params request.UpdateGroup) (result reply.UpdateGroup,mErr errcode.Err)  {
	err := dao.Group.DB.UpdateGroupRelation(c,&db.UpdateGroupRelationParams{
		Name:        params.Name,
		Description: params.Description,
		ID:          params.RelationID,
	})
	if err != nil {
		return result,errcode.ErrServer
	}
	result = reply.UpdateGroup{
		Name:        params.Name,
		Description: params.Description,
	}
	return result,nil
}