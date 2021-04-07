package models

import (
	"gorm.io/gorm"
)

type Item_Images struct {
	gorm.Model
	ItemID   uint `gorm:"not null"`
	ImageURL string
}
