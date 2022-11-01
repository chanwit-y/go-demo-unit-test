package router

import (
	"go-demo-unit-test/internal/customer"
	"go-demo-unit-test/internal/product"
	"go-demo-unit-test/internal/user_profile"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(app *fiber.App) {

	routerV1 := app.Group("api/v1")

	customer.Router(routerV1)
	product.Router(routerV1)
	user_profile.Router(routerV1)

}
