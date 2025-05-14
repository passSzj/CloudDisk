package user

import (
	"go-cloud-disk/model"
	"go-cloud-disk/serializer"
	"go-cloud-disk/utils/logger"
)

type UserInfoService struct {
}

// GetUserInfo get user info by userid
func (service *UserInfoService) GetUserInfo(userid string) serializer.Response {
	var user model.User

	err := model.DB.Model(&model.User{}).Where("uuid = ?", userid).First(&user).Error
	if err != nil {
		logger.Log().Error("[UserInfoService.GetUserInfo] Fail to find user")
		return serializer.ParamsErr("NotFound", err)
	}

	return serializer.Success(serializer.BuildUser(user))
}
