package db

import (
	"context"

	"github.com/0RAJA/chat_app/src/pkg/tool"
)

// AcceptApplicationTx 接受申请并建立好友关系和双方关系设置
func (store *SqlStore) AcceptApplicationTx(c context.Context, account1, account2 *GetAccountByIDRow) error {
	return store.execTx(c, func(queries *Queries) error {
		var err error
		err = tool.DoThat(err, func() error {
			return queries.UpdateApplication(c, &UpdateApplicationParams{
				Account1ID: account1.ID,
				Account2ID: account2.ID,
				Status:     ApplicationstatusValue1,
			})
		})
		id1, id2 := account1.ID, account2.ID
		if id1 > id2 {
			id1, id2 = id2, id1
		}
		var relationID int64
		err = tool.DoThat(err, func() error {
			relationID, err = queries.CreateFriendRelation(c, &CreateFriendRelationParams{Account1ID: id1, Account2ID: id2})
			return err
		})
		err = tool.DoThat(err, func() error {
			return queries.CreateSetting(c, &CreateSettingParams{
				AccountID:  account1.ID,
				RelationID: relationID,
				IsLeader:   false,
			})
		})
		err = tool.DoThat(err, func() error {
			return queries.CreateSetting(c, &CreateSettingParams{
				AccountID:  account2.ID,
				RelationID: relationID,
				IsLeader:   false,
			})
		})
		return err
	})
}
