package entity

import "gorm.io/gorm"

type Bill struct {
	gorm.Model
	CustomerID   uint           `json:"customer_id"`
	TotalAmount  uint           `json:"total_amount"`
	Bill_Details []Bill_Details `gorm:"constraint:OnDelete:CASCADE;"`
}
