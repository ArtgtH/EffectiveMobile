package methods

import (
	"EffectiveMobile/src/database"
	"EffectiveMobile/src/database/models"
)

func GetGroupIDByTitle(title string) uint {
	db := database.DB
	group := models.Group{}

	db.Find(&group, "title = ?", title)

	return group.ID
}

func GetGroupTitleByID(id uint) string {
	db := database.DB
	group := models.Group{}

	db.Find(&group, "id = ?", id)

	return group.Title
}
