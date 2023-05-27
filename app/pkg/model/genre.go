package model

import (
	"gorm.io/gorm"
)

type Genre struct {
	gorm.Model
	Name string `json:"name"`
}

func NewGenre(name string) *Genre {
	return &Genre{
		Name: name,
	}
}

func CreateGenre(db *gorm.DB, genre *Genre) (*Genre, error) {
	result := db.Create(&genre)
	return genre, result.Error
}
