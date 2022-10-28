package main

import (
	routers "go-demo-unit-test/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:   "Go Demo unit test",
		Immutable: true,
	})
	app.Use(recover.New())
	app.Use(cors.New())

	routers.PublicRoutes(app)
	app.Listen(":3000")
}
