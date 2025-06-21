package authrouter

import (
	"github.com/anle/codebase/global"
	"github.com/anle/codebase/internal/wire"
	"github.com/gin-gonic/gin"
)

type UploadRouter struct{}

func (ar *UploadRouter) InitUploadRouter(router *gin.RouterGroup) {
	uploadHandler, _ := wire.InitUploadRouterHandler(global.Mdb)

	AuthRouterPublic := router.Group("/upload")

	{
		AuthRouterPublic.POST("/", uploadHandler.Upload)
	}
}

func NewUploadRouter() *UploadRouter {
	return &UploadRouter{}
}
