package file

import (
	"github.com/0RAJA/chat_app/src/dao/file/upload"
	"github.com/0RAJA/chat_app/src/global"
)

func Init(config upload.Config) *upload.OSS {
	o := upload.Config{
		BucketUrl:       global.PvSettings.AliyunOSS.BucketUrl,
		BasePath:        global.PvSettings.AliyunOSS.BasePath,
		Endpoint:        global.PvSettings.AliyunOSS.Endpoint,
		AccessKeyId:     global.PvSettings.AliyunOSS.AccessKeyId,
		AccessKeySecret: global.PvSettings.AliyunOSS.AccessKeySecret,
		BucketName:      global.PvSettings.AliyunOSS.BucketName,
	}
	return &upload.OSS{Config: o}
}
