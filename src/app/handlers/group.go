package handlers

import (
	"EffectiveMobile/src/app/services"
	"EffectiveMobile/src/database/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GroupHandler struct {
	groupService *services.GroupService
}

func NewGroupHandler(groupService *services.GroupService) *GroupHandler {
	return &GroupHandler{groupService: groupService}
}

// CreateGroup godoc
// @Tags Groups
// @Summary Create group
// @Accept json
// @Produce json
// @Param group body models.GroupRequest true "Group DTO"
// @Success 201 {object} models.GroupResponse
// @Failure 400 "Bad Request"
// @Failure 500 "Internal Server Error"
// @Router /api/groups [post]
func (h *GroupHandler) CreateGroup(c *gin.Context) {
	group := models.GroupRequest{}

	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resGroup, err := h.groupService.CreateGroup(group)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": resGroup})
}

// GetGroup godoc
// @Tags Groups
// @Summary Return group by given ID
// @Produce json
// @Param id path integer true "Group ID"
// @Success 200 {object} models.GroupResponse
// @Router /api/groups/{id} [get]
func (h *GroupHandler) GetGroup(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resGroup, err := h.groupService.GetGroup(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": resGroup})
}

// GetGroups godoc
// @Tags Groups
// @Summary Return all groups
// @Produce json
// @Success 200 {object} models.GroupResponse
// @Router /api/groups [get]
func (h *GroupHandler) GetGroups(c *gin.Context) {
	resGroups, err := h.groupService.GetGroups()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": resGroups})
}
