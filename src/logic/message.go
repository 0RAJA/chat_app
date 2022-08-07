package logic

import (
	"database/sql"
	"errors"

	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/0RAJA/chat_app/src/dao"
	db "github.com/0RAJA/chat_app/src/dao/postgres/sqlc"
	"github.com/0RAJA/chat_app/src/global"
	mid "github.com/0RAJA/chat_app/src/middleware"
	"github.com/0RAJA/chat_app/src/model"
	"github.com/0RAJA/chat_app/src/model/reply"
	"github.com/0RAJA/chat_app/src/myerr"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

type message struct {
}

// ExistsSetting 是否存在关系
// 参数: accountID, relationID
// 成功: 是否存在,nil
// 失败: 打印错误日志 errcode.ErrServer
func ExistsSetting(c *gin.Context, accountID, relationID int64) (bool, errcode.Err) {
	ok, err := dao.Group.DB.ExistsSetting(c, &db.ExistsSettingParams{AccountID: accountID, RelationID: relationID})
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return false, errcode.ErrServer
	}
	return ok, nil
}

// GetMsgInfoByID 获取消息详情
// 参数: msgID 消息ID
// 成功: 消息详情,nil
// 失败: 打印错误日志 errcode.ErrServer,myerr.MsgNotExists
func GetMsgInfoByID(c *gin.Context, msgID int64) (*db.GetMsgByIDRow, errcode.Err) {
	result, err := dao.Group.DB.GetMsgByID(c, msgID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, myerr.MsgNotExists
		}
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	return result, nil
}

func (message) GetMsgsByRelationIDAndTime(c *gin.Context, params model.GetMsgsByRelationIDAndTimeParams) (reply.GetMsgsByRelationIDAndTime, errcode.Err) {
	ok, merr := ExistsSetting(c, params.AccountID, params.RelationID)
	if merr != nil {
		return reply.GetMsgsByRelationIDAndTime{}, merr
	}
	if !ok {
		return reply.GetMsgsByRelationIDAndTime{}, myerr.AuthPermissionsInsufficient
	}
	data, err := dao.Group.DB.GetMsgsByRelationIDAndTime(c, &db.GetMsgsByRelationIDAndTimeParams{
		RelationID: params.RelationID,
		CreateAt:   params.LastTime,
		Limit:      params.Limit,
		Offset:     params.Offset,
	})
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return reply.GetMsgsByRelationIDAndTime{}, errcode.ErrServer
	}
	if len(data) == 0 {
		return reply.GetMsgsByRelationIDAndTime{List: []*reply.MsgInfoWithRly{}}, nil
	}
	result := make([]*reply.MsgInfoWithRly, 0, len(data))
	for _, v := range data {
		var content string
		var extend *model.MsgExtend
		if !v.IsRevoke {
			content = v.MsgContent
			extend, err = model.JsonToExpand(v.MsgExtend)
			if err != nil {
				global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
				return reply.GetMsgsByRelationIDAndTime{}, errcode.ErrServer
			}
		}
		var readIDs []int64
		if params.AccountID == v.AccountID.Int64 {
			readIDs = v.ReadIds
		}
		var rlyMsg *reply.RlyMsg
		if v.RlyMsgID.Valid {
			rlyMsgInfo, merr := GetMsgInfoByID(c, v.RlyMsgID.Int64)
			if merr != nil {
				continue
			}
			var rlyContent string
			var rlyExtend *model.MsgExtend
			if !rlyMsgInfo.IsRevoke {
				rlyContent = rlyMsgInfo.MsgContent
				rlyExtend, err = model.JsonToExpand(rlyMsgInfo.MsgExtend)
				if err != nil {
					global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
					return reply.GetMsgsByRelationIDAndTime{}, errcode.ErrServer
				}
			}
			rlyMsg = &reply.RlyMsg{
				MsgID:      v.RlyMsgID.Int64,
				MsgType:    rlyMsgInfo.MsgType,
				MsgContent: rlyContent,
				MsgExtend:  rlyExtend,
				IsRevoke:   rlyMsgInfo.IsRevoke,
			}
		}
		result = append(result, &reply.MsgInfoWithRly{
			MsgInfo: reply.MsgInfo{
				ID:         v.ID,
				NotifyType: string(v.NotifyType),
				MsgType:    v.MsgType,
				MsgContent: content,
				Extend:     extend,
				FileID:     v.FileID.Int64,
				AccountID:  v.AccountID.Int64,
				RelationID: v.RelationID,
				CreateAt:   v.CreateAt,
				IsRevoke:   v.IsRevoke,
				IsTop:      v.IsTop,
				IsPin:      v.IsPin,
				PinTime:    v.PinTime,
				ReadIds:    readIDs,
				ReplyCount: v.ReplyCount,
			},
			RlyMsg: rlyMsg,
		})
	}
	return reply.GetMsgsByRelationIDAndTime{List: result, Total: data[0].Total}, nil
}

func (message) GetPinMsgsByRelationID(c *gin.Context, params model.GetPinMsgsByRelationIDParams) (reply.GetPinMsgsByRelationID, errcode.Err) {
	ok, merr := ExistsSetting(c, params.AccountID, params.RelationID)
	if merr != nil {
		return reply.GetPinMsgsByRelationID{}, merr
	}
	if !ok {
		return reply.GetPinMsgsByRelationID{}, myerr.AuthPermissionsInsufficient
	}
	data, err := dao.Group.DB.GetPinMsgsByRelationID(c, &db.GetPinMsgsByRelationIDParams{
		RelationID: params.RelationID,
		Limit:      params.Limit,
		Offset:     params.Offset,
	})
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return reply.GetPinMsgsByRelationID{}, errcode.ErrServer
	}
	if len(data) == 0 {
		return reply.GetPinMsgsByRelationID{List: []*reply.MsgInfo{}}, nil
	}
	result := make([]*reply.MsgInfo, 0, len(data))
	for _, v := range data {
		var content string
		var extend *model.MsgExtend
		if !v.IsRevoke {
			content = v.MsgContent
			extend, err = model.JsonToExpand(v.MsgExtend)
			if err != nil {
				global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
				return reply.GetPinMsgsByRelationID{}, errcode.ErrServer
			}
		}
		var readIDs []int64
		if params.AccountID == v.AccountID.Int64 {
			readIDs = v.ReadIds
		}
		result = append(result, &reply.MsgInfo{
			ID:         v.ID,
			NotifyType: string(v.NotifyType),
			MsgType:    v.MsgType,
			MsgContent: content,
			Extend:     extend,
			FileID:     v.FileID.Int64,
			AccountID:  v.AccountID.Int64,
			RelationID: v.RelationID,
			CreateAt:   v.CreateAt,
			IsRevoke:   v.IsRevoke,
			IsTop:      v.IsTop,
			IsPin:      v.IsPin,
			PinTime:    v.PinTime,
			ReadIds:    readIDs,
			ReplyCount: v.ReplyCount,
		})
	}
	return reply.GetPinMsgsByRelationID{List: result, Total: data[0].Total}, nil
}

func (message) GetRlyMsgsInfoByMsgID(c *gin.Context, params model.GetRlyMsgsInfoByMsgIDParams) (reply.GetRlyMsgsInfoByMsgID, errcode.Err) {
	ok, merr := ExistsSetting(c, params.AccountID, params.RelationID)
	if merr != nil {
		return reply.GetRlyMsgsInfoByMsgID{}, merr
	}
	if !ok {
		return reply.GetRlyMsgsInfoByMsgID{}, myerr.AuthPermissionsInsufficient
	}
	data, err := dao.Group.DB.GetRlyMsgsInfoByMsgID(c, &db.GetRlyMsgsInfoByMsgIDParams{
		RelationID: params.RelationID,
		RlyMsgID:   params.RlyMsgID,
		Limit:      params.Limit,
		Offset:     params.Offset,
	})
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return reply.GetRlyMsgsInfoByMsgID{}, errcode.ErrServer
	}
	if len(data) == 0 {
		return reply.GetRlyMsgsInfoByMsgID{List: []*reply.MsgInfo{}}, nil
	}
	result := make([]*reply.MsgInfo, 0, len(data))
	for _, v := range data {
		var content string
		var extend *model.MsgExtend
		if !v.IsRevoke {
			content = v.MsgContent
			extend, err = model.JsonToExpand(v.MsgExtend)
			if err != nil {
				global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
				return reply.GetRlyMsgsInfoByMsgID{}, errcode.ErrServer
			}
		}
		var readIDs []int64
		if params.AccountID == v.AccountID.Int64 {
			readIDs = v.ReadIds
		}
		result = append(result, &reply.MsgInfo{
			ID:         v.ID,
			NotifyType: string(v.NotifyType),
			MsgType:    v.MsgType,
			MsgContent: content,
			Extend:     extend,
			FileID:     v.FileID.Int64,
			AccountID:  v.AccountID.Int64,
			RelationID: v.RelationID,
			CreateAt:   v.CreateAt,
			IsRevoke:   v.IsRevoke,
			IsTop:      v.IsTop,
			IsPin:      v.IsPin,
			PinTime:    v.PinTime,
			ReadIds:    readIDs,
			ReplyCount: v.ReplyCount,
		})
	}
	return reply.GetRlyMsgsInfoByMsgID{List: result, Total: data[0].Total}, nil
}

func (message) GetTopMsgByRelationID(c *gin.Context, params model.GetTopMsgByRelationIDParams) (*reply.GetTopMsgByRelationID, errcode.Err) {
	ok, merr := ExistsSetting(c, params.AccountID, params.RelationID)
	if merr != nil {
		return nil, merr
	}
	if !ok {
		return nil, myerr.AuthPermissionsInsufficient
	}
	data, err := dao.Group.DB.GetTopMsgByRelationID(c, params.RelationID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	var content string
	var extend *model.MsgExtend
	if !data.IsRevoke {
		content = data.MsgContent
		extend, err = model.JsonToExpand(data.MsgExtend)
		if err != nil {
			global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
			return nil, errcode.ErrServer
		}
	}
	var readIDs []int64
	if params.AccountID == data.AccountID.Int64 {
		readIDs = data.ReadIds
	}
	return &reply.GetTopMsgByRelationID{
		MsgInfo: reply.MsgInfo{
			ID:         data.ID,
			NotifyType: string(data.NotifyType),
			MsgType:    string(data.MsgType),
			MsgContent: content,
			Extend:     extend,
			FileID:     data.FileID.Int64,
			AccountID:  data.AccountID.Int64,
			RelationID: data.RelationID,
			CreateAt:   data.CreateAt,
			IsRevoke:   data.IsRevoke,
			IsTop:      data.IsTop,
			IsPin:      data.IsPin,
			PinTime:    data.PinTime,
			ReadIds:    readIDs,
			ReplyCount: data.ReplyCount,
		},
	}, nil
}

func (message) UpdateMsgPin(c *gin.Context, params model.UpdateMsgPinParams) errcode.Err {
	ok, merr := ExistsSetting(c, params.AccountID, params.RelationID)
	if merr != nil {
		return merr
	}
	if !ok {
		return myerr.AuthPermissionsInsufficient
	}
	msgInfo, merr := GetMsgInfoByID(c, params.MsgID)
	if merr != nil {
		return merr
	}
	if msgInfo.IsPin == params.IsPin {
		return nil
	}
	err := dao.Group.DB.UpdateMsgPin(c, &db.UpdateMsgPinParams{
		ID:    params.MsgID,
		IsPin: params.IsPin,
	})
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	// TODO:推送pin通知
	return nil
}

func (message) UpdateMsgTop(c *gin.Context, params model.UpdateMsgTopParams) errcode.Err {
	ok, merr := ExistsSetting(c, params.AccountID, params.RelationID)
	if merr != nil {
		return merr
	}
	if !ok {
		return myerr.AuthPermissionsInsufficient
	}
	msgInfo, merr := GetMsgInfoByID(c, params.MsgID)
	if merr != nil {
		return merr
	}
	if msgInfo.IsTop == params.IsTop {
		return nil
	}
	var err error
	if params.IsTop {
		err = dao.Group.DB.UpdateMsgTopTrueByMsgIDWithTx(c, params.RelationID, params.MsgID)
	} else {
		err = dao.Group.DB.UpdateMsgTopFalseByMsgID(c, params.MsgID)
	}
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	// TODO:推送置顶通知
	return nil
}

func (message) RevokeMsg(c *gin.Context, params model.RevokeMsgParams) errcode.Err {
	msgInfo, merr := GetMsgInfoByID(c, params.MsgID)
	if merr != nil {
		return merr
	}
	// 检查权限
	if params.AccountID != msgInfo.AccountID.Int64 {
		return myerr.AuthPermissionsInsufficient
	}
	if msgInfo.IsRevoke {
		return myerr.MsgAlreadyRevoke
	}
	err := dao.Group.DB.RevokeMsgWithTx(c, params.MsgID, msgInfo.IsTop, msgInfo.IsPin)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	if msgInfo.IsTop {
		// TODO: 推送取消置顶通知
	}
	return nil
}

func (message) CreateFileMsg(c *gin.Context, params model.CreateFileMsgParams) (*reply.CreateFileMsg, errcode.Err) {
	// 检查权限
	ok, merr := ExistsSetting(c, params.AccountID, params.RelationID)
	if merr != nil {
		return nil, merr
	}
	if !ok {
		return nil, myerr.AuthPermissionsInsufficient
	}
	// 上传文件
	fileInfo, merr := PublishFile(c, model.PublishFile{
		File:       params.File,
		RelationID: params.RelationID,
		AccountID:  params.AccountID,
	})
	if merr != nil {
		return nil, merr
	}
	var isRly bool
	var rlyID int64
	if params.RlyMsgID > 0 {
		rlyInfo, merr := GetMsgInfoByID(c, params.RlyMsgID)
		if merr != nil {
			return nil, merr
		}
		if rlyInfo.IsRevoke {
			return nil, myerr.RlyMsgHasRevoked
		}
		isRly = true
		rlyID = params.RlyMsgID
	}
	extend, _ := model.ExpandToJson(nil)
	result, err := dao.Group.DB.CreateMsg(c, &db.CreateMsgParams{
		NotifyType: db.MsgnotifytypeCommon,
		MsgType:    string(model.MsgTypeFile),
		MsgContent: fileInfo.Url,
		MsgExtend:  extend,
		FileID:     sql.NullInt64{Int64: fileInfo.ID, Valid: true},
		AccountID:  sql.NullInt64{Int64: params.AccountID, Valid: true},
		RlyMsgID:   sql.NullInt64{Int64: rlyID, Valid: isRly},
		RelationID: 0,
	})
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	return &reply.CreateFileMsg{
		ID:         result.ID,
		MsgContent: result.MsgContent,
		FileID:     result.FileID.Int64,
		CreateAt:   result.CreateAt,
	}, nil
}
