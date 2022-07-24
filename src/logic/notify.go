package logic

import (
	"database/sql"
	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/0RAJA/chat_app/src/dao"
	db "github.com/0RAJA/chat_app/src/dao/postgres/sqlc"
	"github.com/0RAJA/chat_app/src/global"
	mid "github.com/0RAJA/chat_app/src/middleware"
	"github.com/0RAJA/chat_app/src/model/reply"
	"github.com/0RAJA/chat_app/src/model/request"
	"github.com/0RAJA/chat_app/src/myerr"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgtype"
)

type notify struct {

}

func (notify)CreateNotify(c *gin.Context,params *request.CreateNotify) (reply.GroupNotify,errcode.Err)  {
	r,err := dao.Group.DB.CreateGroupNotify(c,&db.CreateGroupNotifyParams{
		RelationID: sql.NullInt64{Int64: params.RelationID, Valid: true},
		MsgContent: params.MsgContent,
		MsgExpand:  pgtype.JSON{
			Bytes:  []byte(params.MsgContent),
			Status: 0,
		},
		AccountID:  sql.NullInt64{Int64: params.AccountID, Valid: true},
		ReadIds:    nil,
	})
	result := reply.GroupNotify{}
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return result,errcode.ErrServer
	}
	result = reply.GroupNotify{
		ID:         r.ID,
		RelationID: r.RelationID.Int64,
		MsgContent: r.MsgContent,
		MsgExpand:  r.MsgExpand,
		AccountID:  r.AccountID.Int64,
		CreateAt:   r.CreateAt,
		ReadIds:    nil,
	}

	return result,nil
}

func (notify)UpdateNotify(c *gin.Context,params *request.UpdateNotify) (result reply.UpdateNotify,mErr errcode.Err) {
	_,err := dao.Group.DB.UpdateGroupNotify(c,&db.UpdateGroupNotifyParams{
		RelationID: sql.NullInt64{Int64: params.RelationID, Valid: true},
		MsgContent: params.MsgContent,
		MsgExpand:  pgtype.JSON{Bytes: []byte(params.MsgExpand)},
		AccountID:  sql.NullInt64{Int64: params.AccountID,Valid: true},
		CreateAt:   params.UpdateAt,
		ReadIds:    nil,
	})
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return result,errcode.ErrServer
	}
	return result,nil
}
func (notify)GetNotifyByID(c *gin.Context,relationID int64) (result reply.GetNotify,mErr errcode.Err) {
	r,err := dao.Group.DB.GetGroupNotifyByID(c,sql.NullInt64{
		Int64: relationID,
		Valid: false,
	})
	if err != nil {
		if err != sql.ErrNoRows {
			return result, myerr.NotifyNotExist
		}
		return result, errcode.ErrServer
	}
	result = reply.GetNotify{
		ID:         r.ID,
		RelationID: r.RelationID.Int64,
		MsgContent: r.MsgContent,
		MsgExpand:  r.MsgExpand,
		AccountID:  r.AccountID.Int64,
		CreateAt:   r.CreateAt,
		ReadIds:    nil,
	}
	return result,nil
}