package api

import (
	"github.com/gin-gonic/gin"
	"go-cloud-disk/serializer"
)

func Ping(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "Pong",
	})
}
