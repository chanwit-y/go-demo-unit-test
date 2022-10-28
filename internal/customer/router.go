package customer

import (
	"go-demo-unit-test/pkg/db"

	"github.com/gofiber/fiber/v2"
)

func Router(r fiber.Router) {
	db, _ := db.GetDB()
	repository := NewRepository(db)
	service := NewService(repository)
	handler := NewHandler(service)

	groupRoute := r.Group("/customer")
	groupRoute.Get("", handler.Get)
	groupRoute.Post("", handler.Create)
}
