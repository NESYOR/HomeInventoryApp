package models

import (
	"time"

	"gorm.io/gorm"
)

type Item_Infos struct {
	gorm.Model
	Purchase_Date       time.Time
	Expiration_Date     time.Time
	Last_Used           time.Time
	Purchase_Price      float64
	MSRP                float64
	UserID              uint `gorm:"not null"`
	ItemID              uint `gorm:"not null"`
	InventoryLocationID uint `gorm:"not null"`
	CompanyID           uint `gorm:"not null"`
}
