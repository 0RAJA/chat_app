package upload

import (
	"errors"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"mime/multipart"
	"path"
	"time"
)
type Config struct {
	BucketUrl       string
	BasePath        string
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
}

type OSS struct {
	Config Config
}

var ErrFileOpen = errors.New("文件打开失败")

func (o *OSS) UploadFile(file *multipart.FileHeader) (string, string, error) {
	bucket, err := o.newBucket()
	if err != nil {
		return "", "", errors.New("function OSS.NewBucket() Failed, err:" + err.Error())
	}

	// 读取本地文件。
	f, openError := file.Open()
	if openError != nil {
		return "", "", ErrFileOpen
	}
	defer f.Close() // 创建文件 defer 关闭
	// 上传阿里云路径 文件名格式 自己可以改 建议保证唯一性
	yunFileTmpPath := o.Config.BasePath + time.Now().Format("2006-01-02-15:04:05.99") + path.Ext(file.Filename)

	// 上传文件流。
	err = bucket.PutObject(yunFileTmpPath, f)
	if err != nil {
		return "", "", errors.New("function formUploader.Put() Failed, err:" + err.Error())
	}

	return o.Config.BucketUrl + "/" + yunFileTmpPath, yunFileTmpPath, nil
}

// DeleteFile 删除文件
// 通过key删除对应文件
func (o *OSS) DeleteFile(key ...string) (oss.DeleteObjectsResult, error) {
	if len(key) == 0 {
		return oss.DeleteObjectsResult{}, nil
	}
	bucket, err := o.newBucket()
	if err != nil {
		return oss.DeleteObjectsResult{}, errors.New("function OSS.NewBucket() Failed, err:" + err.Error())
	}
	// 删除多个文件。objectName表示删除OSS文件时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	// 如需删除文件夹，请将objectName设置为对应的文件夹名称。如果文件夹非空，则需要将文件夹下的所有object删除后才能删除该文件夹。
	result, err := bucket.DeleteObjects(key)
	if err != nil {
		return oss.DeleteObjectsResult{}, errors.New("function bucketManager.Delete() Filed, err:" + err.Error())
	}

	return result, nil
}

func (o *OSS) newBucket() (*oss.Bucket, error) {
	// 创建OSSClient实例。
	client, err := oss.New(o.Config.Endpoint, o.Config.AccessKeyId, o.Config.AccessKeySecret)
	if err != nil {
		return nil, err
	}

	// 获取存储空间。
	bucket, err := client.Bucket(o.Config.BucketName)
	if err != nil {
		return nil, err
	}

	return bucket, nil
}
