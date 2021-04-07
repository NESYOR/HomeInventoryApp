package models

import (
	"gorm.io/gorm"
)

type ItemImages struct {
	gorm.Model
	ItemID   uint `gorm:"not null"`
	ImageURL string
	Items Items `gorm:"foreignKey:ItemID;references:id"`
}
