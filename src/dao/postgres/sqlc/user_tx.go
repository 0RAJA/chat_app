package db

import (
	"context"

	"github.com/0RAJA/chat_app/src/pkg/tool"
)

// DeleteUserTx 删除用户和他的所有账户并删除与之相关的好友关系
func (store *SqlStore) DeleteUserTx(c context.Context, userID int64) error {
	return store.execTx(c, func(queries *Queries) error {
		var err error
		accountIDs, err := queries.DeleteAccountsByUserID(c, userID)
		if err != nil {
			return err
		}
		err = tool.DoThat(err, func() error { return queries.DeleteFriendRelationsByAccountIDs(c, accountIDs) })
		err = tool.DoThat(err, func() error { return queries.DeleteUser(c, userID) })
		return err
	})
}
