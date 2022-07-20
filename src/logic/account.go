package logic

import (
	"errors"

	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/0RAJA/chat_app/src/dao"
	db "github.com/0RAJA/chat_app/src/dao/postgres/sqlc"
	"github.com/0RAJA/chat_app/src/global"
	mid "github.com/0RAJA/chat_app/src/middleware"
	"github.com/0RAJA/chat_app/src/model"
	"github.com/0RAJA/chat_app/src/model/common"
	"github.com/0RAJA/chat_app/src/model/reply"
	"github.com/0RAJA/chat_app/src/myerr"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

type account struct {
}

func (account) CreateAccount(c *gin.Context, userID int64, name, avatar string) (*reply.CreateAccount, errcode.Err) {
	// 判断账户数量
	accountNum, err := dao.Group.DB.CountAccountByUserID(c, userID)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	if accountNum >= global.PbSettings.Rule.AccountNumMax {
		return nil, myerr.AccountNumExcessive
	}
	// 判断账户名称是否存在
	ok, err := dao.Group.DB.ExistsAccountByNameAndUserID(c, &db.ExistsAccountByNameAndUserIDParams{
		UserID: userID,
		Name:   name,
	})
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	if ok {
		return nil, myerr.AccountNameExists
	}
	arg := &db.CreateAccountParams{
		ID:     global.GenID.GetID(),
		UserID: userID,
		Name:   name,
		Avatar: avatar,
	}
	// 创建账户
	if err := dao.Group.DB.CreateAccount(c, arg); err != nil {
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
func getAccountInfoByID(c *gin.Context, accountID int64) (*db.Account, errcode.Err) {
	accountInfo, err := dao.Group.DB.GetAccountByID(c, accountID)
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
	accountInfo, merr := getAccountInfoByID(c, accountID)
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
	accountInfo, merr := getAccountInfoByID(c, accountID)
	if merr != nil {
		return merr
	}
	if accountInfo.UserID != userID {
		return myerr.AuthPermissionsInsufficient
	}
	if err := dao.Group.DB.DeleteAccount(c, accountID); err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	return nil
}

func (account) UpdateAccount(c *gin.Context, accountID int64, name, avatar, gender, signature string) errcode.Err {
	err := dao.Group.DB.UpdateAccount(c, &db.UpdateAccountParams{
		Name:      name,
		Avatar:    avatar,
		Gender:    db.Gender(gender),
		Signature: signature,
		ID:        accountID,
	})
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	return nil
}

func (account) GetAccountByID(c *gin.Context, accountID int64) (*reply.GetAccountByID, errcode.Err) {
	accountInfo, merr := getAccountInfoByID(c, accountID)
	if merr != nil {
		return nil, merr
	}
	return &reply.GetAccountByID{
		ID:        accountInfo.ID,
		Name:      accountInfo.Name,
		Avatar:    accountInfo.Avatar,
		Gender:    string(accountInfo.Gender),
		Signature: accountInfo.Signature,
		CreateAt:  accountInfo.CreateAt,
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
		})
	}
	return reply.GetAccountsByUserID{
		List:  result,
		Total: int64(len(result)),
	}, nil
}

func (account) GetAccountsByName(c *gin.Context, name string, limit, offset int32) (reply.GetAccountsByName, errcode.Err) {
	accounts, err := dao.Group.DB.GetAccountsByName(c, &db.GetAccountsByNameParams{Name: name, Limit: limit, Offset: offset})
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return reply.GetAccountsByName{}, errcode.ErrServer
	}
	var total int64
	result := make([]*reply.AccountInfo, 0, len(accounts))
	for _, v := range accounts {
		result = append(result, &reply.AccountInfo{
			ID:     v.ID,
			Name:   v.Name,
			Avatar: v.Avatar,
		})
	}
	if len(accounts) > 0 {
		total = accounts[0].Total
	}
	return reply.GetAccountsByName{
		List:  result,
		Total: total,
	}, nil
}
