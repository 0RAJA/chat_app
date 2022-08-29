package logic

import (
	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/0RAJA/chat_app/src/dao"
	db "github.com/0RAJA/chat_app/src/dao/postgres/sqlc"
	"github.com/0RAJA/chat_app/src/global"
	mid "github.com/0RAJA/chat_app/src/middleware"
	"github.com/0RAJA/chat_app/src/model"
	"github.com/0RAJA/chat_app/src/model/reply"
	"github.com/0RAJA/chat_app/src/model/request"
	"github.com/0RAJA/chat_app/src/myerr"
	"github.com/0RAJA/chat_app/src/task"
	"github.com/gin-gonic/gin"
)

type mGroup struct {
}

func (mGroup) CreateGroup(c *gin.Context, accountID int64, name string, desc string) (relationID int64, mErr errcode.Err) {
	err := dao.Group.DB.AddSettingWithTx(c, dao.Group.Redis, relationID, accountID, true)

	if err != nil {
		global.Logger.Error(err.Error())
		return 0, errcode.ErrServer
	}
	return relationID, nil
}

func (mGroup) TransferGroup(c *gin.Context, relationID int64, fID int64, tID int64) (result reply.TransferGroup, mErr errcode.Err) {
	t, err := dao.Group.DB.ExistsIsLeader(c, &db.ExistsIsLeaderParams{
		RelationID: relationID,
		AccountID:  fID,
	})
	if err != nil {
		global.Logger.Error(err.Error())
		return result, errcode.ErrServer
	}
	if !t {
		return result, myerr.NotLeader
	}
	t, err = dao.Group.DB.ExistsSetting(c, &db.ExistsSettingParams{
		AccountID:  tID,
		RelationID: relationID,
	})
	if err != nil {
		return result, errcode.ErrServer
	}
	if !t {
		return result, myerr.NotGroupMember
	}
	err = dao.Group.DB.TransferGroup(c, relationID, fID, tID)
	if err != nil {
		global.Logger.Error(err.Error())
		return result, errcode.ErrServer
	}
	accessToken, _ := mid.GetToken(c.Request.Header)
	global.Worker.SendTask(task.TransferGroup(accessToken, tID, relationID))
	return reply.TransferGroup{}, nil
}
func (mGroup) DissolveGroup(c *gin.Context, relationId int64, accountID int64) (result reply.DissolveGroup, mErr errcode.Err) {
	t, err := dao.Group.DB.ExistsIsLeader(c, &db.ExistsIsLeaderParams{
		RelationID: relationId,
		AccountID:  accountID,
	})
	if err != nil {
		global.Logger.Error(err.Error())
		return result, errcode.ErrServer
	}
	if !t {
		return result, myerr.NotLeader
	}
	err = dao.Group.DB.DeleteRelationWithTx(c, dao.Group.Redis, relationId)
	if err != nil {
		global.Logger.Error(err.Error())
		return result, errcode.ErrServer
	}
	accessToken, _ := mid.GetToken(c.Request.Header)
	global.Worker.SendTask(task.DissolveGroup(accessToken, relationId))
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

func (mGroup) InviteAccount(c *gin.Context, relationID int64, tID []int64, fID int64) (result reply.InviteAccount, mErr errcode.Err) {
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
	result.InviteMember = make([]int64, 0, len(tID))
	for _, v := range tID {
		f1, e1 := dao.Group.DB.ExistsFriendSetting(c, &db.ExistsFriendSettingParams{
			Account1ID: v,
			Account2ID: fID,
		})
		if e1 != nil {
			continue
		}
		f2, e2 := dao.Group.DB.ExistsSetting(c, &db.ExistsSettingParams{
			AccountID:  v,
			RelationID: relationID,
		})
		if e2 != nil {
			continue
		}
		if f1 && !f2 {
			err = dao.Group.DB.AddSettingWithTx(c, dao.Group.Redis, relationID, v, false)
			if err != nil {
				global.Logger.Error(err.Error())
				continue
			}
			result.InviteMember = append(result.InviteMember, v)
		}

	}

	accessToken, _ := mid.GetToken(c.Request.Header)
	for _, v := range tID {
		global.Worker.SendTask(task.InviteAccount(accessToken, relationID, v))
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

	err = dao.Group.DB.DeleteSettingWithTx(c, dao.Group.Redis, relationID, accountID)

	if err != nil {
		global.Logger.Error(err.Error())
		return result, errcode.ErrServer
	}
	accessToken, _ := mid.GetToken(c.Request.Header)
	global.Worker.SendTask(task.QuitGroup(accessToken, relationID, accountID))
	return result, nil
}

func (mGroup) GroupList(c *gin.Context, accountID int64) (reply.GetGroup, errcode.Err) {
	data, err := dao.Group.DB.GetGroupList(c, accountID)
	if err != nil {
		return reply.GetGroup{}, errcode.ErrServer
	}
	groupList := make([]model.SettingGroup, 0, len(data))
	for _, v := range data {
		t := model.SettingGroup{
			SettingInfo: model.SettingInfo{
				RelationID:   v.RelationID,
				RelationType: "group",
				NickName:     v.NickName,
				IsNotDisturb: v.IsNotDisturb,
				IsPin:        v.IsPin,
				PinTime:      v.PinTime,
				IsShow:       v.IsShow,
				LastShow:     v.LastShow,
			},
			GroupInfo: &model.SettingGroupInfo{
				RelationID:  v.RelationID,
				Name:        v.GroupName.(string),
				Description: v.Description.(string),
				Avatar:      v.GroupAvatar.(string),
			},
		}
		groupList = append(groupList, t)
	}
	return reply.GetGroup{
		List:  groupList,
		Total: int64(len(data)),
	}, nil
}
func (mGroup) GetGroupByName(c *gin.Context, accountID int64, limit int32, offset int32, name string) (reply.GetGroup, errcode.Err) {
	data, err := dao.Group.DB.GetGroupSettingsByName(c, &db.GetGroupSettingsByNameParams{
		AccountID: accountID,
		Limit:     limit,
		Offset:    offset,
		Name:      name,
	})
	if err != nil {
		return reply.GetGroup{}, errcode.ErrServer
	}
	groupList := make([]model.SettingGroup, 0, data[0].Total)
	for _, v := range data {
		t := model.SettingGroup{
			SettingInfo: model.SettingInfo{
				RelationID:   v.RelationID,
				RelationType: "group",
				NickName:     v.NickName,
				IsNotDisturb: v.IsNotDisturb,
				IsPin:        v.IsPin,
				PinTime:      v.PinTime,
				IsShow:       v.IsShow,
				LastShow:     v.LastShow,
			},
			GroupInfo: &model.SettingGroupInfo{
				RelationID:  v.RelationID,
				Name:        v.GroupName.(string),
				Description: v.Description.(string),
				Avatar:      v.GroupAvatar.(string),
			},
		}
		groupList = append(groupList, t)
	}
	return reply.GetGroup{
		List:  groupList,
		Total: data[0].Total,
	}, nil
}
func (mGroup) GeGroupMembers(c *gin.Context, accountID int64, relationID int64) ([]reply.GetGroupMembers, errcode.Err) {
	t, err := dao.Group.DB.ExistsSetting(c, &db.ExistsSettingParams{
		AccountID:  accountID,
		RelationID: relationID,
	})
	if err != nil {
		return nil, errcode.ErrServer
	}
	if !t {
		return nil, myerr.NotGroupMember
	}
	data, err := dao.Group.DB.GetGroupMembersByID(c, relationID)
	if err != nil {
		global.Logger.Error(err.Error())
	}
	result := make([]reply.GetGroupMembers, 0, len(data))
	for _, v := range data {
		t := reply.GetGroupMembers{
			ID:       v.ID,
			Name:     v.Name,
			Avatar:   v.Avatar,
			NickName: v.NickName.String,
			IsLeader: v.IsLeader.Bool,
		}
		result = append(result, t)
	}
	return result, nil
}
