package app

import (
	"EffectiveMobile/src/app/handlers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Songs Swagger API
// @version 1.0
// @description Swagger API for test Golang Project Songs.
// @termsOfService http://swagger.io/terms/

// @BasePath /api/
// @host localhost:8080
// @schemes http

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		api.GET("/songs/:id", handlers.GetSong)
		api.GET("/groups/:id", handlers.GetGroup)
		api.GET("/songs", handlers.GetSongs)
		api.GET("/groups", handlers.GetGroups)
		api.POST("/groups", handlers.CreateGroup)
		api.POST("/songs", handlers.CreateSong)
	}

	return router
}
