package models

import (
	"gorm.io/gorm"
)

type Item_Types struct {
	gorm.Model
	Name string `gorm:not null`
}
