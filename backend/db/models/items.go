package models

import (
	"gorm.io/gorm"
)

type Items struct {
	gorm.Model
	Name        string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	SKU         string    `gorm:"not null"`
	UserID      uint      `gorm:"not null"`
	ItemTypeID  uint      `gorm:"not null"`
	CompanyID   uint      `gorm:"not null"`
	SizeID      uint      `gorm:"not null"`
	Users       Users     `gorm:"foreignKey:UserID;references:id"`
	ItemTypes   ItemTypes `gorm:"foreignKey:ItemTypeID;references:id"`
	Companies   Companies `gorm:"foreignKey:CompanyID;references:id"`
	Sizes Sizes `gorm:"foreignKey:SizeID;references:id"`
}
