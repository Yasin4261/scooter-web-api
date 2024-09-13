package routes

import (
	"scoter-web-api/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRentalRoutes(app *fiber.App) {
	rental := app.Group("/rentals")

	// Kiralama başlatma endpointi
	rental.Post("/", controllers.StartRental)

	// Kullanıcının kiralamalarını getirme endpointi
	rental.Get("/user/:user_id", controllers.GetUserRentals)
}
