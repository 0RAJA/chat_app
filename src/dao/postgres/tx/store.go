package tx

import (
	"context"
	"fmt"

	db "github.com/0RAJA/chat_app/src/dao/postgres/sqlc"
	"github.com/0RAJA/chat_app/src/dao/redis/query"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type TXer interface {
	// CreateApplicationTx 先判断是否存在申请，不存在则创建申请
	CreateApplicationTx(c context.Context, arg *db.CreateApplicationParams) error
	// AcceptApplicationTx account1接受account2申请并建立好友关系和双方的关系设置同时发送消息通知并添加到redis
	AcceptApplicationTx(c context.Context, rdb *query.Queries, account1, account2 *db.GetAccountByIDRow) (*db.Message, error)
	// CreateAccountWithTx 创建账户并建立和自己的关系
	CreateAccountWithTx(c context.Context, rdb *query.Queries, maxAccountNum int32, arg *db.CreateAccountParams) error
	// DeleteAccountWithTx 删除账户并删除与之相关的好友关系
	DeleteAccountWithTx(c context.Context, rdb *query.Queries, accountID int64) error
	// UploadGroupAvatar 创建群组头像文件
	UploadGroupAvatar(c context.Context, arg db.CreateFileParams) error
	// DissolveGroup 删除群关系并删除所有群员
	DissolveGroup(c context.Context, relationID int64) error
	// UpdateMsgTopTrueByMsgIDWithTx 更新此消息置顶(会删除其他置顶)
	UpdateMsgTopTrueByMsgIDWithTx(c context.Context, relationID, msgID int64) error
	// RevokeMsgWithTx 撤回消息，如果消息置顶或pin则统统取消
	RevokeMsgWithTx(c context.Context, msgID int64, isTop, isPin bool) error
	// DeleteRelationWithTx 从数据库中删除关系并删除redis中的关系
	DeleteRelationWithTx(c context.Context, rdb *query.Queries, relationID int64) error
	// AddSettingWithTx 向数据库和redis中同时添加群员
	AddSettingWithTx(c context.Context, rdb *query.Queries, relationID, accountID int64, isLeader bool) error
	// DeleteSettingWithTx 向数据库和redis中同时删除群员
	DeleteSettingWithTx(c context.Context, rdb *query.Queries, relationID int64, accountID int64) error
	// TransferGroup 转让群
	TransferGroup(c context.Context, relationID int64, fID int64, tID int64) error
}

type SqlStore struct {
	*db.Queries
	DB *pgxpool.Pool
}

// 通过事务执行回调函数
func (store *SqlStore) execTx(ctx context.Context, fn func(queries *db.Queries) error) error {
	tx, err := store.DB.BeginTx(ctx, pgx.TxOptions{
		IsoLevel:       pgx.ReadCommitted,
		AccessMode:     pgx.ReadWrite,
		DeferrableMode: pgx.Deferrable,
	})
	if err != nil {
		return err
	}
	q := store.WithTx(tx) // 使用开启的事务创建一个查询
	if err := fn(q); err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err:%v,rb err:%v", err, rbErr)
		}
		return err
	}
	return tx.Commit(ctx)
}
