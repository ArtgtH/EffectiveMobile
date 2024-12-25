package handlers

import (
	"EffectiveMobile/src/database"
	"EffectiveMobile/src/database/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateGroup(c *gin.Context) {
	db := database.DB
	group := models.GroupDTO{}

	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newGroup := models.Group{
		Title: group.Title,
		Songs: []models.Song{},
	}

	if err := db.Create(&newGroup).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": newGroup})
}

func GetGroup(c *gin.Context) {
	db := database.DB

	group := new(models.Group)
	if err := db.Where("id = ?", c.Param("id")).First(&group).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": group})
}

func GetGroups(c *gin.Context) {
	db := database.DB
	var groups []models.Group

	if err := db.Find(&groups).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": groups})
}
