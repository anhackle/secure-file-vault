package initialize

import (
	"github.com/anle/codebase/global"
	"github.com/anle/codebase/internal/middlewares"
	uploadrouter "github.com/anle/codebase/internal/upload/interfaces/http/router"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	MainGroup := r.Group("/v1")
	MainGroup.Use(middlewares.CORSMiddleware())

	{
		uploadRouter := uploadrouter.NewUploadRouter()
		uploadRouter.InitUploadRouter(MainGroup)

		shareRouter := uploadrouter.NewShareRouter()
		shareRouter.InitShareRouter(MainGroup)
	}

	return r
}
