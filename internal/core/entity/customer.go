package entity

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	CustomerName string `json:"customer_name"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	Bill         []Bill `gorm:"constraint:OnDelete:CASCADE;"`
}
