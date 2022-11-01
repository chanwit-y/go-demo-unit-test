package user_profile

import (
	"go-demo-unit-test/pkg/util"

	"github.com/gofiber/fiber/v2"
)

func Router(r fiber.Router) {
	httpClient := util.NewHttpClient()
	service := NewService(httpClient)
	handler := NewHandler(service)

	groupRoute := r.Group("/user-profile")
	groupRoute.Get("", handler.GetByEmail)
}
