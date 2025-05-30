package share

import (
	"go-cloud-disk/model"
	"go-cloud-disk/serializer"
	"go-cloud-disk/utils/logger"
)

type ShareGetAllService struct {
}

// GetAllShare get user's all share
func (service *ShareGetAllService) GetAllShare(userId string) serializer.Response {
	// get shares from database
	var shares []model.Share
	if err := model.DB.Where("owner = ?", userId).Find(&shares).Error; err != nil {
		logger.Log().Error("[ShareGetAllService.GetAllShare] Fail to get share info: ", err)
		return serializer.DBErr("", err)
	}

	return serializer.Success(serializer.BuildShares(shares))
}
