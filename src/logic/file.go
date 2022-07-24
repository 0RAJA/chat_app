package logic

import (
	"database/sql"
	"fmt"
	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/0RAJA/chat_app/src/dao"
	db "github.com/0RAJA/chat_app/src/dao/postgres/sqlc"
	"github.com/0RAJA/chat_app/src/global"
	mid "github.com/0RAJA/chat_app/src/middleware"
	"github.com/0RAJA/chat_app/src/model/reply"
	"github.com/0RAJA/chat_app/src/model/request"
	"github.com/0RAJA/chat_app/src/myerr"
	"github.com/0RAJA/chat_app/src/upload/oss"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type file struct {
}

func (file) PublishFile(c *gin.Context, params request.PublishFile, filetype string, file *multipart.FileHeader) (reply.PublishFile, errcode.Err) {
	var con = oss.Config{
		BucketUrl:       global.PvSettings.AliyunOSS.BucketUrl,
		BasePath:        global.PvSettings.AliyunOSS.BasePath,
		Endpoint:        global.PvSettings.AliyunOSS.Endpoint,
		AccessKeyId:     global.PvSettings.AliyunOSS.AccessKeyId,
		AccessKeySecret: global.PvSettings.AliyunOSS.AccessKeySecret,
		BucketName:      global.PvSettings.AliyunOSS.BucketName,
	}
	o := oss.Init(con)
	url, key, err := o.UploadFile(file)
	result := reply.PublishFile{}
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return result, myerr.FiledStore
	}

	r, err := dao.Group.DB.CreateFile(c, &db.CreateFileParams{
		FileName: params.File.Filename,
		FileType: db.Filetype(filetype),
		FileSize: params.File.Size,
		Key:      key,
		Url:      url,
		RelationID: sql.NullInt64{
			Int64: params.RelationID,
			Valid: true,
		},
		AccountID: sql.NullInt64{
			Int64: params.AccountID,
			Valid: true,
		},
	})
	if err != nil {
		fmt.Println(err)
		return result, errcode.ErrServer
	}
	result = reply.PublishFile{
		ID:       r.ID,
		FileType: filetype,
		FileSize: params.File.Size,
		Url:      url,
		CreateAt: r.CreateAt,
	}
	return result, nil
}

func (file) DeleteFile(c *gin.Context, fileID int64) (result reply.DeleteFile, mErr errcode.Err) {
	k, err := dao.Group.DB.GetFileKeyByID(c, &db.GetFileKeyByIDParams{FileID: fileID})
	if err != nil {
		if err != sql.ErrNoRows {
			return result, myerr.FileNotExist
		}
		return result, errcode.ErrServer
	}
	fmt.Println(k)
	key := k.(string)
	var con = oss.Config{
		BucketUrl:       global.PvSettings.AliyunOSS.BucketUrl,
		BasePath:        global.PvSettings.AliyunOSS.BasePath,
		Endpoint:        global.PvSettings.AliyunOSS.Endpoint,
		AccessKeyId:     global.PvSettings.AliyunOSS.AccessKeyId,
		AccessKeySecret: global.PvSettings.AliyunOSS.AccessKeySecret,
		BucketName:      global.PvSettings.AliyunOSS.BucketName,
	}
	o := oss.Init(con)
	r, err := o.DeleteFile(key)
	fmt.Println(r)
	if err != nil {
		return result, myerr.FileDeleteFailed
	}
	err = dao.Group.DB.DeleteFileByID(c, &db.DeleteFileByIDParams{
		FileID: fileID,
	})
	if err != nil {
		return result, errcode.ErrServer
	}
	return result, nil
}

func (file) GetRelationFile(c *gin.Context, relationID int64) ([]reply.File, errcode.Err) {
	list, err := dao.Group.DB.GetFileByRelationID(c, sql.NullInt64{Int64: relationID, Valid: true})
	result := make([]reply.File, 0, 20)
	if err != nil {

		if err != sql.ErrNoRows {
			return result, myerr.FileNotExist
		}
		return result, errcode.ErrServer

	}
	for _, v := range list {
		r := reply.File{
			FileID:    v.ID,
			FileName:  v.FileName,
			FileType:  string(v.FileType),
			FileSize:  v.FileSize,
			Url:       v.Url,
			AccountID: v.AccountID.Int64,
			CreateAt:  v.CreateAt,
		}
		result = append(result, r)
	}
	return result, nil
}
