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

	err := dbConn.AutoMigrate(&entity.Product{})
	if err != nil {
		log.Fatalf("Error during migration: %v", err)
	}

	productRepo := db.NewProductRepository(dbConn)
	productService := service.NewProductService(productRepo)
	http.SetUpRoutes(app, productService)
	app.Listen(":8080")
}
