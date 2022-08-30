package setting

import (
	"github.com/0RAJA/Rutils/pkg/upload/oss/aliyun"
	"github.com/0RAJA/chat_app/src/global"
)

type oss struct {
}

func (oss) Init() {
	global.OSS = aliyun.Init(aliyun.Config{
		BucketUrl:       global.PvSettings.AliyunOSS.BucketUrl,
		BasePath:        global.PvSettings.AliyunOSS.BasePath,
		Endpoint:        global.PvSettings.AliyunOSS.Endpoint,
		AccessKeyId:     global.PvSettings.AliyunOSS.AccessKeyId,
		AccessKeySecret: global.PvSettings.AliyunOSS.AccessKeySecret,
		BucketName:      global.PvSettings.AliyunOSS.BucketName,
	})
}
