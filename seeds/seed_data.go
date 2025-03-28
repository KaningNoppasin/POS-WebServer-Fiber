package seeds

import (
	"github.com/KaningNoppasin/Web-Server-Fiber/internal/core/entity"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	db.Create(&entity.Product{
		ProductBarcode: "001",
		ProductName:    "Product - 001",
		ImagePath:      "uploads/default_image.jpeg",
		Price:          20,
		Stock: entity.Stock{
			Quantity: 10,
		},
	})
	db.Create(&entity.Product{
		ProductBarcode: "002",
		ProductName:    "Product - 002",
		ImagePath:      "uploads/default_image.jpeg",
		Price:          30,
		Stock: entity.Stock{
			Quantity: 20,
		},
	})
	db.Create(&entity.Product{
		ProductBarcode: "003",
		ProductName:    "Product - 003",
		ImagePath:      "uploads/Teapot-Multi.webp",
		Price:          40,
		Stock: entity.Stock{
			Quantity: 5,
		},
	})
	db.Create(&entity.Product{
		ProductBarcode: "004",
		ProductName:    "Product - 004",
		ImagePath:      "uploads/Ruby-Red-Wine-Glasses.jpg",
		Price:          50,
		Stock: entity.Stock{
			Quantity: 5,
		},
	})
	// 8C 9D 84 81
	// 91b089b404f42a22c878e405618b9139ab4436bf2dba9b45e8ba2bdb1ef1d9fe
	db.Create(&entity.Customer{
		CustomerName: "NN",
		Phone:        "0812345670",
		Email:        "nn@test.com",
		CardUID:      "91b089b404f42a22c878e405618b9139ab4436bf2dba9b45e8ba2bdb1ef1d9fe",
	})
	// CB 97 40 E3
	// ccb05b3531190ad2b7691826f13d2853ce68ab3abc5df07c6ef4d20b2691035c
	db.Create(&entity.Customer{
		CustomerName: "John",
		Phone:        "0987654321",
		Email:        "john@test.com",
		CardUID:      "ccb05b3531190ad2b7691826f13d2853ce68ab3abc5df07c6ef4d20b2691035c",
	})
	db.Create(&entity.Bill{
		CustomerID:  1,
		TotalAmount: 90,
	})
	db.Create(&entity.Bill{
		CustomerID:  2,
		TotalAmount: 100,
	})
	db.Create(&entity.Bill{
		CustomerID:  2,
		TotalAmount: 150,
	})
	db.Create(&entity.Bill_Details{
		BillID:    1,
		ProductID: 1,
		Quantity:  3,
		Total:     60,
	})
	db.Create(&entity.Bill_Details{
		BillID:    1,
		ProductID: 2,
		Quantity:  1,
		Total:     30,
	})
	db.Create(&entity.Bill_Details{
		BillID:    2,
		ProductID: 1,
		Quantity:  5,
		Total:     100,
	})
	db.Create(&entity.Bill_Details{
		BillID:    3,
		ProductID: 2,
		Quantity:  5,
		Total:     150,
	})
}
