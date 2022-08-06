package db

import (
	"context"

	"github.com/0RAJA/chat_app/src/pkg/tool"
)

// UpdateMsgTopTrueByMsgIDWithTx 更新此消息置顶(会删除其他置顶)
func (store *SqlStore) UpdateMsgTopTrueByMsgIDWithTx(c context.Context, relationID, msgID int64) error {
	return store.execTx(c, func(queries *Queries) error {
		var err error
		err = tool.DoThat(err, func() error { return queries.UpdateMsgTopFalseByRelationID(c, relationID) })
		err = tool.DoThat(err, func() error { return queries.UpdateMsgTopTrueByMsgID(c, msgID) })
		return err
	})
}

// RevokeMsgWithTx 撤回消息，如果消息置顶或pin则统统取消
func (store *SqlStore) RevokeMsgWithTx(c context.Context, msgID int64, isTop, isPin bool) error {
	return store.execTx(c, func(queries *Queries) error {
		var err error
		err = tool.DoThat(err, func() error {
			return queries.UpdateMsgRevoke(c, &UpdateMsgRevokeParams{
				ID:       msgID,
				IsRevoke: true,
			})
		})
		if isTop {
			err = tool.DoThat(err, func() error { return queries.UpdateMsgTopFalseByMsgID(c, msgID) })
		}
		if isPin {
			err = tool.DoThat(err, func() error {
				return queries.UpdateMsgPin(c, &UpdateMsgPinParams{
					ID:    msgID,
					IsPin: false,
				})
			})
		}
		return err
	})
}
