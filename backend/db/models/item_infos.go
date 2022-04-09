package models

import (
	"time"

	"gorm.io/gorm"
)

type ItemInfos struct {
	gorm.Model
	PurchaseDate        time.Time
	ExpirationDate      time.Time
	LastUsed            time.Time
	PurchasePrice       float64
	MSRP                float64
	UserID              uint `gorm:"not null"`
	ItemID              uint `gorm:"not null"`
	InventoryLocationID uint `gorm:"not null"`
	CompanyID           uint `gorm:"not null"`
	Users               Users `gorm:"foreignKey:UserID;references:id"`
	Items               Items `gorm:"foreignKey:ItemID;references:id"`
	InventoryLocations  InventoryLocations `gorm:"foreignKey:InventoryLocationID;references:id"`
	Companies           Companies `gorm:"foreignKey:CompanyID;references:id"`
}
