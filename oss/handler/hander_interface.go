package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/limingxinleo/go-oss-server/oss"
)

type HandlerInterface interface {
	Handle(ctx *gin.Context, config *oss.Config) (string, error)
}
