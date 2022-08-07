package tx

import (
	"context"
	"database/sql"

	db "github.com/0RAJA/chat_app/src/dao/postgres/sqlc"
)

// UploadGroupAvatar 创建群组头像文件
func (store *SqlStore) UploadGroupAvatar(c context.Context, arg db.CreateFileParams) error {
	return store.execTx(c, func(queries *db.Queries) error {
		var err error
		_, err = queries.GetGroupAvatar(c, arg.RelationID)
		if err != nil {
			if err.Error() == "no rows in result set" {
				_, err = queries.CreateFile(c, &db.CreateFileParams{
					FileName:   arg.FileName,
					FileType:   "img",
					FileSize:   arg.FileSize,
					Key:        arg.Key,
					Url:        arg.Url,
					RelationID: arg.RelationID,
					AccountID:  sql.NullInt64{},
				})
			} else {
				return err
			}
		} else {
			err = queries.UpdateGroupAvatar(c, &db.UpdateGroupAvatarParams{
				Url:        arg.Url,
				RelationID: arg.RelationID,
			})
		}
		data, err := queries.GetGroupRelationByID(c, arg.RelationID.Int64)
		if err != nil {
			return err
		}

		return queries.UpdateGroupRelation(c, &db.UpdateGroupRelationParams{
			Name:        data.Name,
			Description: data.Description,
			Avatar:      arg.Url,
			ID:          arg.RelationID.Int64,
		})
	})
}
