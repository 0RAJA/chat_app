package db

import (
	"context"

	"github.com/0RAJA/chat_app/src/pkg/setting"
)

// CreateAccountTx 创建账户并建立和自己的关系
func (store *SqlStore) CreateAccountTx(c context.Context, arg *CreateAccountParams) error {
	return store.execTx(c, func(queries *Queries) error {
		var err error
		err = setting.DoThat(err, func() error { return queries.CreateAccount(c, arg) })
		var relationID int64
		err = setting.DoThat(err, func() error {
			relationID, err = queries.CreateFriendRelation(c, &CreateFriendRelationParams{Account1ID: arg.ID, Account2ID: arg.ID})
			return err
		})
		err = setting.DoThat(err, func() error {
			return queries.CreateSetting(c, &CreateSettingParams{
				AccountID:  arg.ID,
				RelationID: relationID,
				NickName:   arg.Name,
				IsLeader:   false,
			})
		})
		return err
	})
}

// DeleteAccountWithTx 删除账户并删除与之相关的好友关系
func (store *SqlStore) DeleteAccountWithTx(c context.Context, accountID int64) error {
	return store.execTx(c, func(queries *Queries) error {
		var err error
		err = setting.DoThat(err, func() error { return queries.DeleteFriendRelationsByAccountID(c, accountID) })
		return setting.DoThat(err, func() error { return queries.DeleteAccount(c, accountID) })
	})
}
