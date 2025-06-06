package server

import (
	"github.com/gin-gonic/gin"
	"go-cloud-disk/api"
	"go-cloud-disk/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.MaxMultipartMemory = 8 << 20 // set upload speed
	r.Use(middleware.Cors())
	r.GET("ping", api.Ping)

	v1 := r.Group("/api/v1")
	{
		v1.POST("user/login", api.UserLogin)
		v1.POST("user/register", api.UserRegiser)
		v1.POST("user/email", api.ConfirmUserEmail)

		v1.GET("share/:shareId", api.GetShareInfo)

		auth := v1.Group("")
		auth.Use(middleware.JWTAuth(), middleware.CasbinAuth())
		{
			auth.GET("user/:id", api.UserInfo)
			auth.GET("user", api.UserMyInfo)
			auth.PUT("user", api.UpdateUserInfo)

			auth.GET("file/:fileid", api.GetDownloadURL)
			auth.POST("file", api.UploadFile)
			auth.PUT("file", api.UpdateFile)
			auth.DELETE("file/:fileid", api.DeleteFile)

			auth.GET("filefolder/:filefolderid/file", api.GetFilefolderAllFile)
			auth.GET("filefolder/:filefolderid/filefolder", api.GetFilefolderAllFilefolder)
			auth.POST("filefolder", api.CreateFileFolder)
			auth.PUT("filefolder", api.UpdateFileFolder)
			auth.DELETE("filefolder/:filefolderid", api.DeleteFileFolder)

			auth.GET("filestore/:filestoreId", api.GetFileStoreInfo)

			auth.GET("share", api.GetUserAllShare)
			auth.POST("share", api.CreateShare)
			auth.DELETE("share/:shareId", api.DeleteShare)
			auth.POST("share/file", api.ShareSaveFile)

			auth.GET("rank/day", api.GetDailyRank)

			admin := auth.Group("admin")
			admin.Use(middleware.AdminAuth())
			{
				admin.POST("user", api.SearchUser)
				admin.PUT("user", api.UpdateUserAuth)

				admin.POST("share", api.SearchShare)
				admin.DELETE("share/:shareId", api.AdminDeleteShare)

				admin.DELETE("file/:fileId", api.AdminDeleteFile)

				admin.GET("filestore/:userId", api.AdminGetFileStoreInfo)
				admin.PUT("filestore", api.UserFileStoreUpdate)
			}
		}
	}

	return r
}
