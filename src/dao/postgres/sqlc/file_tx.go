package db

import (
	"context"
	"database/sql"
)
// UploadGroupAvatar 创建群组头像文件
func (store *SqlStore)UploadGroupAvatar(c context.Context,arg CreateFileParams) error {
	return store.execTx(c, func(queries *Queries) error {
		var err error
		_,err = queries.GetGroupAvatar(c,arg.RelationID)
		if err != nil {
			if err.Error() == "no rows in result set"{
				_,err = queries.CreateFile(c,&CreateFileParams{
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
			err = queries.UpdateGroupAvatar(c,&UpdateGroupAvatarParams{
				Url:        "",
				RelationID: arg.RelationID,
			})
		}
		return err
	})
}
