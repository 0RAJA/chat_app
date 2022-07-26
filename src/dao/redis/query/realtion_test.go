package query_test

import (
	"context"
	"sort"
	"testing"

	"github.com/0RAJA/Rutils/pkg/utils"
	"github.com/0RAJA/chat_app/src/dao"
	"github.com/stretchr/testify/require"
)

func TestQueries_AddGroupRelationAccount(t *testing.T) {
	tests := []struct {
		name string
		f    func()
	}{
		{
			name: "good",
			f: func() {
				require.NoError(t, dao.Group.Redis.DelAllGroupRelation(context.Background()))
				groupNum := utils.RandomInt(1, 10)
				groupMap := make(map[int64][]int64, groupNum)
				for i := int64(0); i < groupNum; i++ {
					accountNum := utils.RandomInt(1, 10)
					groupMap[i] = make([]int64, accountNum)
					for j := int64(0); j < accountNum; j++ {
						groupMap[i][j] = j
					}
				}
				require.NoError(t, dao.Group.Redis.ReloadGroupRelationIDs(context.Background(), groupMap))
				for i := int64(0); i < groupNum; i++ {
					accounts, err := dao.Group.Redis.GetAccountsByGroupRelationID(context.Background(), i)
					require.NoError(t, err)
					sort.Slice(accounts, func(i, j int) bool { return accounts[i] < accounts[j] })
					require.EqualValues(t, groupMap[i], accounts)
				}
				accounts, err := dao.Group.Redis.GetAccountsByGroupRelationID(context.Background(), -1)
				require.NoError(t, err)
				require.Empty(t, accounts)
				groupIdx := utils.RandomInt(0, groupNum)
				require.NoError(t, dao.Group.Redis.DelGroupRelation(context.Background(), groupIdx))
				accounts, err = dao.Group.Redis.GetAccountsByGroupRelationID(context.Background(), groupIdx)
				require.NoError(t, err)
				require.Empty(t, accounts)
				accountNum := utils.RandomInt(1, 10)
				accountIDs := make([]int64, 0, accountNum)
				for i := int64(0); i < accountNum; i++ {
					accountIDs = append(accountIDs, i)
				}
				require.NoError(t, dao.Group.Redis.AddGroupRelationAccount(context.Background(), groupIdx, accountIDs...))
				result, err := dao.Group.Redis.GetAccountsByGroupRelationID(context.Background(), groupIdx)
				require.NoError(t, err)
				sort.Slice(accounts, func(i, j int) bool { return accounts[i] < accounts[j] })
				require.EqualValues(t, accountIDs, result)
				accountID := utils.RandomInt(0, accountNum-1)
				require.NoError(t, dao.Group.Redis.DelGroupRelationAccount(context.Background(), groupIdx, accountID))
				result, err = dao.Group.Redis.GetAccountsByGroupRelationID(context.Background(), groupIdx)
				require.NoError(t, err)
				sort.Slice(accounts, func(i, j int) bool { return accounts[i] < accounts[j] })
				require.Len(t, result, int(accountNum-1))
				for i := range result {
					if result[i] == accountID {
						require.Fail(t, "accountID should not be in result")
					}
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f()
		})
	}
}
