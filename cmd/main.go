package main

import (
	"log"

	"github.com/KaningNoppasin/Web-Server-Fiber/internal/adapter/db"
	"github.com/KaningNoppasin/Web-Server-Fiber/internal/adapter/http"
	"github.com/KaningNoppasin/Web-Server-Fiber/internal/core/entity"
	"github.com/KaningNoppasin/Web-Server-Fiber/internal/core/service"
	"github.com/KaningNoppasin/Web-Server-Fiber/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Static("/uploads", "./uploads")

	dbConn := database.ConnectDB()

	// dbConn.Migrator().DropTable(
	// 	&entity.Product{},
	// 	&entity.Stock{},
	// 	&entity.Customer{},
	// 	&entity.Bill{},
	// 	&entity.Bill_Details{},
	// )

	err := dbConn.AutoMigrate(
		&entity.Product{},
		&entity.Stock{},
		&entity.Customer{},
		&entity.Bill{},
		&entity.Bill_Details{},
	)

	if err != nil {
		log.Fatalf("Error during migration: %v", err)
	}
	// seeds.Seed(dbConn)

	productRepo := db.NewProductRepository(dbConn)
	productService := service.NewProductService(productRepo)

	stockRepo := db.NewStockRepository(dbConn)
	stockService := service.NewStockService(stockRepo)

	customerRepo := db.NewCustomerRepository(dbConn)
	customerService := service.NewCustomerService(customerRepo)

	billRepo := db.NewBillRepository(dbConn)
	billService := service.NewBillService(billRepo)

	http.SetUpRoutes(
		app,
		productService,
		stockService,
		customerService,
		billService,
	)

	app.Listen(":8080")
}
