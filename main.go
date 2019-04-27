package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/limingxinleo/go-oss-server/core"
	"github.com/limingxinleo/go-oss-server/oss/handler"
	"log"
	"reflect"
)
import "os"
import "github.com/limingxinleo/go-oss-server/oss"

var handlers map[string]interface{}
var config *oss.Config

func init() {
	godotenv.Load()
	config = new(oss.Config)
	config.EndPoint = os.Getenv("END_POINT")
	config.AccessKeyId = os.Getenv("ACCESS_KEY_ID")
	config.AccessKeySecret = os.Getenv("ACCESS_KEY_SECRET")

	handlers = make(map[string]interface{})
	handlers["simple_uploader"] = handler.SimpleUploader{}

	log.Println(handlers)
}

func main() {
	r := gin.Default()

	r.POST("/:handler/:bucket", func(c *gin.Context) {
		handlerName := c.Param("handler")

		response := core.NewHttpResponse(c)

		if handlers[handlerName] == nil {
			response.Failed(500, "Handler is invalid.")
			return
		}

		t := reflect.ValueOf(handlers[handlerName]).Type()
		handler := reflect.New(t).Elem().Interface().(handler.HandlerInterface)

		result, err := handler.Handle(c, config)

		if err != nil {
			log.Println("Error:", err)
			response.Failed(500, "文件上传失败")
			return
		}

		response.Success(result)
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
