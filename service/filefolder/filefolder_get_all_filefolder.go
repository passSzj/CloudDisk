package filefolder

import (
	"go-cloud-disk/model"
	"go-cloud-disk/serializer"
	"go-cloud-disk/utils/logger"
)

type FileFolderGetAllFileFolderService struct {
}

// GetAllFileFolder get user all filefolder form filefolder
func (service *FileFolderGetAllFileFolderService) GetAllFileFolder(userId string, fileFolderID string) serializer.Response {
	var filefolder []model.FileFolder
	if err := model.DB.Where("parent_folder_id = ? and owner_id = ?", fileFolderID, userId).Find(&filefolder).Error; err != nil {
		logger.Log().Error("[FileFolderGetAllFileFolderService.GetAllFileFolder] Fail to find filefolder: ", err)
		return serializer.DBErr("", err)
	}
	return serializer.Success(serializer.BuildFileFolders(filefolder))
}
