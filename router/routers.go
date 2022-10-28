package router

import (
	"go-demo-unit-test/internal/customer"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(app *fiber.App) {

	routerV1 := app.Group("api/v1")

	customer.Router(routerV1)
	// upload.Router(routerV1)

}
