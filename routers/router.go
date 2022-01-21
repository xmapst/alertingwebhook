package routers

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/xmapst/alertingwebhook/cmd"
	_ "github.com/xmapst/alertingwebhook/docs"
	"github.com/xmapst/alertingwebhook/handlers/google"
)

func Router() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	if cmd.DebugEnabled {
		pprof.Register(router)
		// swagger doc
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	googleGroup := router.Group("/google")
	{
		googleGroup.POST("/notification", google.Notification)
	}
	return router
}
