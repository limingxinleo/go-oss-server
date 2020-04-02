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

func (s Deleter) Handle(ctx *gin.Context, config *oss.Config) (string, error) {
	bucketName := ctx.Param("bucket")
	objects := ctx.Query("objects")

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

	log.Println(bucket)
	//res,_ = bucket.DeleteObjects(objects)

	return objects, err
}