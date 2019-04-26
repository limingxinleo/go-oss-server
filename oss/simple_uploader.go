package oss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/google/uuid"
	"mime/multipart"
	"path"
)

func SimpleUpload(config *Config, bucketName string, fileHeader *multipart.FileHeader, object string) (string, error) {
	file, _ := fileHeader.Open()
	defer file.Close()

	// 创建OSSClient实例。
	client, err := oss.New(config.EndPoint, config.AccessKeyId, config.AccessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	if object == "" {
		uuid := uuid.New()
		ext := path.Ext(fileHeader.Filename)
		object = uuid.String() + ext
	}

	// 上传Byte数组。
	err = bucket.PutObject(object, file)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	return object, nil
}
