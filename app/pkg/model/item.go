package model

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name         string `json:"mame"`
	Introduction string `json:"introduction"`
	Price        int    `json:"price"`
	IsActive     bool   `json:"is_active"`
}

func GetItems(db *gorm.DB) ([]*Item, error) {
	items := []*Item{}
	result := db.Find(&items)
	return items, result.Error
}

func CreateItem(db *gorm.DB, item *Item) (*Item, error) {
	result := db.Create(&item)
	return item, result.Error
}

func NewItem(name string, introduction string, price int) *Item {
	return &Item{
		Name:         name,
		Introduction: introduction,
		Price:        price,
	}
}

func GetItemById(db *gorm.DB, ID int) (*Item, error) {
	item := Item{}
	result := db.First(&item, ID)
	return &item, result.Error
}
