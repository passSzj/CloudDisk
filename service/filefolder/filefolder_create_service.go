package filefolder

import (
	"go-cloud-disk/model"
	"go-cloud-disk/serializer"
	"go-cloud-disk/utils/logger"
)

type FileFolderCreateService struct {
	ParentFolderID string `json:"parent" form:"parent" binding:"required"`
	FileFolderName string `json:"name" form:"name" binding:"required"`
}

// CreateFileFolder create filefolder to user database
func (service *FileFolderCreateService) CreateFileFolder(userId string) serializer.Response {
	// check if user match
	var fileFolder model.FileFolder
	var err error
	if err = model.DB.Where("uuid = ?", service.ParentFolderID).Find(&fileFolder).Error; err != nil {
		logger.Log().Error("[FileFolderCreateService.CreateFileFolder] Fail to get filefolder info: ", err)
		return serializer.DBErr("", err)
	}
	if fileFolder.OwnerID != userId {
		return serializer.NotAuthErr("")
	}

	// insert filefolder to database
	createFilerFolder := model.FileFolder{
		FileFolderName: service.FileFolderName,
		ParentFolderID: service.ParentFolderID,
		FileStoreID:    fileFolder.FileStoreID,
		OwnerID:        userId,
		Size:           0,
	}

	if err = model.DB.Create(&createFilerFolder).Error; err != nil {
		logger.Log().Error("[FileFolderCreateService.CreateFileFolder] Fail to create filefolder: ", err)
		return serializer.DBErr("", err)
	}
	return serializer.Success(serializer.BuildFileFolder(createFilerFolder))
}
