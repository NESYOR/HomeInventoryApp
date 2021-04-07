package models

import (
	"gorm.io/gorm"
)

type Countries struct {
	gorm.Model
	Name string `gorm:"not null"`
	Code string `gorm:"not null"`
}
