package logic

import (
	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/0RAJA/chat_app/src/dao"
	db "github.com/0RAJA/chat_app/src/dao/postgres/sqlc"
	"github.com/0RAJA/chat_app/src/global"
	"github.com/0RAJA/chat_app/src/model/reply"
	"github.com/0RAJA/chat_app/src/model/request"
	"github.com/0RAJA/chat_app/src/myerr"
	"github.com/gin-gonic/gin"
)

type mGroup struct {
}

func (mGroup) CreateGroup(c *gin.Context, accountID int64, name string, desc string) (relationID int64, mErr errcode.Err) {
	relationID, err := dao.Group.DB.CreateGroupRelation(c, &db.CreateGroupRelationParams{
		Name:        name,
		Description: desc,
		Avatar:      global.PbSettings.Rule.DefaultAvatarURL,
	})
	if err != nil {
		return 0, errcode.ErrServer
	}
	err = dao.Group.DB.CreateSetting(c, &db.CreateSettingParams{
		AccountID:  accountID,
		RelationID: relationID,
		IsLeader:   true,
		IsSelf:     false,
	})
	if err != nil {
		return 0, errcode.ErrServer
	}
	err = dao.Group.Redis.AddGroupRelationAccount(c, relationID, accountID)
	if err != nil {
		return 0, errcode.ErrServer
	}
	return relationID, nil
}
func (mGroup) TransferGroup(c *gin.Context, relationID int64, fID int64, tID int64) (reply.TransferGroup, errcode.Err) {
	t, err := dao.Group.DB.ExistsIsLeader(c, &db.ExistsIsLeaderParams{
		RelationID: relationID,
		AccountID:  fID,
	})
	if err != nil {
		global.Logger.Error(err.Error())
		return reply.TransferGroup{}, errcode.ErrServer
	}
	if !t {
		return reply.TransferGroup{}, myerr.NotLeader
	}
	t, err = dao.Group.DB.ExistsSetting(c, &db.ExistsSettingParams{
		AccountID:  tID,
		RelationID: relationID,
	})
	if err != nil {
		return reply.TransferGroup{}, errcode.ErrServer
	}
	if !t {
		return reply.TransferGroup{}, myerr.NotGroupMember
	}
	err = dao.Group.DB.TransferIsSelfFalse(c, &db.TransferIsSelfFalseParams{
		RelationID: relationID,
		AccountID:  fID,
	})
	if err != nil {
		return reply.TransferGroup{}, errcode.ErrServer
	}
	err = dao.Group.DB.TransferIsSelfTrue(c, &db.TransferIsSelfTrueParams{
		RelationID: relationID,
		AccountID:  tID,
	})
	if err != nil {
		return reply.TransferGroup{}, errcode.ErrServer
	}
	return reply.TransferGroup{}, nil
}
func (mGroup) DissolveGroup(c *gin.Context, relationId int64, accountID int64) (result reply.DissolveGroup, mErr errcode.Err) {
	t, err := dao.Group.DB.ExistsIsLeader(c, &db.ExistsIsLeaderParams{
		RelationID: relationId,
		AccountID:  accountID,
	})
	if err != nil {
		return result, errcode.ErrServer
	}
	if !t {
		return result, myerr.NotLeader
	}
	err = dao.Group.DB.DissolveGroup(c, relationId)
	if err != nil {
		return result, errcode.ErrServer
	}
	err = dao.Group.Redis.DelGroupRelation(c, relationId)
	if err != nil {
		return result, errcode.ErrServer
	}
	return result, nil
}
func (mGroup) UpdateGroup(c *gin.Context, params request.UpdateGroup, accountID int64) (result reply.UpdateGroup, mErr errcode.Err) {
	t, err := dao.Group.DB.ExistsSetting(c, &db.ExistsSettingParams{
		AccountID:  accountID,
		RelationID: params.RelationID,
	})
	if err != nil {
		global.Logger.Error(err.Error())
		return result, errcode.ErrServer
	}
	if !t {
		return result, myerr.NotGroupMember
	}
	data, err := dao.Group.DB.GetGroupRelationByID(c, params.RelationID)
	if err != nil {
		global.Logger.Error(err.Error())
		return result, errcode.ErrServer
	}
	err = dao.Group.DB.UpdateGroupRelation(c, &db.UpdateGroupRelationParams{
		Name:        params.Name,
		Description: params.Description,
		ID:          params.RelationID,
		Avatar:      data.Avatar,
	})
	if err != nil {
		global.Logger.Error(err.Error())
		return result, errcode.ErrServer
	}
	result = reply.UpdateGroup{
		Name:        params.Name,
		Description: params.Description,
	}
	return result, nil
}

func (mGroup) InviteAccount(c *gin.Context, relationID int64, tID int64, fID int64) (result reply.InviteAccount, mErr errcode.Err) {
	t, err := dao.Group.DB.ExistsSetting(c, &db.ExistsSettingParams{
		AccountID:  fID,
		RelationID: relationID,
	})
	if err != nil {
		return result, errcode.ErrServer
	}
	if !t {
		return result, myerr.NotGroupMember
	}
	err = dao.Group.DB.CreateSetting(c, &db.CreateSettingParams{
		AccountID:  tID,
		RelationID: relationID,
		IsLeader:   false,
		IsSelf:     false,
	})
	if err != nil {
		global.Logger.Error(err.Error())
		return result, errcode.ErrServer
	}
	err = dao.Group.Redis.AddGroupRelationAccount(c, relationID, tID)
	if err != nil {
		global.Logger.Error(err.Error())
		return result, errcode.ErrServer
	}
	return result, nil
}
func (mGroup) QuitGroup(c *gin.Context, relationID int64, accountID int64) (result reply.QuitGroup, mErr errcode.Err) {
	t, err := dao.Group.DB.ExistsIsLeader(c, &db.ExistsIsLeaderParams{
		RelationID: relationID,
		AccountID:  accountID,
	})
	if err != nil {
		return result, errcode.ErrServer
	}
	if t {
		return result, myerr.IsLeader
	}
	err = dao.Group.DB.DeleteSetting(c, &db.DeleteSettingParams{
		AccountID:  accountID,
		RelationID: relationID,
	})
	if err != nil {
		return result, errcode.ErrServer
	}
	err = dao.Group.Redis.DelGroupRelationAccount(c, relationID, accountID)
	if err != nil {
		return result, errcode.ErrServer
	}
	return result, nil
}
