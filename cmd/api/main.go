package main

import (
	"log"
	"scoter-web-api/internal/config"
	"scoter-web-api/internal/repositories"
	"scoter-web-api/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db := config.ConnectDB()

	repositories.InitRepositories(db) // Repository'yi başlatın.

	app := fiber.New()

	routes.SetupRentalRoutes(app)
	routes.SetupScooterRoutes(app)
	routes.SetupUserRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
