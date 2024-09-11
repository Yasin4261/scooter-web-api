package main

import (
	"log"
	"scoter-web-api/internal/config"
	"scoter-web-api/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Veritabanı bağlantısını kur
	config.ConnectDB()

	routes.SetupRoutes(app) // Fiber router'ı ayarlıyoruz

	log.Println("Server is running on port 8080...")
	log.Fatal(app.Listen(":8080")) // Sunucuyu başlatıyoruz
}
