package query

import (
	"context"

	"github.com/0RAJA/Rutils/pkg/utils"
	"github.com/go-redis/redis/v8"
)

const KeyGroup = "KeyGroup"

const DelAllPrefixLua = "local redisKeys = redis.call('keys', KEYS[1] .. '*');for i, k in pairs(redisKeys) do redis.call('expire', k, 0);end"

// ReloadRelationIDs 重新加载群聊名单
func (q *Queries) ReloadRelationIDs(c context.Context, groupMap map[int64][]int64) error {
	if err := q.rdb.Eval(c, DelAllPrefixLua, []string{KeyGroup}).Err(); err != nil && err != redis.Nil {
		return err
	}
	pipe := q.rdb.TxPipeline()
	for relationID, ids := range groupMap {
		data := make([]interface{}, len(ids))
		for i, id := range ids {
			data[i] = utils.IDToSting(id)
		}
		q.rdb.SAdd(c, utils.LinkStr(KeyGroup, utils.IDToSting(relationID)), data...)
	}
	_, err := pipe.Exec(c)
	return err
}

// DelRelations 删除部分群聊名单
func (q *Queries) DelRelations(c context.Context, relationIDs ...int64) error {
	if len(relationIDs) == 0 {
		return nil
	}
	pipe := q.rdb.TxPipeline()
	for _, relationID := range relationIDs {
		pipe.Del(c, utils.LinkStr(KeyGroup, utils.IDToSting(relationID)))
	}
	_, err := pipe.Exec(c)
	return err
}

// AddRelationAccount 向群聊名单中增加人员
func (q *Queries) AddRelationAccount(c context.Context, relationID int64, accountIDs ...int64) error {
	if len(accountIDs) == 0 {
		return nil
	}
	data := make([]interface{}, len(accountIDs))
	for i, v := range accountIDs {
		data[i] = utils.IDToSting(v)
	}
	return q.rdb.SAdd(c, utils.LinkStr(KeyGroup, utils.IDToSting(relationID)), data...).Err()
}

// DelRelationAccount 从一个群中删除多个人员
func (q *Queries) DelRelationAccount(c context.Context, relationID int64, accountIDs ...int64) error {
	if len(accountIDs) == 0 {
		return nil
	}
	data := make([]interface{}, len(accountIDs))
	for i, v := range accountIDs {
		data[i] = utils.IDToSting(v)
	}
	return q.rdb.SRem(c, utils.LinkStr(KeyGroup, utils.IDToSting(relationID)), data...).Err()
}

// DelAccountFromRelations 从多个群中删除一个人员
func (q *Queries) DelAccountFromRelations(c context.Context, accountID int64, relationIDs ...int64) error {
	if len(relationIDs) == 0 {
		return nil
	}
	pipe := q.rdb.TxPipeline()
	for _, relationID := range relationIDs {
		pipe.SRem(c, utils.LinkStr(KeyGroup, utils.IDToSting(relationID)), utils.IDToSting(accountID))
	}
	_, err := pipe.Exec(c)
	return err
}

// GetAccountsByRelationID 获取群聊名单中的所有人员
func (q *Queries) GetAccountsByRelationID(c context.Context, relationID int64) ([]int64, error) {
	data, err := q.rdb.SMembers(c, utils.LinkStr(KeyGroup, utils.IDToSting(relationID))).Result()
	if err != nil {
		return nil, err
	}
	ret := make([]int64, len(data))
	for i, v := range data {
		ret[i] = utils.StringToIDMust(v)
	}
	return ret, nil
}

// DelAllRelations 删除所有群聊名单(用于测试)
func (q *Queries) DelAllRelations(c context.Context) error {
	if err := q.rdb.Eval(c, DelAllPrefixLua, []string{KeyGroup}).Err(); err != nil && err != redis.Nil {
		return err
	}
	return nil
}
