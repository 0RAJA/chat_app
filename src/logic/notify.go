package logic

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/0RAJA/chat_app/src/model"
	"github.com/0RAJA/chat_app/src/task"

	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/0RAJA/chat_app/src/dao"
	db "github.com/0RAJA/chat_app/src/dao/postgres/sqlc"
	"github.com/0RAJA/chat_app/src/global"
	mid "github.com/0RAJA/chat_app/src/middleware"
	"github.com/0RAJA/chat_app/src/model/reply"
	"github.com/0RAJA/chat_app/src/model/request"
	"github.com/0RAJA/chat_app/src/myerr"
	"github.com/gin-gonic/gin"
)

type notify struct {
}

func (notify) CreateNotify(c *gin.Context, params *request.CreateNotify, accountID int64) (reply.GroupNotify, errcode.Err) {
	result := reply.GroupNotify{}
	t, err := dao.Group.DB.ExistsIsLeader(c, &db.ExistsIsLeaderParams{
		AccountID:  accountID,
		RelationID: params.RelationID,
	})
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return result, errcode.ErrServer
	}
	if !t {
		return result, myerr.NotLeader
	}
	expand, _ := model.ExpandToJson(params.MsgExtend)
	r, err := dao.Group.DB.CreateGroupNotify(c, &db.CreateGroupNotifyParams{
		RelationID: sql.NullInt64{Int64: params.RelationID, Valid: true},
		MsgContent: params.MsgContent,
		MsgExpand:  expand,
		AccountID:  sql.NullInt64{Int64: accountID, Valid: true},
		ReadIds:    []int64{accountID},
		CreateAt:   time.Now(),
	})
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return result, errcode.ErrServer
	}
	msgExpand, err := model.JsonToExpand(r.MsgExpand)
	if err != nil {
		global.Logger.Error(err.Error())
		return result, errcode.ErrServer
	}
	result = reply.GroupNotify{
		ID:         r.ID,
		RelationID: r.RelationID.Int64,
		MsgContent: r.MsgContent,
		MsgExpand:  msgExpand,
		AccountID:  r.AccountID.Int64,
		CreateAt:   r.CreateAt,
		ReadIds:    nil,
	}
	accessToken, _ := mid.GetToken(c.Request.Header)
	global.Worker.SendTask(task.CreateNotify(accessToken, accountID, params.RelationID, r.MsgContent, msgExpand))
	return result, nil
}

func (notify) UpdateNotify(c *gin.Context, params *request.UpdateNotify, accountID int64) (result reply.UpdateNotify, mErr errcode.Err) {

	t, err := dao.Group.DB.ExistsSetting(c, &db.ExistsSettingParams{
		AccountID:  accountID,
		RelationID: params.RelationID,
	})
	if err != nil {
		return result, errcode.ErrServer
	}
	if !t {
		return result, myerr.NotGroupMember
	}
	expand, _ := model.ExpandToJson(params.MsgExtend)
	_, err = dao.Group.DB.UpdateGroupNotify(c, &db.UpdateGroupNotifyParams{
		RelationID: sql.NullInt64{Int64: params.RelationID, Valid: true},
		MsgContent: params.MsgContent,
		MsgExpand:  expand,
		AccountID:  sql.NullInt64{Int64: accountID, Valid: true},
		ReadIds:    []int64{accountID},
		CreateAt:   time.Now(),
		ID:         params.ID,
	})
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return result, errcode.ErrServer
	}
	if err != nil {
		global.Logger.Error(err.Error())
		return result, errcode.ErrServer
	}
	accessToken, _ := mid.GetToken(c.Request.Header)
	global.Worker.SendTask(task.UpdateNotify(accessToken, accountID, params.RelationID, params.MsgContent, params.MsgExtend))
	return result, nil
}
func (notify) GetNotifyByID(c *gin.Context, relationID int64, accountId int64) (result reply.GetNotify, mErr errcode.Err) {
	t, err := dao.Group.DB.ExistsSetting(c, &db.ExistsSettingParams{
		AccountID:  accountId,
		RelationID: relationID,
	})
	fmt.Println(relationID, accountId)
	if err != nil {
		return result, errcode.ErrServer
	}
	if !t {
		return result, myerr.NotGroupMember
	}
	r, err := dao.Group.DB.GetGroupNotifyByID(c, sql.NullInt64{
		Int64: relationID,
		Valid: true,
	})
	if err != nil {
		if err != sql.ErrNoRows {
			return result, myerr.NotifyNotExist
		}
		return result, errcode.ErrServer
	}

	for _, v := range r {
		msgExpand, err := model.JsonToExpand(v.MsgExpand)
		if err != nil {
			global.Logger.Error(err.Error())
			return reply.GetNotify{}, errcode.ErrServer
		}
		re := reply.Notify{
			ID:         v.ID,
			RelationID: v.RelationID.Int64,
			MsgContent: v.MsgContent,
			MsgExpand:  msgExpand,
			AccountID:  v.AccountID.Int64,
			CreateAt:   v.CreateAt,
			ReadIds:    v.ReadIds,
		}
		result.List = append(result.List, re)
	}
	result.Total = int64(len(r))
	return result, nil
}

func (notify) DeleteNotify(c *gin.Context, id, relationID, accountID int64) errcode.Err {
	t, err := dao.Group.DB.ExistsIsLeader(c, &db.ExistsIsLeaderParams{
		AccountID:  accountID,
		RelationID: relationID,
	})
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	if !t {
		return myerr.NotLeader
	}
	err = dao.Group.DB.DeleteGroupNotify(c, id)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	return nil
}
