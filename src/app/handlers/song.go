package handlers

import (
	"EffectiveMobile/src/database"
	"EffectiveMobile/src/database/methods"
	"EffectiveMobile/src/database/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateSong godoc
// @Tags Songs
// @Summary Create song
// @Accept json
// @Produce json
// @Param song body models.SongRequest true "Song DTO"
// @Success 201 {object} models.SongResponse
// @Failure 400 "Bad Request"
// @Failure 500 "Internal Server Error"
// @Router /api/songs [post]
func CreateSong(c *gin.Context) {
	db := database.DB
	song := models.SongRequest{}

	if err := c.ShouldBindJSON(&song); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	groupID := methods.GetGroupIDByTitle(song.GroupName)

	newSong := models.Song{
		Title:   song.Title,
		GroupID: groupID,
	}

	if err := db.Create(&newSong).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	groupTitle := methods.GetGroupTitleByID(newSong.GroupID)

	resSong := models.SongResponse{
		Id:        newSong.ID,
		Title:     newSong.Title,
		GroupName: groupTitle,
	}

	c.JSON(http.StatusCreated, gin.H{"data": resSong})
}

// GetSong godoc
// @Tags Songs
// @Summary Return song by given ID
// @Produce json
// @Param id path integer true "Song ID"
// @Success 200 {object} models.SongResponse
// @Router /api/songs/{id} [get]
func GetSong(c *gin.Context) {
	db := database.DB
	song := models.Song{}
	if err := db.First(&song, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	groupTitle := methods.GetGroupTitleByID(song.GroupID)
	resSong := models.SongResponse{
		Id:        song.ID,
		Title:     song.Title,
		GroupName: groupTitle,
	}
	c.JSON(http.StatusOK, gin.H{"data": resSong})
}

// GetSongs godoc
// @Tags Songs
// @Summary Return all songs
// @Produce json
// @Success 200 {object} models.SongResponse
// @Router /api/songs [get]
func GetSongs(c *gin.Context) {
	db := database.DB
	var songs []models.Song
	if err := db.Find(&songs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	resSongs := make([]models.SongResponse, len(songs))
	for idx, song := range songs {
		groupTitle := methods.GetGroupTitleByID(song.GroupID)
		resSongs[idx] = models.SongResponse{
			Id:        song.ID,
			Title:     song.Title,
			GroupName: groupTitle,
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": resSongs})
}
