package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sabitalizade/go-resttAPI/database"
	"github.com/sabitalizade/go-resttAPI/routers"
)

func main() {
    database.Connect()
    app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
    routers.Setup(app)
 app.Listen(":4000")
}