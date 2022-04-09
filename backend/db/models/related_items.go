package models

import (
	"gorm.io/gorm"
)

type RelatedItems struct {
	gorm.Model
	ItemID uint `gorm:"not null"`
	Items Items `gorm:"foreignKey:ItemID;references:id"`
}
