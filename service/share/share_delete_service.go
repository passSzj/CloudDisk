package share

import (
	"go-cloud-disk/model"
	"go-cloud-disk/serializer"
	"go-cloud-disk/utils/logger"
)

type ShareDeleteService struct {
}

func (service *ShareDeleteService) DeleteShare(shareId string, userId string) serializer.Response {
	// get shares from database
	var share model.Share
	if err := model.DB.Where("uuid = ? and owner = ?", shareId, userId).First(&share).Error; err != nil {
		logger.Log().Error("[ShareDeleteService.DeleteShare] Fail to get share info: ", err)
		return serializer.DBErr("", err)
	}

	// delay double delete, ensure the safe of info
	share.DeleteShareInfoInRedis()
	if err := model.DB.Delete(&share).Error; err != nil {
		logger.Log().Error("[ShareDeleteService.DeleteShare] Fail to delete share: ", err)
		return serializer.DBErr("", err)
	}
	share.DeleteShareInfoInRedis()

	return serializer.Success(nil)
}
