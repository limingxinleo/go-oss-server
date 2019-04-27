package handler

import (
	"fmt"
	aoss "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/limingxinleo/go-oss-server/oss"
	"mime/multipart"
	"path"
)

type SimpleUploader struct {
}

func (s SimpleUploader) Handle(ctx *gin.Context, config *oss.Config) (string, error) {
	bucket := ctx.Param("bucket")
	object := ctx.Query("object")
	file, _ := ctx.FormFile("file")

	result, err := s.SimpleUpload(config, bucket, file, object)

	return result, err
}

func (s SimpleUploader) SimpleUpload(config *oss.Config, bucketName string, fileHeader *multipart.FileHeader, object string) (string, error) {
	file, _ := fileHeader.Open()
	defer file.Close()

	// 创建OSSClient实例。
	client, err := aoss.New(config.EndPoint, config.AccessKeyId, config.AccessKeySecret)
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
