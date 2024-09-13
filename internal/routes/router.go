package routes

import (
	"scoter-web-api/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/api/scooters", controllers.GetAllScooters)
	app.Post("/api/scooters", controllers.CreateScooter)
	app.Put("/api/scooters/:id/status", controllers.UpdateScooterStatus)
	app.Put("/api/scooters/:id/location", controllers.UpdateScooterLocation)
}
