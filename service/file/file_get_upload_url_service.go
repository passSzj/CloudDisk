package file

import (
	"github.com/google/uuid"
	"go-cloud-disk/disk"
	"go-cloud-disk/serializer"
	"go-cloud-disk/utils/logger"
)

type GetUploadURLService struct {
	FileType string `form:"filetype" json:"filetype" binding:"required,min=1"`
}

type getUploadURLResponse struct {
	Url      string `json:"url"`
	FileUuid string `json:"file_uuid"`
}

func (service *GetUploadURLService) GetUploadURL(fileowner string) serializer.Response {
	fileID := uuid.New().String()
	fileName := fileID + "." + service.FileType
	url, err := disk.BaseCloudDisk.GetUploadPresignedURL(fileowner, "", fileName)
	if err != nil {
		logger.Log().Error("[GetUploadURLService.GetUploadURL] Fail to get upload url: ", err)
		return serializer.InternalErr("", err)
	}

	return serializer.Success(getUploadURLResponse{
		Url:      url,
		FileUuid: fileID,
	})
}
