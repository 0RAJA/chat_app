package upload

import (
	"errors"
	"github.com/0RAJA/chat_app/src/global"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"go.uber.org/zap"
	"mime/multipart"
	"time"
)

var (
	Endpoint        = global.PvSettings.AliyunOSS.Endpoint
	AccessKeyId     = global.PvSettings.AliyunOSS.AccessKeyId
	AccessKeySecret = global.PvSettings.AliyunOSS.AccessKeySecret
	BucketName      = global.PvSettings.AliyunOSS.BucketName
	BucketUrl       = global.PvSettings.AliyunOSS.BucketUrl
	BasePath        = global.PvSettings.AliyunOSS.BasePath
)

type AliyunOSS struct{}

// UploadFile 上传文件
// 返回 访问地址，文件key，error
func (*AliyunOSS) UploadFile(file *multipart.FileHeader) (string, string, error) {
	//bucket, err := NewBucket()
	//if err != nil {
	//	global.Logger.Error("function AliyunOSS.NewBucket() Failed", zap.Any("err", err.Error()))
	//	return "", "", errors.New("function AliyunOSS.NewBucket() Failed, err:" + err.Error())
	//}

	// 读取本地文件。
	f, openError := file.Open()
	if openError != nil {
		global.Logger.Error("function file.Open() Failed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Failed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭
	// 上传阿里云路径 文件名格式 自己可以改 建议保证唯一性
	// yunFileTmpPath := filepath.Join("uploads", time.Now().Format("2006-01-02")) + "/" + file.Filename
	yunFileTmpPath := BasePath + time.Now().Format("2006-01-02 15:04:05.999999999") + "/" + file.Filename

	//上传文件流。
	//err = bucket.PutObject(yunFileTmpPath, f)
	//if err != nil {
	//	global.Logger.Error("function formUploader.Put() Failed", zap.Any("err", err.Error()))
	//	return "", "", errors.New("function formUploader.Put() Failed, err:" + err.Error())
	//}

	return BucketUrl + "/" + yunFileTmpPath, yunFileTmpPath, nil
}

// DeleteFile 删除文件
// 通过key删除对应文件
func (*AliyunOSS) DeleteFile(key string) error {
	bucket, err := NewBucket()
	if err != nil {
		global.Logger.Error("function AliyunOSS.NewBucket() Failed", zap.Any("err", err.Error()))
		return errors.New("function AliyunOSS.NewBucket() Failed, err:" + err.Error())
	}

	// 删除单个文件。objectName表示删除OSS文件时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	// 如需删除文件夹，请将objectName设置为对应的文件夹名称。如果文件夹非空，则需要将文件夹下的所有object删除后才能删除该文件夹。
	err = bucket.DeleteObject(key)
	if err != nil {
		global.Logger.Error("function bucketManager.Delete() Filed", zap.Any("err", err.Error()))
		return errors.New("function bucketManager.Delete() Filed, err:" + err.Error())
	}

	return nil
}

func NewBucket() (*oss.Bucket, error) {
	// 创建OSSClient实例。
	client, err := oss.New(Endpoint, AccessKeyId, AccessKeySecret)
	if err != nil {
		return nil, err
	}

	// 获取存储空间。
	bucket, err := client.Bucket(BucketName)
	if err != nil {
		return nil, err
	}

	return bucket, nil
}
