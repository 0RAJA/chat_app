package logic

import (
	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/0RAJA/chat_app/src/dao"
	db "github.com/0RAJA/chat_app/src/dao/postgres/sqlc"
	"github.com/0RAJA/chat_app/src/dao/postgres/tx"
	"github.com/0RAJA/chat_app/src/global"
	mid "github.com/0RAJA/chat_app/src/middleware"
	"github.com/0RAJA/chat_app/src/model/reply"
	"github.com/0RAJA/chat_app/src/myerr"
	"github.com/0RAJA/chat_app/src/task"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

type application struct {
}

// 通过AccountID获取申请信息，返回申请信息
// 成功: 申请信息
// 错误: 打印日志, myerr.ApplicationNotExists,errcode.ErrServer
func getApplication(c *gin.Context, account1ID, account2ID int64) (*db.Application, errcode.Err) {
	aply, err := dao.Group.DB.GetApplicationByID(c, &db.GetApplicationByIDParams{Account1ID: account1ID, Account2ID: account2ID})
	switch err {
	case nil:
		return aply, nil
	case pgx.ErrNoRows:
		return nil, myerr.ApplicationNotExists
	default:
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
}

func (application) Create(c *gin.Context, account1ID, account2ID int64, applyMsg string) errcode.Err {
	// 不能自己对自己发送申请
	if account1ID == account2ID {
		return myerr.ApplicationNotValid
	}
	// 判断是否已经存在好友关系
	id1, id2 := sortID(account1ID, account2ID)
	exist, err := dao.Group.DB.ExistsFriendRelation(c, &db.ExistsFriendRelationParams{Account1ID: id1, Account2ID: id2})
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	if exist {
		return myerr.RelationExists
	}
	// 创建申请
	err = dao.Group.DB.CreateApplicationTx(c, &db.CreateApplicationParams{Account1ID: account1ID, Account2ID: account2ID, ApplyMsg: applyMsg})
	switch err {
	case nil:
		// 提示对方有新的申请信息
		global.Worker.SendTask(task.Application(account2ID))
		return nil
	case tx.ErrApplicationExists:
		return myerr.ApplicationExists
	default:
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
}

func (application) Delete(c *gin.Context, account1ID, account2ID int64) errcode.Err {
	aply, merr := getApplication(c, account1ID, account2ID)
	if merr != nil {
		return merr
	}
	if aply.Account1ID != account1ID {
		return myerr.AuthPermissionsInsufficient
	}
	if err := dao.Group.DB.DeleteApplication(c, &db.DeleteApplicationParams{Account1ID: account1ID, Account2ID: account2ID}); err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	return nil
}

// Accept account1ID是被申请者，account2ID是申请者
func (application) Accept(c *gin.Context, account1ID, account2ID int64) errcode.Err {
	aply, merr := getApplication(c, account2ID, account1ID)
	if merr != nil {
		return merr
	}
	if aply.Status == db.ApplicationstatusValue1 {
		return myerr.ApplicationRepeatOpt
	}
	account1Info, merr := getAccountInfoByID(c, account1ID, account1ID)
	if merr != nil {
		return merr
	}
	account2Info, merr := getAccountInfoByID(c, account2ID, account1ID)
	if merr != nil {
		return merr
	}
	msgInfo, err := dao.Group.DB.AcceptApplicationTx(c, dao.Group.Redis, account1Info, account2Info)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	// 推送信息
	global.Worker.SendTask(task.PublishMsg("", reply.MsgInfo{
		ID:         msgInfo.ID,
		NotifyType: string(msgInfo.NotifyType),
		MsgType:    msgInfo.MsgType,
		MsgContent: msgInfo.MsgContent,
		RelationID: msgInfo.RelationID,
		CreateAt:   msgInfo.CreateAt,
	}, nil))
	return nil
}

func (application) Refuse(c *gin.Context, account1ID, account2ID int64, refuseMsg string) errcode.Err {
	aply, merr := getApplication(c, account2ID, account1ID)
	if merr != nil {
		return merr
	}
	if aply.Status == db.ApplicationstatusValue2 {
		return myerr.ApplicationRepeatOpt
	}
	if err := dao.Group.DB.UpdateApplication(c, &db.UpdateApplicationParams{
		Account1ID: account1ID,
		Account2ID: account2ID,
		Status:     db.ApplicationstatusValue2,
		RefuseMsg:  refuseMsg,
	}); err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	return nil
}

func (application) List(c *gin.Context, accountID int64, limit, offset int32) (reply.ListApplications, errcode.Err) {
	aplys, err := dao.Group.DB.GetApplications(c, &db.GetApplicationsParams{AccountID: accountID, Offset: offset, Limit: limit})
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return reply.ListApplications{}, errcode.ErrServer
	}
	if len(aplys) == 0 {
		return reply.ListApplications{List: []*reply.ApplicationInfo{}}, nil
	}
	result := make([]*reply.ApplicationInfo, 0, len(aplys))
	for _, aply := range aplys {
		name, avatar := aply.Account1Name, aply.Account1Avatar
		if aply.Account1ID == accountID {
			name, avatar = aply.Account2Name, aply.Account2Avatar
		}
		result = append(result, &reply.ApplicationInfo{
			Account1ID: aply.Account1ID,
			Account2ID: aply.Account2ID,
			ApplyMsg:   aply.ApplyMsg,
			RefuseMsg:  aply.RefuseMsg,
			Status:     string(aply.Status),
			CreateAt:   aply.CreateAt,
			UpdateAt:   aply.UpdateAt,
			Name:       name,
			Avatar:     avatar,
		})
	}
	return reply.ListApplications{
		List:  result,
		Total: aplys[0].Total,
	}, nil
}
