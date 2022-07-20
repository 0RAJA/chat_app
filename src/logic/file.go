package logic

import (
	"fmt"
	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/0RAJA/chat_app/src/dao"
	db "github.com/0RAJA/chat_app/src/dao/postgres/sqlc"
	"github.com/0RAJA/chat_app/src/model/reply"
	"github.com/0RAJA/chat_app/src/model/request"
	"github.com/0RAJA/chat_app/src/myerr"
	"github.com/0RAJA/chat_app/src/upload"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type file struct {

}

func (file) PublishFile(c *gin.Context,params request.PublishFile,filetype string,file *multipart.FileHeader) (reply.PublishFile,errcode.Err) {
	oss := upload.NewOSS()
	url,key,err := oss.UploadFile(file)
	result := reply.PublishFile{}
	if err != nil {
		return result,myerr.OSSFiledStore
	}
	r,err := dao.Group.DB.CreateFile(c,&db.CreateFileParams{
		FileName:   params.File.Filename,
		FileType:   filetype,
		FileSize:   params.File.Size,
		Key:        key,
		Url:        url,
		RelationID: params.RelationID,
		AccountID:  params.AccountID,
	})
	if err != nil {
		fmt.Println(err)
		return result,errcode.ErrServer
	}
	result = reply.PublishFile{
		FileType:filetype,
		FileSize: params.File.Size,
		Url:      "url",
		CreateAt: r.CreateAt,
	}
	return result,nil
}