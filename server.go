package main

import "github.com/gin-gonic/gin"
import "os"
import "github.com/limingxinleo/go-oss-server/oss"
import "github.com/joho/godotenv"
import "log"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := new(oss.Config)
	config.EndPoint = os.Getenv("END_POINT")
	config.AccessKeyId = os.Getenv("ACCESS_KEY_ID")
	config.AccessKeySecret = os.Getenv("ACCESS_KEY_SECRET")

	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		bucket := c.DefaultQuery("bucket", "public")
		object := c.Query("object")
		file, _ := c.FormFile("file")

		c.JSON(200, gin.H{
			"message": gin.H{
				"bucket":   bucket,
				"object":   object,
				"fileName": file.Filename,
			},
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
