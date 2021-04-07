package models

import (
	"gorm.io/gorm"
)

type States struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Code        string `gorm:"not null"`
	CountriesID uint   `gorm:not null`
}
