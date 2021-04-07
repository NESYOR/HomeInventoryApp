package models

import (
	"gorm.io/gorm"
)

type Shapes struct {
	gorm.Model
	Name string `gorm:not null`
}
