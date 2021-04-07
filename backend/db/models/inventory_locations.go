package models

import (
	"gorm.io/gorm"
)

type InventoryLocations struct {
	gorm.Model
	Name        string `gorm:not null`
	Description string `gorm:not null`
	Image_Url   string
}
