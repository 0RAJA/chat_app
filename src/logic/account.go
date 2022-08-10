package logic

import (
	"errors"

	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/0RAJA/chat_app/src/dao"
	db "github.com/0RAJA/chat_app/src/dao/postgres/sqlc"
	"github.com/0RAJA/chat_app/src/dao/postgres/tx"
	"github.com/0RAJA/chat_app/src/global"
	mid "github.com/0RAJA/chat_app/src/middleware"
	"github.com/0RAJA/chat_app/src/model"
	"github.com/0RAJA/chat_app/src/model/common"
	"github.com/0RAJA/chat_app/src/model/reply"
	"github.com/0RAJA/chat_app/src/myerr"
	"github.com/0RAJA/chat_app/src/task"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

type account struct {
}

func (account) CreateAccount(c *gin.Context, userID int64, name, avatar, gender, signature string) (*reply.CreateAccount, errcode.Err) {
	arg := &db.CreateAccountParams{
		ID:        global.GenID.GetID(),
		UserID:    userID,
		Name:      name,
		Avatar:    avatar,
		Gender:    db.Gender(gender),
		Signature: signature,
	}
	// 创建账户以及和自己的关系
	err := dao.Group.DB.CreateAccountWithTx(c, dao.Group.Redis, global.PbSettings.Rule.AccountNumMax, arg)
	switch err {
	case tx.ErrAccountOverNum:
		return nil, myerr.AccountNumExcessive
	case tx.ErrAccountNameExists:
		return nil, myerr.AccountNameExists
	case nil:
	default:
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	// 生成账户Token
	token, payload, err := newToken(model.AccountToken, arg.ID)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	return &reply.CreateAccount{
		ID:     arg.ID,
		Name:   arg.Name,
		Avatar: arg.Avatar,
		AccountToken: common.Token{
			Token:     token,
			ExpiredAt: payload.ExpiredAt,
		},
	}, nil
}

// 通过账户ID获取账户信息
func getAccountInfoByID(c *gin.Context, accountID, selfID int64) (*db.GetAccountByIDRow, errcode.Err) {
	accountInfo, err := dao.Group.DB.GetAccountByID(c, &db.GetAccountByIDParams{
		TargetID: accountID,
		SelfID:   selfID,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, myerr.AccountNotFound
		}
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	return accountInfo, nil
}

func (account) GetAccountToken(c *gin.Context, userID, accountID int64) (*reply.GetAccountToken, errcode.Err) {
	accountInfo, merr := getAccountInfoByID(c, accountID, accountID)
	if merr != nil {
		return nil, merr
	}
	if accountInfo.UserID != userID {
		return nil, myerr.AuthPermissionsInsufficient
	}
	token, payload, err := newToken(model.AccountToken, accountID)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	return &reply.GetAccountToken{
		AccountToken: common.Token{
			Token:     token,
			ExpiredAt: payload.ExpiredAt,
		},
	}, nil
}

func (account) DeleteAccount(c *gin.Context, userID, accountID int64) errcode.Err {
	accountInfo, merr := getAccountInfoByID(c, accountID, accountID)
	if merr != nil {
		return merr
	}
	if accountInfo.UserID != userID {
		return myerr.AuthPermissionsInsufficient
	}
	err := dao.Group.DB.DeleteAccountWithTx(c, dao.Group.Redis, accountID)
	switch err {
	case tx.ErrAccountGroupLeader:
		return myerr.AccountGroupLeader
	case nil:
		return nil
	default:
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
}

func (account) UpdateAccount(c *gin.Context, accountID int64, name, gender, signature string) errcode.Err {
	err := dao.Group.DB.UpdateAccount(c, &db.UpdateAccountParams{
		Name:      name,
		Gender:    db.Gender(gender),
		Signature: signature,
		ID:        accountID,
	})
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	accessToken, _ := mid.GetToken(c.Request.Header)
	// 推送更新信息
	global.Worker.SendTask(task.UpdateAccount(accessToken, accountID, name, gender, signature))
	return nil
}

func (account) GetAccountByID(c *gin.Context, accountID, selfID int64) (*reply.GetAccountByID, errcode.Err) {
	accountInfo, merr := getAccountInfoByID(c, accountID, selfID)
	if merr != nil {
		return nil, merr
	}
	return &reply.GetAccountByID{
		ID:         accountInfo.ID,
		Name:       accountInfo.Name,
		Avatar:     accountInfo.Avatar,
		Gender:     string(accountInfo.Gender),
		Signature:  accountInfo.Signature,
		CreateAt:   accountInfo.CreateAt,
		RelationID: accountInfo.RelationID.Int64,
	}, nil
}

func (account) GetAccountsByUserID(c *gin.Context, userID int64) (reply.GetAccountsByUserID, errcode.Err) {
	accounts, err := dao.Group.DB.GetAccountsByUserID(c, userID)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return reply.GetAccountsByUserID{}, errcode.ErrServer
	}
	result := make([]*reply.AccountInfo, 0, len(accounts))
	for _, v := range accounts {
		result = append(result, &reply.AccountInfo{
			ID:     v.ID,
			Name:   v.Name,
			Avatar: v.Avatar,
			Gender: string(v.Gender),
		})
	}
	return reply.GetAccountsByUserID{
		List:  result,
		Total: int64(len(result)),
	}, nil
}

func (account) GetAccountsByName(c *gin.Context, accountID int64, name string, limit, offset int32) (reply.GetAccountsByName, errcode.Err) {
	accounts, err := dao.Group.DB.GetAccountsByName(c, &db.GetAccountsByNameParams{AccountID: accountID, Name: name, Limit: limit, Offset: offset})
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return reply.GetAccountsByName{}, errcode.ErrServer
	}
	if len(accounts) == 0 {
		return reply.GetAccountsByName{}, nil
	}
	result := make([]*reply.AccountFriendInfo, 0, len(accounts))
	for _, v := range accounts {
		result = append(result, &reply.AccountFriendInfo{
			AccountInfo: reply.AccountInfo{ID: v.ID,
				Name:   v.Name,
				Avatar: v.Avatar,
				Gender: string(v.Gender),
			},
			RelationID: v.RelationID.Int64,
		})
	}
	return reply.GetAccountsByName{
		List:  result,
		Total: accounts[0].Total,
	}, nil
}
