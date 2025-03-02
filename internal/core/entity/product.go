package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductBarcode string `json:"product_barcode" gorm:"unique"`
	ProductName    string `json:"product_name"`
	ImagePath      string `json:"image_path"`
	Price          uint   `json:"price"`
	Stock          Stock  `gorm:"constraint:OnDelete:CASCADE;"`
}
