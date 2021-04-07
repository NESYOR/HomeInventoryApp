package models

import (
	"gorm.io/gorm"
)

type States struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Code        string `gorm:"not null"`
	CountryID uint   `gorm:not null`
	Countries     Countries `gorm:"foreignKey:CountryID;references:id"`

}
