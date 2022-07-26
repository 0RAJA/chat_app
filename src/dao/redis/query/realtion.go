package query

import (
	"context"

	"github.com/0RAJA/Rutils/pkg/utils"
	"github.com/go-redis/redis/v8"
)

const KeyGroup = "KeyGroup"

const DelAllPrefixLua = "local redisKeys = redis.call('keys', KEYS[1] .. '*');for i, k in pairs(redisKeys) do redis.call('expire', k, 0);end"

// ReloadGroupRelationIDs 重新加载群聊名单
func (q *Queries) ReloadGroupRelationIDs(c context.Context, groupMap map[int64][]int64) error {
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

// DelGroupRelation 删除一个群聊名单
func (q *Queries) DelGroupRelation(c context.Context, relationID int64) error {
	return q.rdb.Del(c, utils.LinkStr(KeyGroup, utils.IDToSting(relationID))).Err()
}

// AddGroupRelationAccount 向群聊名单中增加人员
func (q *Queries) AddGroupRelationAccount(c context.Context, relationID int64, accountIDs ...int64) error {
	if len(accountIDs) == 0 {
		return nil
	}
	data := make([]interface{}, len(accountIDs))
	for i, v := range accountIDs {
		data[i] = utils.IDToSting(v)
	}
	return q.rdb.SAdd(c, utils.LinkStr(KeyGroup, utils.IDToSting(relationID)), data...).Err()
}

// DelGroupRelationAccount 从群聊名单中删除人员
func (q *Queries) DelGroupRelationAccount(c context.Context, relationID int64, accountIDs ...int64) error {
	if len(accountIDs) == 0 {
		return nil
	}
	data := make([]interface{}, len(accountIDs))
	for i, v := range accountIDs {
		data[i] = utils.IDToSting(v)
	}
	return q.rdb.SRem(c, utils.LinkStr(KeyGroup, utils.IDToSting(relationID)), data...).Err()
}

// GetAccountsByGroupRelationID 获取群聊名单中的所有人员
func (q *Queries) GetAccountsByGroupRelationID(c context.Context, relationID int64) ([]int64, error) {
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

// DelAllGroupRelation 删除所有群聊名单(用于测试)
func (q *Queries) DelAllGroupRelation(c context.Context) error {
	if err := q.rdb.Eval(c, DelAllPrefixLua, []string{KeyGroup}).Err(); err != nil && err != redis.Nil {
		return err
	}
	return nil
}
