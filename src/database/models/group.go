package models

import (
	"gorm.io/gorm"
	"math/rand"
)

type Group struct {
	gorm.Model
	Title string
	Songs []Song
	Joke  int
}

type GroupDTO struct {
	Title string `json:"title"`
}

func (group *Group) BeforeCreate(_ *gorm.DB) error {
	group.Joke = rand.Intn(100)
	return nil
}
