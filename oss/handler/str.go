package handler

import (
	aoss "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/limingxinleo/go-oss-server/oss"
	"log"
	"strings"
)

type Str struct {
}

func (s Str) Handle(ctx *gin.Context, config *oss.Config) (string, error) {
	bucket := ctx.Param("bucket")
	object := ctx.Query("object")
	data := ctx.PostForm("str")

	result, err := s.putObject(config, bucket, data, object)

	return result, err
}

func (s Str) putObject(config *oss.Config, bucketName string, data string, object string) (string, error) {

	// 创建OSSClient实例。
	client, err := aoss.New(config.EndPoint, config.AccessKeyId, config.AccessKeySecret)
	if err != nil {
		log.Println("Error:", err)
		return "", nil
	}

	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		log.Println("Error:", err)
		return "", nil
	}

	if object == "" {
		uuid := uuid.New()
		object = uuid.String()
	}

	// 上传字符串。
	err = bucket.PutObject(object, strings.NewReader(data))
	if err != nil {
		log.Println("Error:", err)
		return "", nil
	}

	return object, nil
}
