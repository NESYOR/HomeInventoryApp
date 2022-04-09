package models

import "gorm.io/gorm"

type Addresses struct {
	gorm.Model
	StreetAddress1 string `gorm:"not null"`
	StreetAddress2 string
	City           string `gorm:"not null"`
	StateID        uint   `gorm:"not null"`
	CountryID      uint   `gorm:"not null"`
	zipcode        string `gorm:"not null"`
	longitude      string
	latitude       string
	States States `gorm:"foreignKey:StateID;references:id"`
	Countries     Countries `gorm:"foreignKey:CountryID;references:id"`
}
