package filefolder

import (
	"go-cloud-disk/model"
	"go-cloud-disk/serializer"
	"go-cloud-disk/utils/logger"
)

type FileFolderGetAllFileService struct {
}

// GetAllFile get all file in user filefolder
func (service *FileFolderGetAllFileService) GetAllFile(userId string, fileFolderID string) serializer.Response {
	// check if user match
	var fileFolder model.FileFolder
	if err := model.DB.Where("uuid = ? and owner_id = ?", fileFolderID, userId).Find(&fileFolder).Error; err != nil {
		logger.Log().Error("[FileFolderGetAllFileService.GetAllFile] Fail to get Filefoldedr: ", err)
		return serializer.DBErr("", err)
	}

	var files []model.File
	if err := model.DB.Where("parent_folder_id = ?", fileFolderID).Find(&files).Error; err != nil {
		logger.Log().Error("[FileFolderGetAllFileService.GetAllFile] Fail to get Files: ", err)
		return serializer.DBErr("", err)
	}
	return serializer.Success(serializer.BuildFiles(files))
}
