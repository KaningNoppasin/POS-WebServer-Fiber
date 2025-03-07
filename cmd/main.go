package main

import (
	"log"

	"github.com/KaningNoppasin/Web-Server-Fiber/internal/adapter/db"
	"github.com/KaningNoppasin/Web-Server-Fiber/internal/adapter/http"
	"github.com/KaningNoppasin/Web-Server-Fiber/internal/core/entity"
	"github.com/KaningNoppasin/Web-Server-Fiber/internal/core/service"
	"github.com/KaningNoppasin/Web-Server-Fiber/pkg/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
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

	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}​\n",
	}))

	productRepo := db.NewProductRepository(dbConn)
	productService := service.NewProductService(productRepo)

	stockRepo := db.NewStockRepository(dbConn)
	stockService := service.NewStockService(stockRepo)

	customerRepo := db.NewCustomerRepository(dbConn)
	customerService := service.NewCustomerService(customerRepo)

	billRepo := db.NewBillRepository(dbConn)
	billService := service.NewBillService(billRepo)

	bill_detailsRepo := db.NewBill_DetailsRepository(dbConn)
	bill_detailsService := service.NewBill_DetailsService(bill_detailsRepo)

	http.SetUpRoutes(
		app,
		productService,
		stockService,
		customerService,
		billService,
		bill_detailsService,
	)

	app.Listen(":8080")
}
