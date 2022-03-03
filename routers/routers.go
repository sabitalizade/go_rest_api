package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sabitalizade/go-resttAPI/controller"
)

func Setup(app *fiber.App) {
	app.Post("/api/register",controller.Register)
	app.Post("/api/login",controller.Login)
	app.Get("/api/user",controller.User)
	app.Post("/api/logout",controller.Logout)
}