package logic

import (
	"context"
	"database/sql"
	"fmt"
	"mime/multipart"

	upload "github.com/0RAJA/Rutils/pkg/upload/oss"
	"github.com/0RAJA/Rutils/pkg/upload/oss/aliyun"
	"github.com/0RAJA/chat_app/src/model"
	"github.com/0RAJA/chat_app/src/pkg/gtype"

	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/0RAJA/chat_app/src/dao"
	db "github.com/0RAJA/chat_app/src/dao/postgres/sqlc"
	"github.com/0RAJA/chat_app/src/global"
	"github.com/0RAJA/chat_app/src/model/reply"
	"github.com/0RAJA/chat_app/src/myerr"
	"github.com/gin-gonic/gin"
)

type file struct {
}

var oss upload.OSS

// PublishFile 上传文件，传出context与relationID,accountID,file(*multipart.FileHeader),返回 model.PublishFileRe
// 错误代码 1003:系统错误 8001:文件存储失败(aly) 8004:文件过大
func PublishFile(c context.Context, params model.PublishFile) (model.PublishFileRe, errcode.Err) {
	result := model.PublishFileRe{}
	filetype, mErr := gtype.GetFileType(params.File)
	if mErr != nil {
		return result, errcode.ErrServer
	}
	if filetype != "img" && filetype != "png" && filetype != "jpg" {
		if params.File.Size > global.PbSettings.Rule.BiggestFileSize {
			return result, myerr.FileTooBig
		}
		filetype = "file"
	} else {
		filetype = "img"
	}
	oss = aliyun.Init(aliyun.Config{
		BucketUrl:       global.PvSettings.AliyunOSS.BucketUrl,
		BasePath:        global.PvSettings.AliyunOSS.BasePath,
		Endpoint:        global.PvSettings.AliyunOSS.Endpoint,
		AccessKeyId:     global.PvSettings.AliyunOSS.AccessKeyId,
		AccessKeySecret: global.PvSettings.AliyunOSS.AccessKeySecret,
		BucketName:      global.PvSettings.AliyunOSS.BucketName,
	})
	url, key, err := oss.UploadFile(params.File)
	if err != nil {
		global.Logger.Error(err.Error())
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
	result = model.PublishFileRe{
		ID:       r.ID,
		FileType: filetype,
		FileSize: params.File.Size,
		Url:      url,
		CreateAt: r.CreateAt,
	}
	return result, nil
}

func (file) DeleteFile(c context.Context, fileID int64) (result reply.DeleteFile, mErr errcode.Err) {
	key, err := dao.Group.DB.GetFileKeyByID(c, fileID)
	if err != nil {
		if err != sql.ErrNoRows {
			return result, myerr.FileNotExist
		}
		return result, errcode.ErrServer
	}
	oss = aliyun.Init(aliyun.Config{
		BucketUrl:       global.PvSettings.AliyunOSS.BucketUrl,
		BasePath:        global.PvSettings.AliyunOSS.BasePath,
		Endpoint:        global.PvSettings.AliyunOSS.Endpoint,
		AccessKeyId:     global.PvSettings.AliyunOSS.AccessKeyId,
		AccessKeySecret: global.PvSettings.AliyunOSS.AccessKeySecret,
		BucketName:      global.PvSettings.AliyunOSS.BucketName,
	})
	if key != "" {
		_, err = oss.DeleteFile(key)
		if err != nil {
			return result, myerr.FileDeleteFailed
		}
	}
	err = dao.Group.DB.DeleteFileByID(c, fileID)
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

func (file) UploadAccountAvatar(c *gin.Context, accountId int64, file *multipart.FileHeader) (reply.UploadAvatar, errcode.Err) {
	oss = aliyun.Init(aliyun.Config{
		BucketUrl:       global.PvSettings.AliyunOSS.BucketUrl,
		BasePath:        global.PvSettings.AliyunOSS.BasePath,
		Endpoint:        global.PvSettings.AliyunOSS.Endpoint,
		AccessKeyId:     global.PvSettings.AliyunOSS.AccessKeyId,
		AccessKeySecret: global.PvSettings.AliyunOSS.AccessKeySecret,
		BucketName:      global.PvSettings.AliyunOSS.BucketName,
	})
	url, key, err := oss.UploadFile(file)
	result := reply.UploadAvatar{}
	if err != nil {
		global.Logger.Error(err.Error())
		return result, myerr.FiledStore
	}
	err = dao.Group.DB.UploadAccountAvatar(c, accountId, url, key)
	return reply.UploadAvatar{Url: url}, nil
}
func (file) UploadGroupAvatar(c *gin.Context, file *multipart.FileHeader, relationID int64, accountID int64) (reply.UploadAvatar, errcode.Err) {
	var url, key string
	var err error
	result := reply.UploadAvatar{}
	t, err := dao.Group.DB.ExistsSetting(c, &db.ExistsSettingParams{
		AccountID:  accountID,
		RelationID: relationID,
	})
	if err != nil {
		global.Logger.Error(err.Error())
		return result, errcode.ErrServer
	}
	if !t {
		return result, myerr.NotGroupMember
	}
	oss = aliyun.Init(aliyun.Config{
		BucketUrl:       global.PvSettings.AliyunOSS.BucketUrl,
		BasePath:        global.PvSettings.AliyunOSS.BasePath,
		Endpoint:        global.PvSettings.AliyunOSS.Endpoint,
		AccessKeyId:     global.PvSettings.AliyunOSS.AccessKeyId,
		AccessKeySecret: global.PvSettings.AliyunOSS.AccessKeySecret,
		BucketName:      global.PvSettings.AliyunOSS.BucketName,
	})
	if file != nil {
		url, key, err = oss.UploadFile(file)
		if err != nil {
			global.Logger.Error(err.Error())
			return result, myerr.FiledStore
		}
	}
	err = dao.Group.DB.UploadGroupAvatar(c, db.CreateFileParams{
		FileName:   "groupAvatar",
		FileType:   "",
		FileSize:   0,
		Key:        key,
		Url:        url,
		RelationID: sql.NullInt64{Int64: relationID, Valid: true},
		AccountID:  sql.NullInt64{},
	})
	if err != nil {
		global.Logger.Error(err.Error())
		return result, errcode.ErrServer
	}
	if file == nil {
		return reply.UploadAvatar{Url: global.PbSettings.Rule.DefaultAvatarURL}, nil
	}
	return reply.UploadAvatar{Url: url}, nil
}
func (file) GetFileDetailsByID(c *gin.Context, fileID int64) (reply.File, errcode.Err) {
	result, err := dao.Group.DB.GetFileDetailsByID(c, fileID)
	if err != nil {
		return reply.File{}, errcode.ErrServer
	}
	return reply.File{
		FileID:    result.ID,
		FileName:  result.FileName,
		FileType:  string(result.FileType),
		FileSize:  result.FileSize,
		Url:       result.Url,
		AccountID: result.AccountID.Int64,
		CreateAt:  result.CreateAt,
	}, nil
}
