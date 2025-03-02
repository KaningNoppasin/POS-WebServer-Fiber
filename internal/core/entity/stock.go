package entity

import "gorm.io/gorm"

type Stock struct {
	gorm.Model
	ProductID uint `gorm:"uniqueIndex;not null"`
	Quantity  uint `json:"quantity"`
}
