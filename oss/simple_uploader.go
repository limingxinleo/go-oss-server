package oss

import (
	"mime/multipart"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/google/uuid"
	"fmt"
)

func SimpleUpload(config *Config, bucketName string, fileHeader *multipart.FileHeader) (string, error) {
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

	uuid := uuid.New()
	object := uuid.String() + ".png"

	// 上传Byte数组。
	err = bucket.PutObject(object, file)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	return object, nil
}
