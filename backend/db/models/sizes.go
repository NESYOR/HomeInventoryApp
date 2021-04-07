package models

import (
	"gorm.io/gorm"
)

type Sizes struct {
	gorm.Model
	Name    string  `gorm:"not null"`
	Length  float64 `gorm:"not null"`
	Width   float64 `gorm:"not null"`
	Height  float64 `gorm:"not null"`
	ShapeID uint    `gorm:"not null"`
	Volume  float64
}

//Countries Countries `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
