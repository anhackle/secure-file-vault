package authrouter

import (
	"github.com/anle/codebase/global"
	"github.com/anle/codebase/internal/wire"
	"github.com/gin-gonic/gin"
)

type ShareRouter struct{}

func (ar *ShareRouter) InitShareRouter(router *gin.RouterGroup) {
	shareHandler, _ := wire.InitShareRouterHandler(global.Mdb, global.MinioClient, global.Rdb)
	ShareRouterPublic := router.Group("/share")

	{
		ShareRouterPublic.POST("/", shareHandler.Share)
		ShareRouterPublic.GET("/:id", shareHandler.ProcessSharedURL)
	}
}

func NewShareRouter() *ShareRouter {
	return &ShareRouter{}
}
