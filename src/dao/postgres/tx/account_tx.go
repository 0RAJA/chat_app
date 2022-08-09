package tx

import (
	"context"
	"errors"

	db "github.com/0RAJA/chat_app/src/dao/postgres/sqlc"
	"github.com/0RAJA/chat_app/src/dao/redis/query"
	"github.com/0RAJA/chat_app/src/pkg/tool"
)

var (
	ErrAccountOverNum     = errors.New("账户数量超过限制")
	ErrAccountNameExists  = errors.New("账户名已存在")
	ErrAccountGroupLeader = errors.New("账户是群主")
)

// CreateAccountWithTx 检查数量，账户名后创建账户并建立和自己的关系
func (store *SqlStore) CreateAccountWithTx(c context.Context, rdb *query.Queries, maxAccountNum int32, arg *db.CreateAccountParams) error {
	return store.execTx(c, func(queries *db.Queries) error {
		var err error
		var accountNum int32
		// 检查数量
		err = tool.DoThat(err, func() error {
			accountNum, err = queries.CountAccountByUserID(c, arg.UserID)
			return err
		})
		if accountNum >= maxAccountNum {
			return ErrAccountOverNum
		}
		// 检查账户名
		var exists bool
		err = tool.DoThat(err, func() error {
			exists, err = queries.ExistsAccountByNameAndUserID(c, &db.ExistsAccountByNameAndUserIDParams{UserID: arg.UserID, Name: arg.Name})
			return err
		})
		if exists {
			return ErrAccountNameExists
		}
		err = tool.DoThat(err, func() error { return queries.CreateAccount(c, arg) })
		var relationID int64
		err = tool.DoThat(err, func() error {
			relationID, err = queries.CreateFriendRelation(c, &db.CreateFriendRelationParams{Account1ID: arg.ID, Account2ID: arg.ID})
			return err
		})
		err = tool.DoThat(err, func() error {
			return queries.CreateSetting(c, &db.CreateSettingParams{
				AccountID:  arg.ID,
				RelationID: relationID,
				IsSelf:     true,
			})
		})
		// 添加自己一个人的关系到redis
		err = tool.DoThat(err, func() error { return rdb.AddRelationAccount(c, relationID, arg.ID) })
		return err
	})
}

// DeleteAccountWithTx 删除账户并删除与之相关的关系
func (store *SqlStore) DeleteAccountWithTx(c context.Context, rdb *query.Queries, accountID int64) error {
	return store.execTx(c, func(queries *db.Queries) error {
		var err error
		// 判断是否是群主
		var isLeader bool
		err = tool.DoThat(err, func() error {
			isLeader, err = queries.ExistsGroupLeaderByAccountIDWithLock(c, accountID)
			return err
		})
		if isLeader {
			return ErrAccountGroupLeader
		}
		// 删除好友
		var friendRelationIDs []int64
		err = tool.DoThat(err, func() error {
			friendRelationIDs, err = queries.DeleteFriendRelationsByAccountID(c, accountID)
			return err
		})
		// 删除群
		var groupRelationIDs []int64
		err = tool.DoThat(err, func() error {
			groupRelationIDs, err = queries.DeleteSettingsByAccountID(c, accountID)
			return err
		})
		// 删除账户
		err = tool.DoThat(err, func() error { return queries.DeleteAccount(c, accountID) })
		// 从redis中删除对应的关系
		err = tool.DoThat(err, func() error { return rdb.DelRelations(c, friendRelationIDs...) })
		err = tool.DoThat(err, func() error { return rdb.DelAccountFromRelations(c, accountID, groupRelationIDs...) })
		return err
	})
}
