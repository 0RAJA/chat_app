package db

import (
	"context"
	"errors"
)

var ErrApplicationExists = errors.New("申请已存在")

// CreateApplicationTx 使用事务先判断是否存在申请，不存在则创建申请
func (store *SqlStore) CreateApplicationTx(c context.Context, arg *CreateApplicationParams) error {
	return store.execTx(c, func(queries *Queries) error {
		// 查看申请是否存在
		ok, err := queries.ExistsApplicationByIDWithLock(c, &ExistsApplicationByIDWithLockParams{
			Account1ID: arg.Account1ID,
			Account2ID: arg.Account2ID,
		})
		if err != nil {
			return err
		}
		if ok {
			return ErrApplicationExists
		}
		// 创建申请
		return queries.CreateApplication(c, arg)
	})
}
