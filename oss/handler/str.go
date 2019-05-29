package handler

import (
	aoss "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/limingxinleo/go-oss-server/oss"
	"github.com/limingxinleo/go-oss-server/oss/option"
	"log"
	"strings"
)

type Str struct {
}

type StrBody struct {
	Content string        `form:"content" json:"content" xml:"content"  binding:"required"`
	Header  option.Header `form:"header" json:"header" xml:"header"`
}

func (s Str) Handle(ctx *gin.Context, config *oss.Config) (string, error) {
	var body StrBody

	bucket := ctx.Param("bucket")
	object := ctx.Query("object")
	if err := ctx.ShouldBindJSON(&body); err != nil {
		return "", err
	}

	result, err := s.putObject(config, bucket, body, object)

	return result, err
}

func (s Str) putObject(config *oss.Config, bucketName string, body StrBody, object string) (string, error) {

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

	options := []aoss.Option{}
	if body.Header.ContentType != "" {
		options = append(options, aoss.ContentType(body.Header.ContentType))
	}

	// 上传字符串。
	err = bucket.PutObject(object, strings.NewReader(body.Content), options...)
	if err != nil {
		log.Println("Error:", err)
		return "", nil
	}

	return object, nil
}
