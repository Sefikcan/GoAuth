package routes

import (
	"goauth/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/v1/auth/register", controllers.Register)
	app.Post("/api/v1/auth/login", controllers.Login)
	app.Get("/api/v1/auth", controllers.GetUser)
	app.Post("/api/v1/auth/logout", controllers.Logout)
}
