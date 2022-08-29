package tx

import (
	"context"
	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/0RAJA/chat_app/src/dao"
	"github.com/0RAJA/chat_app/src/global"

	db "github.com/0RAJA/chat_app/src/dao/postgres/sqlc"
	"github.com/0RAJA/chat_app/src/dao/redis/query"
	"github.com/0RAJA/chat_app/src/pkg/tool"
)

// AddSettingWithTx 向数据库和redis中同时添加群员
func (store *SqlStore) AddSettingWithTx(c context.Context, rdb *query.Queries, accountID int64, isLeader bool, name, desc string) error {
	return store.execTx(c, func(queries *db.Queries) error {
		relationID, err := dao.Group.DB.CreateGroupRelation(c, &db.CreateGroupRelationParams{
			Name:        name,
			Description: desc,
			Avatar:      global.PbSettings.Rule.DefaultAvatarURL,
		})
		if err != nil {
			return errcode.ErrServer
		}
		err = queries.CreateSetting(c, &db.CreateSettingParams{
			AccountID:  accountID,
			RelationID: relationID,
			IsLeader:   isLeader,
			IsSelf:     false,
		})
		if err != nil {
			return err
		}
		return rdb.AddRelationAccount(c, relationID, accountID)
	})
}

// DeleteSettingWithTx 向数据库和redis中同时删除群员
func (store *SqlStore) DeleteSettingWithTx(c context.Context, rdb *query.Queries, relationID int64, accountID int64) error {
	return store.execTx(c, func(queries *db.Queries) error {
		err := queries.DeleteSetting(c, &db.DeleteSettingParams{
			AccountID:  accountID,
			RelationID: relationID,
		})
		if err != nil {
			return err
		}
		return rdb.DelRelationAccount(c, relationID, accountID)
	})
}

// TransferGroup 转让群
func (store *SqlStore) TransferGroup(c context.Context, relationID int64, fID int64, tID int64) error {
	return store.execTx(c, func(queries *db.Queries) error {
		var err error
		err = tool.DoThat(err, func() error {
			return queries.TransferIsSelfFalse(c, &db.TransferIsSelfFalseParams{
				RelationID: relationID,
				AccountID:  fID,
			})
		})
		err = tool.DoThat(err, func() error {
			return queries.TransferIsSelfTrue(c, &db.TransferIsSelfTrueParams{
				RelationID: relationID,
				AccountID:  tID,
			})
		})
		return err
	})
}
