package handler

import (
	"fmt"
	aoss "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"github.com/limingxinleo/go-oss-server/oss"
	"log"
)

type Deleter struct {
}

type ObjectArray struct {
	Objects []string        `form:"objects" json:"objects" xml:"objects"  binding:"required"`
}

func (s Deleter) Handle(ctx *gin.Context, config *oss.Config) (string, error) {
	var body ObjectArray
	bucketName := ctx.Param("bucket")
	if err := ctx.ShouldBindJSON(&body); err != nil {
		return "", err
	}

	client, err := aoss.New(config.EndPoint, config.AccessKeyId, config.AccessKeySecret)
	if err != nil {
		log.Println("Error:", err)
		return "", nil
	}

	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	bucket.DeleteObjects(body.Objects)

	return "", err
}