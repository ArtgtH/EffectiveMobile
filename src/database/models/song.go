package models

import (
	"gorm.io/gorm"
)

type Song struct {
	gorm.Model
	Title   string
	GroupID uint
}

type SongRequest struct {
	Title     string `json:"title"`
	GroupName string `json:"group_name"`
}

type SongResponse struct {
	Id        uint   `json:"id"`
	Title     string `json:"title"`
	GroupName string `json:"group_name"`
}
