package core

import "github.com/gin-gonic/gin"

type HttpResponse struct {
	ctx *gin.Context
}

func NewHttpResponse(ctx *gin.Context) *HttpResponse {
	res := new(HttpResponse)
	res.ctx = ctx
	return res
}

func (http HttpResponse) Success(data interface{}) {
	http.ctx.JSON(200, gin.H{
		"code": 0,
		"data": data,
	})
}

func (http HttpResponse) Failed(code int, message string) {
	http.ctx.JSON(200, gin.H{
		"code":    code,
		"message": message,
	})
}
