package upload

import (
	"mime/multipart"
)

// OSS 对象存储接口
type OSS interface {
	UploadFile(file *multipart.FileHeader) (string, string, error)
	DeleteFile(key string) error
}

// NewOSS OSS的实例化方法
func NewOSS() OSS {
	return &AliyunOSS{}
}
