package filefolder

import (
	"go-cloud-disk/model"
	"go-cloud-disk/serializer"
	"go-cloud-disk/utils/logger"
)

type DeleteFileFolderService struct {
}

// DeleteFileFolder tmp delete filefolder, this func will update when add size model
func (service *DeleteFileFolderService) DeleteFileFolder(userId string, fileFolderId string) serializer.Response {
	// check if user auth match this filefolder
	var fileFolder model.FileFolder
	var err error
	if err := model.DB.Where("uuid = ?", fileFolderId).Find(&fileFolder).Error; err != nil {
		logger.Log().Error("[DeleteFileFolderService.DeleteFileFolder] Fail to find filefolder info: ", err)
		return serializer.DBErr("", err)
	}
	if fileFolder.OwnerID != userId {
		return serializer.NotAuthErr("")
	}

	// delete filefolder form list and protect filefolder from duplicate delete
	if fileFolder.ParentFolderID == "root" || fileFolder.ParentFolderID == "" {
		return serializer.ParamsErr("CanDeleteRoot", nil)
	}
	t := model.DB.Begin()
	defer func() {
		if err != nil {
			t.Rollback()
		} else {
			t.Commit()
		}
	}()

	// delete filefolder and file that in DeleteFileFolder
	deleteFileFolderIDs := []string{}
	deleteFileFolderIDs = append(deleteFileFolderIDs, fileFolderId)
	for len(deleteFileFolderIDs) > 0 {
		deleteFileFolders := []model.FileFolder{}
		deleteIDs := []string{}
		// get filefolder that in deleteFileFolder
		if err := t.Select("uuid").Where("parent_folder_id in (?)", deleteFileFolderIDs).Find(&deleteFileFolders).Error; err != nil {
			logger.Log().Error("[DeleteFileFolderService.DeleteFileFolder] Fail to find delete filefolders info: ", err)
			return serializer.DBErr("", err)
		}
		// get will delete filefolder id
		for _, filefolder := range deleteFileFolders {
			deleteIDs = append(deleteIDs, filefolder.Uuid)
		}
		if err := t.Where("uuid in (?)", deleteFileFolderIDs).Delete(&model.FileFolder{}).Error; err != nil {
			logger.Log().Error("[DeleteFileFolderService.DeleteFileFolder] Fail to delete filefolder: ", err)
			return serializer.DBErr("", err)
		}
		if err := t.Where("parent_folder_id in (?)", deleteFileFolderIDs).Delete(&model.FileFolder{}).Error; err != nil {
			logger.Log().Error("[DeleteFileFolderService.DeleteFileFolder] Fail to delete file: ", err)
			return serializer.DBErr("", err)
		}
		deleteFileFolderIDs = deleteIDs
	}

	// delete filefolder size from parent filefolder
	if fileFolder.ParentFolderID != "root" {
		var parentFileFolder model.FileFolder
		if err := t.Where("uuid = ?", fileFolder.ParentFolderID).Find(&parentFileFolder).Error; err != nil {
			logger.Log().Error("[DeleteFileFolderService.DeleteFileFolder] Fail to find filefolder info: ", err)
			return serializer.DBErr("", err)
		}
		if err := parentFileFolder.SubFileFolderSize(t, fileFolder.Size); err != nil {
			logger.Log().Error("[DeleteFileFolderService.DeleteFileFolder] Fail to update parent filefolder info: ", err)
			return serializer.DBErr("", err)
		}
	}

	// delete filefolder size from userSotre
	var userStore model.FileStore
	if err := t.Where("uuid = ? and owner_id = ?", fileFolder.FileStoreID, userId).Find(&userStore).Error; err != nil {
		logger.Log().Error("[DeleteFileFolderService.DeleteFileFolder] Fail to find filestore: ", err)
		return serializer.DBErr("", err)
	}
	userStore.AddCurrentSize(fileFolder.Size)
	if err = t.Save(&userStore).Error; err != nil {
		logger.Log().Error("[DeleteFileFolderService.DeleteFileFolder] Fail to update filestore: ", err)
		return serializer.DBErr("", err)
	}

	return serializer.Success(nil)
}
