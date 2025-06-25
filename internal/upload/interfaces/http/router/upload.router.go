package authrouter

import (
	"github.com/anle/codebase/global"
	"github.com/anle/codebase/internal/wire"
	"github.com/gin-gonic/gin"
)

type UploadRouter struct{}

func (ar *UploadRouter) InitUploadRouter(router *gin.RouterGroup) {
	// Initialize the cron job for deleting expired files
	DeleteCronHandler, _ := wire.InitDeleteExpiredFileService(global.Mdb, global.MinioClient)
	DeleteCronHandler.RegisterDeleteCron()

	uploadHandler, _ := wire.InitUploadRouterHandler(global.Mdb, global.MinioClient)
	UploadRouterPublic := router.Group("/upload")

	{
		UploadRouterPublic.POST("/", uploadHandler.Upload)
	}
}

func NewUploadRouter() *UploadRouter {
	return &UploadRouter{}
}
