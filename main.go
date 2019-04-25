package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)
import "os"
import "github.com/limingxinleo/go-oss-server/oss"
import "log"

func main() {
	godotenv.Load()

	config := new(oss.Config)
	config.EndPoint = os.Getenv("END_POINT")
	config.AccessKeyId = os.Getenv("ACCESS_KEY_ID")
	config.AccessKeySecret = os.Getenv("ACCESS_KEY_SECRET")

	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		bucket := c.DefaultQuery("bucket", "public")
		file, _ := c.FormFile("file")

		result, err := oss.SimpleUpload(config, bucket, file)
		if err != nil {
			log.Println("Error:", err)
			c.JSON(200, gin.H{
				"code":    500,
				"message": "文件上传失败",
			})
		} else {
			c.JSON(200, gin.H{
				"code": 0,
				"data": result,
			})
		}
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
