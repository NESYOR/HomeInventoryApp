package models

import (
	"gorm.io/gorm"
)

type Related_Items struct {
	gorm.Model
	ItemID uint `gorm:"not null"`
}
