package models

import (
	"gorm.io/gorm"
)

type Items struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	SKU         string `gorm:"not null"`
	UserID      uint   `gorm:"not null"`
	ItemTypesID uint   `gorm:"not null"`
	CompanyID   uint   `gorm:"not null"`
	SizeID      uint   `gorm:"not null"`
}
