package db

import (
	"context"

	"github.com/0RAJA/chat_app/src/pkg/tool"
)

func (store *SqlStore) UpdateMsgTopTrueByMsgIDWithTx(c context.Context, relationID, msgID int64) error {
	return store.execTx(c, func(queries *Queries) error {
		var err error
		err = tool.DoThat(err, func() error { return queries.UpdateMsgTopFalseByRelationID(c, relationID) })
		err = tool.DoThat(err, func() error { return queries.UpdateMsgTopTrueByMsgID(c, msgID) })
		return err
	})
}
