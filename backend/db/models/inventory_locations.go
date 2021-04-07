package models

import (
	"gorm.io/gorm"
)

type Inventory_Locations struct {
	gorm.Model
	Name        string `gorm:not null`
	Description string `gorm:not null`
	Image_Url   string
}
