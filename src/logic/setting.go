package logic

import (
	"context"
	"strings"

	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/0RAJA/chat_app/src/dao"
	db "github.com/0RAJA/chat_app/src/dao/postgres/sqlc"
	"github.com/0RAJA/chat_app/src/global"
	mid "github.com/0RAJA/chat_app/src/middleware"
	"github.com/0RAJA/chat_app/src/model"
	"github.com/0RAJA/chat_app/src/model/chat/server"
	"github.com/0RAJA/chat_app/src/model/reply"
	"github.com/0RAJA/chat_app/src/myerr"
	"github.com/0RAJA/chat_app/src/task"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

type setting struct {
}

// ExistsSetting 是否存在account和relation关系的联系
// 参数: accountID, relationID
// 成功: 是否存在,nil
// 失败: 打印错误日志 errcode.ErrServer
func ExistsSetting(c context.Context, accountID, relationID int64) (bool, errcode.Err) {
	ok, err := dao.Group.DB.ExistsSetting(c, &db.ExistsSettingParams{AccountID: accountID, RelationID: relationID})
	if err != nil {
		global.Logger.Error(err.Error())
		return false, errcode.ErrServer
	}
	return ok, nil
}

func (setting) GetFriends(c *gin.Context, accountID int64) (reply.GetFriends, errcode.Err) {
	data, err := dao.Group.DB.GetFriendSettingsOrderByName(c, accountID)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return reply.GetFriends{}, errcode.ErrServer
	}
	result := reply.GetFriends{
		List:  make([]*model.SettingFriend, 0, len(data)),
		Total: int64(len(data)),
	}
	for _, v := range data {
		result.List = append(result.List, &model.SettingFriend{
			SettingInfo: model.SettingInfo{
				RelationID:   v.RelationID,
				RelationType: string(db.RelationtypeFriend),
				NickName:     v.NickName,
				IsNotDisturb: v.IsNotDisturb,
				IsPin:        v.IsPin,
				PinTime:      v.PinTime,
				IsShow:       v.IsShow,
				LastShow:     v.LastShow,
			},
			FriendInfo: &model.SettingFriendInfo{
				AccountID: v.AccountID,
				Name:      v.AccountName,
				Avatar:    v.AccountAvatar,
			},
		})
	}
	return result, nil
}

func (setting) GetPins(c *gin.Context, accountID int64) (reply.GetPins, errcode.Err) {
	friendData, err := dao.Group.DB.GetFriendPinSettingsOrderByPinTime(c, accountID)
	if err != nil {
		return reply.GetPins{List: []*model.SettingPin{}}, errcode.ErrServer
	}
	groupData, err := dao.Group.DB.GetGroupPinSettingsOrderByPinTime(c, accountID)
	if err != nil {
		return reply.GetPins{List: []*model.SettingPin{}}, errcode.ErrServer
	}
	result := make([]*model.SettingPin, 0, len(friendData)+len(groupData))
	for _, v := range friendData {
		friendInfo := &model.SettingFriendInfo{
			AccountID: v.AccountID,
			Name:      v.NickName,
			Avatar:    v.AccountAvatar,
		}
		result = append(result, &model.SettingPin{
			SettingPinInfo: model.SettingPinInfo{
				RelationID:   v.RelationID,
				RelationType: "friend",
				NickName:     v.NickName,
				PinTime:      v.PinTime,
			},
			GroupInfo:  nil,
			FriendInfo: friendInfo,
		})
	}
	for _, v := range groupData {
		groupType := strings.Split(v.GroupType.String, ",")
		groupInfo := &model.SettingGroupInfo{
			RelationID:  v.RelationID,
			Name:        groupType[0][1:],
			Description: groupType[1],
			Avatar:      groupType[2][:len(groupType[2])-1],
		}
		result = append(result, &model.SettingPin{
			SettingPinInfo: model.SettingPinInfo{
				RelationID:   v.RelationID,
				RelationType: "group",
				NickName:     v.NickName,
				PinTime:      v.PinTime,
			},
			GroupInfo:  groupInfo,
			FriendInfo: nil,
		})
	}
	return reply.GetPins{List: result}, nil
}

func (setting) GetShows(c *gin.Context, accountID int64) (reply.GetShows, errcode.Err) {
	friendData, err := dao.Group.DB.GetFriendShowSettingsOrderByShowTime(c, accountID)
	if err != nil {
		return reply.GetShows{List: []*model.Setting{}}, errcode.ErrServer
	}
	groupData, err := dao.Group.DB.GetGroupShowSettingsOrderByShowTime(c, accountID)
	if err != nil {
		return reply.GetShows{List: []*model.Setting{}}, errcode.ErrServer
	}
	result := make([]*model.Setting, 0, len(friendData)+len(groupData))
	for _, v := range friendData {
		friendInfo := &model.SettingFriendInfo{
			AccountID: v.AccountID,
			Name:      v.NickName,
			Avatar:    v.AccountAvatar,
		}
		result = append(result, &model.Setting{
			SettingInfo: model.SettingInfo{
				RelationID:   v.RelationID,
				RelationType: "friend",
				NickName:     v.NickName,
				IsNotDisturb: v.IsNotDisturb,
				IsPin:        v.IsPin,
				PinTime:      v.PinTime,
				IsShow:       v.IsShow,
				LastShow:     v.LastShow,
			},
			GroupInfo:  nil,
			FriendInfo: friendInfo,
		})
	}
	for _, v := range groupData {
		groupType := strings.Split(v.GroupType.String, ",")
		groupInfo := &model.SettingGroupInfo{
			RelationID:  v.RelationID,
			Name:        groupType[0][1:],
			Description: groupType[1],
			Avatar:      groupType[2][:len(groupType[2])-1],
		}
		result = append(result, &model.Setting{
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
			GroupInfo:  groupInfo,
			FriendInfo: nil,
		})
	}
	return reply.GetShows{List: result}, nil
}

func getFriendRelationByID(c *gin.Context, relationID int64) (*db.GetFriendRelationByIDRow, errcode.Err) {
	relationInfo, err := dao.Group.DB.GetFriendRelationByID(c, relationID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, myerr.RelationNotExists
		}
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	return relationInfo, nil
}

func (setting) DeleteFriend(c *gin.Context, accountID, relationID int64) errcode.Err {
	relationInfo, merr := getFriendRelationByID(c, relationID)
	if merr != nil {
		return merr
	}
	if accountID != relationInfo.Account1ID || accountID != relationInfo.Account2ID {
		return myerr.AuthPermissionsInsufficient
	}
	if err := dao.Group.DB.DeleteRelationWithTx(c, dao.Group.Redis, relationID); err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	accessToken, _ := mid.GetToken(c.Request.Header)
	// 推送删除通知
	global.Worker.SendTask(task.DeleteRelation(accessToken, accountID, relationID))
	return nil
}

func (setting) UpdateNickName(c *gin.Context, accountID, relationID int64, nickName string) errcode.Err {
	settingInfo, err := dao.Group.DB.GetSettingByID(c, &db.GetSettingByIDParams{
		AccountID:  accountID,
		RelationID: relationID,
	})
	switch err {
	case pgx.ErrNoRows:
		return myerr.RelationNotExists
	case nil:
		if settingInfo.NickName == nickName {
			return nil
		}
		if err := dao.Group.DB.UpdateSettingNickName(c, &db.UpdateSettingNickNameParams{
			AccountID:  accountID,
			RelationID: relationID,
			NickName:   nickName,
		}); err != nil {
			global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
			return errcode.ErrServer
		}
		accessToken, _ := mid.GetToken(c.Request.Header)
		// 推送更改昵称通知
		global.Worker.SendTask(task.UpdateNickName(accessToken, accountID, relationID, nickName))
		return nil
	default:
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
}

func (setting) UpdateSettingPin(c *gin.Context, accountID, relationID int64, isPin bool) errcode.Err {
	settingInfo, err := dao.Group.DB.GetSettingByID(c, &db.GetSettingByIDParams{
		AccountID:  accountID,
		RelationID: relationID,
	})
	switch err {
	case pgx.ErrNoRows:
		return myerr.RelationNotExists
	case nil:
		if settingInfo.IsPin == isPin {
			return nil
		}
		if err := dao.Group.DB.UpdateSettingPin(c, &db.UpdateSettingPinParams{
			AccountID:  accountID,
			RelationID: relationID,
			IsPin:      isPin,
		}); err != nil {
			global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
			return errcode.ErrServer
		}
		accessToken, _ := mid.GetToken(c.Request.Header)
		// 推送更改置顶通知
		global.Worker.SendTask(task.UpdateSettingState(accessToken, server.SettingPin, accountID, relationID, isPin))
		return nil
	default:
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
}

func (setting) UpdateSettingDisturb(c *gin.Context, accountID, relationID int64, isNotDisturb bool) errcode.Err {
	settingInfo, err := dao.Group.DB.GetSettingByID(c, &db.GetSettingByIDParams{
		AccountID:  accountID,
		RelationID: relationID,
	})
	switch err {
	case pgx.ErrNoRows:
		return myerr.RelationNotExists
	case nil:
		if settingInfo.IsNotDisturb == isNotDisturb {
			return nil
		}
		if err := dao.Group.DB.UpdateSettingDisturb(c, &db.UpdateSettingDisturbParams{
			AccountID:    accountID,
			RelationID:   relationID,
			IsNotDisturb: isNotDisturb,
		}); err != nil {
			global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
			return errcode.ErrServer
		}
		// 推送更改免打扰通知
		accessToken, _ := mid.GetToken(c.Request.Header)
		global.Worker.SendTask(task.UpdateSettingState(accessToken, server.SettingNotDisturb, accountID, relationID, isNotDisturb))
		return nil
	default:
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
}

func (setting) UpdateSettingShow(c *gin.Context, accountID, relationID int64, isShow bool) errcode.Err {
	settingInfo, err := dao.Group.DB.GetSettingByID(c, &db.GetSettingByIDParams{
		AccountID:  accountID,
		RelationID: relationID,
	})
	switch err {
	case pgx.ErrNoRows:
		return myerr.RelationNotExists
	case nil:
		if settingInfo.IsShow == isShow {
			return nil
		}
		if err := dao.Group.DB.UpdateSettingShow(c, &db.UpdateSettingShowParams{
			AccountID:  accountID,
			RelationID: relationID,
			IsShow:     isShow,
		}); err != nil {
			global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
			return errcode.ErrServer
		}
		// 推送更改是否展示通知
		accessToken, _ := mid.GetToken(c.Request.Header)
		global.Worker.SendTask(task.UpdateSettingState(accessToken, server.SettingShow, accountID, relationID, isShow))
		return nil
	default:
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
}

func (setting) GetFriendsByName(c *gin.Context, accountID int64, name string, limit, offset int32) (reply.GetFriendsByName, errcode.Err) {
	data, err := dao.Group.DB.GetFriendSettingsByName(c, &db.GetFriendSettingsByNameParams{
		AccountID: accountID,
		Limit:     limit,
		Offset:    offset,
		Name:      name,
	})
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return reply.GetFriendsByName{}, errcode.ErrServer
	}
	if len(data) == 0 {
		return reply.GetFriendsByName{List: []*model.SettingFriend{}}, nil
	}
	list := make([]*model.SettingFriend, 0, len(data))
	for _, v := range data {
		list = append(list, &model.SettingFriend{
			SettingInfo: model.SettingInfo{
				RelationID:   v.RelationID,
				RelationType: string(db.RelationtypeFriend),
				NickName:     v.NickName,
				IsNotDisturb: v.IsNotDisturb,
				IsPin:        v.IsPin,
				PinTime:      v.PinTime,
				IsShow:       v.IsShow,
				LastShow:     v.LastShow,
			},
			FriendInfo: &model.SettingFriendInfo{
				AccountID: v.AccountID,
				Name:      v.NickName,
				Avatar:    v.AccountAvatar,
			},
		})
	}
	return reply.GetFriendsByName{
		List:  list,
		Total: data[0].Total,
	}, nil
}
