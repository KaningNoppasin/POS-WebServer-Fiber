package seeds

import (
	"github.com/KaningNoppasin/Web-Server-Fiber/internal/core/entity"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	db.Create(&entity.Product{
		ProductBarcode: "001",
		ProductName:    "Product - 001",
		ImagePath:      "uploads/default_image.jpg",
		Price:          20,
		Stock: entity.Stock{
			Quantity: 10,
		},
	})
	db.Create(&entity.Product{
		ProductBarcode: "002",
		ProductName:    "Product - 002",
		ImagePath:      "uploads/default_image.jpg",
		Price:          30,
		Stock: entity.Stock{
			Quantity: 20,
		},
	})
	db.Create(&entity.Customer{
		CustomerName: "NN",
		Phone:        "0800000000",
		Email:        "nn@test.com",
	})
	db.Create(&entity.Customer{
		CustomerName: "John",
		Phone:        "0800000000",
		Email:        "john@test.com",
	})
	db.Create(&entity.Bill{
		CustomerID:  1,
		TotalAmount: 20,
	})
	db.Create(&entity.Bill{
		CustomerID:  2,
		TotalAmount: 40,
	})
	db.Create(&entity.Bill{
		CustomerID:  2,
		TotalAmount: 100,
	})
	db.Create(&entity.Bill_Details{
		BillID:    1,
		ProductID: 1,
		Quantity:  3,
		Total:     10,
	})
	db.Create(&entity.Bill_Details{
		BillID:    1,
		ProductID: 2,
		Quantity:  1,
		Total:     20,
	})
	db.Create(&entity.Bill_Details{
		BillID:    2,
		ProductID: 1,
		Quantity:  5,
		Total:     10,
	})
	db.Create(&entity.Bill_Details{
		BillID:    3,
		ProductID: 2,
		Quantity:  5,
		Total:     50,
	})
}
