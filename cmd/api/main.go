package main

import (
	"log"
	"scoter-web-api/internal/config"
	"scoter-web-api/internal/repositories"
	"scoter-web-api/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.ConnectDB()

	repositories.InitScooterRepository() // Repository'yi başlatın.

	app := fiber.New()

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
