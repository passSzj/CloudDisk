package api

import (
	"github.com/gin-gonic/gin"
	"go-cloud-disk/serializer"
	"go-cloud-disk/service/share"
)

// CreateShare use fileid and userid to build share
func CreateShare(c *gin.Context) {
	var service share.ShareCreateService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, serializer.ErrorResponse(err))
		return
	}

	userId := c.MustGet("UserId").(string)
	res := service.CreateShare(userId)
	c.JSON(200, res)
}

// GetShareInfo get share info by share id, add view of share
func GetShareInfo(c *gin.Context) {
	var service share.ShareGetInfoService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, serializer.ErrorResponse(err))
		return
	}

	shareId := c.Param("shareId")
	res := service.GetShareInfo(shareId)
	c.JSON(200, res)
}

// GetUserAllShare get user all share info
func GetUserAllShare(c *gin.Context) {
	var service share.ShareGetAllService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, serializer.ErrorResponse(err))
		return
	}

	userId := c.MustGet("UserId").(string)
	res := service.GetAllShare(userId)
	c.JSON(200, res)
}

// DeleteShare delete share by shareid
func DeleteShare(c *gin.Context) {
	var service share.ShareDeleteService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, serializer.ErrorResponse(err))
		return
	}

	shareId := c.Param("shareId")
	userId := c.MustGet("UserId").(string)
	res := service.DeleteShare(shareId, userId)
	c.JSON(200, res)
}

// ShareSaveFile save share file to user filefolder
func ShareSaveFile(c *gin.Context) {
	var service share.ShareSaveFileService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, serializer.ErrorResponse(err))
		return
	}

	userId := c.MustGet("UserId").(string)
	res := service.ShareSaveFile(userId)
	c.JSON(200, res)
}
