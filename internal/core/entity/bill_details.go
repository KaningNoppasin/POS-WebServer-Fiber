package entity

import "gorm.io/gorm"

type Bill_Details struct {
	gorm.Model
	BillID    uint    `gorm:"index"`
	ProductID uint    `gorm:"index"`
	Quantity  uint    `json:"quantity"`
	Total     uint    `json:"total"`
	Bill      Bill    `gorm:"foreignKey:BillID;references:ID;constraint:OnDelete:CASCADE;"`
	Product   Product `gorm:"foreignKey:ProductID;references:ID;constraint:OnDelete:CASCADE;"`
}
