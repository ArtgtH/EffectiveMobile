package handlers

import (
	"EffectiveMobile/src/app/services"
	"EffectiveMobile/src/database/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type SongHandler struct {
	songService  *services.SongService
	groupService *services.GroupService
}

func NewSongHandler(songService *services.SongService, groupService *services.GroupService) *SongHandler {
	return &SongHandler{songService: songService, groupService: groupService}
}

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
func (h *SongHandler) CreateSong(c *gin.Context) {
	song := models.SongRequest{}
	if err := c.ShouldBindJSON(&song); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resSong, err := h.songService.CreateSong(song)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
func (h *SongHandler) GetSong(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resSong, err := h.songService.GetSong(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": resSong})
}

// GetSongs godoc
// @Tags Songs
// @Summary Return all songs
// @Produce json
// @Success 200 {object} models.SongResponse
// @Router /api/songs [get]
func (h *SongHandler) GetSongs(c *gin.Context) {
	resSongs, err := h.songService.GetSongs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": resSongs})
}
