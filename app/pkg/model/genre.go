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

func GetGenreById(db *gorm.DB, ID int) (*Genre, error) {
	genre := Genre{}
	result := db.First(&genre, ID)
	return &genre, result.Error
}

func (genre *Genre) Update(db *gorm.DB, param map[string]interface{}) (*Genre, error) {
	result := db.Model(&genre).Updates(param)
	return genre, result.Error
}

func GetGenres(db *gorm.DB) ([]*Genre, error) {
	genres := []*Genre{}
	result := db.Find(&genres)
	return genres, result.Error
}
