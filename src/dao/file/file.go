package file

import "github.com/0RAJA/chat_app/src/dao/file/upload"

func Init(config upload.Config) *upload.OSS {
	return &upload.OSS{Config: config}
}