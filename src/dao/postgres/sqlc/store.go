package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Store interface {
	Querier
	TXer
}

type TXer interface {
	// CreateApplicationTx 先判断是否存在申请，不存在则创建申请
	CreateApplicationTx(c context.Context, arg *CreateApplicationParams) error
	// AcceptApplicationTx account1接受account2申请并建立好友关系和双方的关系设置
	AcceptApplicationTx(c context.Context, account1, account2 *GetAccountByIDRow) error
	// CreateAccountTx 创建账户并建立和自己的关系
	CreateAccountTx(c context.Context, arg *CreateAccountParams) error
	// DeleteAccountWithTx 删除账户并删除与之相关的好友关系
	DeleteAccountWithTx(c context.Context, accountID int64) error
	// DeleteUserTx 删除用户和他的所有账户并删除与之相关的好友关系
	DeleteUserTx(c context.Context, userID int64) error
	// UploadGroupAvatar 创建群组头像文件
	UploadGroupAvatar(c context.Context,arg CreateFileParams) error
	// DissolveGroup 删除群关系并删除所有群员
	DissolveGroup(c context.Context ,relationID int64) error
	// UpdateMsgTopTrueByMsgIDWithTx 更新此消息置顶(会删除其他置顶)
	UpdateMsgTopTrueByMsgIDWithTx(c context.Context, relationID, msgID int64) error
	// RevokeMsgWithTx 撤回消息，如果消息置顶或pin则统统取消
	RevokeMsgWithTx(c context.Context, msgID int64, isTop, isPin bool) error
}

type SqlStore struct {
	*Queries
	DB *pgxpool.Pool
}

// 通过事务执行回调函数
func (store *SqlStore) execTx(ctx context.Context, fn func(queries *Queries) error) error {
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
