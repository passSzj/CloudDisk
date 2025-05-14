package api

import (
	"github.com/gin-gonic/gin"
	"go-cloud-disk/serializer"
	"go-cloud-disk/service/rank"
)

func GetDailyRank(c *gin.Context) {
	var service rank.GetDailyRankService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, serializer.ErrorResponse(err))
		return
	}

	res := service.GetDailyRank()
	c.JSON(200, res)
}
