package routes

import (
	"scoter-web-api/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupScooterRoutes(app *fiber.App) {
	// Scooter rotaları
	app.Get("/scooters", controllers.GetAllScooters)
	app.Post("/scooters", controllers.CreateScooter)
	app.Put("/scooters/:id/location", controllers.UpdateScooterLocation)
	app.Put("/scooters/:id/status", controllers.UpdateScooterStatus)
}
