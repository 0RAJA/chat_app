package logic

import (
	"errors"
	"fmt"

	"github.com/0RAJA/Rutils/pkg/app/errcode"
	encode "github.com/0RAJA/Rutils/pkg/password"
	"github.com/0RAJA/chat_app/src/dao"
	db "github.com/0RAJA/chat_app/src/dao/postgres/sqlc"
	"github.com/0RAJA/chat_app/src/global"
	mid "github.com/0RAJA/chat_app/src/middleware"
	"github.com/0RAJA/chat_app/src/model/reply"
	"github.com/0RAJA/chat_app/src/myerr"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

type user struct {
}

// TODO: 用户的登陆

// 通过ID获取用户信息
func getUserInfo(c *gin.Context, userID int64) (*db.User, errcode.Err) {
	userInfo, err := dao.Group.DB.GetUserByID(c, userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, myerr.UserNotFound
		}
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	return userInfo, nil
}

func (user) CreateUser(c *gin.Context, emailStr, pwd, code string) (*reply.CreateUser, errcode.Err) {
	// 判断邮箱没有被注册
	if err := CheckEmailNotExists(c, emailStr); err != nil {
		return nil, err
	}
	// 校验验证码
	if !global.EmailMark.CheckCode(emailStr, code) {
		return nil, myerr.EmailCodeNotValid
	}
	hashPassword, err := encode.HashPassword(pwd)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	userInfo, err := dao.Group.DB.CreateUser(c, &db.CreateUserParams{
		Email:    emailStr,
		Password: hashPassword,
	})
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	// 添加邮箱到redis
	if err := dao.Group.Redis.AddEmails(c, emailStr); err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		reTry("addEmail:"+emailStr, func() error { return dao.Group.Redis.AddEmails(c, emailStr) })
	}
	// TODO: 注册后返回token
	return &reply.CreateUser{
		ID:       userInfo.ID,
		Email:    userInfo.Email,
		CreateAt: userInfo.CreateAt,
	}, nil
}

func (user) DeleteUser(c *gin.Context, userID int64) errcode.Err {
	userInfo, err := getUserInfo(c, userID)
	if err != nil {
		return err
	}
	if err := dao.Group.DB.DeleteUser(c, userID); err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	// 从redis删除邮箱
	if err := dao.Group.Redis.DeleteEmail(c, userInfo.Email); err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		reTry("delEmail:"+userInfo.Email, func() error { return dao.Group.Redis.DeleteEmail(c, userInfo.Email) })
	}
	return nil
}

func (user) UpdateUserEmail(c *gin.Context, userID int64, newEmail, code string) errcode.Err {
	// 判断邮箱是否更改
	userInfo, err := getUserInfo(c, userID)
	if err != nil {
		return err
	}
	if userInfo.Email == newEmail {
		return myerr.EmailSame
	}
	// 判断邮箱没有被注册
	if err := CheckEmailNotExists(c, newEmail); err != nil {
		return err
	}
	// 校验验证码
	if !global.EmailMark.CheckCode(newEmail, code) {
		return myerr.EmailCodeNotValid
	}
	if err := dao.Group.DB.UpdateUser(c, &db.UpdateUserParams{
		Email:    newEmail,
		Password: userInfo.Password,
		ID:       userID,
	}); err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	// 更新邮箱到redis
	if err := dao.Group.Redis.UpdateEmail(c, userInfo.Email, newEmail); err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		reTry(fmt.Sprintf("updateEmail:%s,%s", userInfo.Email, newEmail), func() error { return dao.Group.Redis.UpdateEmail(c, userInfo.Email, newEmail) })
	}
	return nil
}

func (user) UpdateUserPassword(c *gin.Context, userID int64, oldPwd, newPwd string) errcode.Err {
	// 检查旧密码是否匹配
	userInfo, merr := getUserInfo(c, userID)
	if merr != nil {
		return merr
	}
	if err := encode.CheckPassword(oldPwd, userInfo.Password); err != nil {
		return myerr.PasswordNotValid
	}
	hashPassword, err := encode.HashPassword(newPwd)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	// 更新
	if err := dao.Group.DB.UpdateUser(c, &db.UpdateUserParams{
		Email:    userInfo.Email,
		Password: hashPassword,
		ID:       userID,
	}); err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	return nil
}
