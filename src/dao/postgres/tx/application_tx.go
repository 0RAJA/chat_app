package tx

import (
	"context"

	db "github.com/0RAJA/chat_app/src/dao/postgres/sqlc"
	"github.com/0RAJA/chat_app/src/dao/redis/query"
	"github.com/0RAJA/chat_app/src/pkg/tool"
)

// AcceptApplicationTx 接受申请并建立好友关系和双方关系设置并添加到redis
func (store *SqlStore) AcceptApplicationTx(c context.Context, rdb *query.Queries, account1, account2 *db.GetAccountByIDRow) error {
	return store.execTx(c, func(queries *db.Queries) error {
		var err error
		err = tool.DoThat(err, func() error {
			return queries.UpdateApplication(c, &db.UpdateApplicationParams{
				Account1ID: account1.ID,
				Account2ID: account2.ID,
				Status:     db.ApplicationstatusValue1,
			})
		})
		id1, id2 := account1.ID, account2.ID
		if id1 > id2 {
			id1, id2 = id2, id1
		}
		var relationID int64
		err = tool.DoThat(err, func() error {
			relationID, err = queries.CreateFriendRelation(c, &db.CreateFriendRelationParams{Account1ID: id1, Account2ID: id2})
			return err
		})
		err = tool.DoThat(err, func() error {
			return queries.CreateSetting(c, &db.CreateSettingParams{
				AccountID:  account1.ID,
				RelationID: relationID,
				IsLeader:   false,
			})
		})
		err = tool.DoThat(err, func() error {
			return queries.CreateSetting(c, &db.CreateSettingParams{
				AccountID:  account2.ID,
				RelationID: relationID,
				IsLeader:   false,
			})
		})
		// 添加关系到redis
		err = tool.DoThat(err, func() error { return rdb.AddRelationAccount(c, relationID, account1.ID, account2.ID) })
		return err
	})
}
