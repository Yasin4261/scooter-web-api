package routes

import (
	"scoter-web-api/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
	// Kullanıcı rotaları
	app.Post("/users/register", controllers.RegisterUser)
	app.Post("/users/login", controllers.LoginUser)
}
