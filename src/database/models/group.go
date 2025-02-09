package models

import (
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	Title string
	Songs []Song
}

type GroupRequest struct {
	Title string `json:"title" validate:"required"`
	Songs []Song `json:"songs"`
}

type GroupResponse struct {
	Id    uint   `json:"id"`
	Title string `json:"title"`
	Songs []Song `json:"songs"`
}
