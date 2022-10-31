package product

import (
	"go-demo-unit-test/pkg/db"

	"github.com/gofiber/fiber/v2"
)

func Router(r fiber.Router) {
	db, _ := db.GetDB()
	repository := NewRepository(db)
	service := NewService(repository)
	handler := NewHandler(service)

	groupRoute := r.Group("/product")
	groupRoute.Get("", handler.Get)
	groupRoute.Post("", handler.Create)
	groupRoute.Post("", handler.Update)
	groupRoute.Post("/:id", handler.Delete)
}
