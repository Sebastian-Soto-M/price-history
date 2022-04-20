package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Id            int
	Name          string
	Url           string
	Price         float32
	OriginalPrice float32
	DateAdded     string
}
