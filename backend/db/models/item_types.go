package models

import (
	"gorm.io/gorm"
)

type ItemTypes struct {
	gorm.Model
	Name string `gorm:not null`
}
