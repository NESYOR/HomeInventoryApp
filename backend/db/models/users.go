package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Email      string `gorm:"not null;unique"`
	Name       string `gorm:not null`
	Password   string `gorm:not null`
	Last_Login time.Time
}
