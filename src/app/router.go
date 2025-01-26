package app

import (
	"EffectiveMobile/src/app/handlers"
	"EffectiveMobile/src/app/services"
	"EffectiveMobile/src/database"
	"EffectiveMobile/src/database/repositories"
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
	db := database.DB

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	groupRepo := repositories.NewGroupRepository(db)
	songRepo := repositories.NewSongRepository(db)

	groupService := services.NewGroupService(groupRepo)
	songService := services.NewSongService(songRepo, groupRepo)

	groupHandler := handlers.NewGroupHandler(groupService)
	songHandler := handlers.NewSongHandler(songService, groupService)

	api := router.Group("/api")
	{
		api.GET("/groups/:id", groupHandler.GetGroup)
		api.GET("/groups", groupHandler.GetGroups)
		api.POST("/groups", groupHandler.CreateGroup)
	}
	{
		api.GET("/songs/:id", songHandler.GetSong)
		api.GET("/songs", songHandler.GetSongs)
		api.POST("/songs", songHandler.CreateSong)
	}

	return router
}
