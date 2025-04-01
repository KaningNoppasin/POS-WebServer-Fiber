package entity

import "gorm.io/gorm"

type Bill_Details struct {
	gorm.Model
	BillID    uint    `gorm:"index"`
	ProductID uint    `json:"product_id" gorm:"index"`
	Quantity  uint    `json:"quantity"`
	Total     uint    `json:"total"`
	Product   Product `gorm:"foreignKey:ProductID;"`
}
