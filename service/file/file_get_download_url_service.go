package file

import (
	"go-cloud-disk/disk"
	"go-cloud-disk/model"
	"go-cloud-disk/serializer"
	"go-cloud-disk/utils/logger"
)

type FileGetDownloadURLService struct {
}

type fileGetDownloadURLResponse struct {
	Url string `json:"dowload_url"`
}

func (service *FileGetDownloadURLService) GetDownloadURL(userId string, fileid string) serializer.Response {
	var file model.File
	if err := model.DB.Where("uuid = ?", fileid).Find(&file).Error; err != nil {
		logger.Log().Error("[fileGetDownloadURLResponse.GetDownloadURL] Fail to find user file: ", err)
		return serializer.DBErr("", err)
	}

	if userId != file.Owner {
		return serializer.NotAuthErr("")
	}

	fileName := file.FileUuid + "." + file.FilePostfix
	url, err := disk.BaseCloudDisk.GetObjectURL(file.FilePath, "", fileName)
	if err != nil {
		logger.Log().Error("[FileGetDownloadURLService.GetDownloadURL] Fail to get download URL: ", err)
		return serializer.InternalErr("", err)
	}
	return serializer.Success(fileGetDownloadURLResponse{
		Url: url,
	})
}
