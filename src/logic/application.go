package logic

import (
	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/0RAJA/chat_app/src/dao"
	db "github.com/0RAJA/chat_app/src/dao/postgres/sqlc"
	"github.com/0RAJA/chat_app/src/global"
	mid "github.com/0RAJA/chat_app/src/middleware"
	"github.com/0RAJA/chat_app/src/model/reply"
	"github.com/0RAJA/chat_app/src/myerr"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

type application struct {
}

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
	// TODO:检查两个账户是否已经有了联系
	if account1ID == account2ID {
		return myerr.ApplicationNotValid
	}
	err := dao.Group.DB.CreateApplicationTx(c, &db.CreateApplicationParams{Account1ID: account1ID, Account2ID: account2ID, ApplyMsg: applyMsg})
	switch err {
	case nil:
		return nil
	case db.ErrApplicationExists:
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

func (application) Accept(c *gin.Context, account1ID, account2ID int64) errcode.Err {
	aply, merr := getApplication(c, account2ID, account1ID)
	if merr != nil {
		return merr
	}
	if aply.Status == db.ApplicationstatusValue1 {
		return myerr.ApplicationRepeatOpt
	}
	if err := dao.Group.DB.UpdateApplication(c, &db.UpdateApplicationParams{
		Account1ID: account2ID,
		Account2ID: account1ID,
		Status:     db.ApplicationstatusValue1,
	}); err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	// TODO: 创建聊天联系
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
		return reply.ListApplications{}, nil
	}
	result := make([]*reply.ApplicationInfo, 0, len(aplys))
	for _, aply := range aplys {
		result = append(result, &reply.ApplicationInfo{
			Account1ID:     aply.Account1ID,
			Account2ID:     aply.Account2ID,
			ApplyMsg:       aply.ApplyMsg,
			RefuseMsg:      aply.RefuseMsg,
			Status:         string(aply.Status),
			CreateAt:       aply.CreateAt,
			UpdateAt:       aply.UpdateAt,
			Account1Avatar: aply.Account2Avatar,
			Account1Name:   aply.Account1Name,
			Account2Avatar: aply.Account2Avatar,
			Account2Name:   aply.Account2Name,
		})
	}
	return reply.ListApplications{
		List:  result,
		Total: aplys[0].Total,
	}, nil
}
