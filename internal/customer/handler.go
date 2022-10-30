package customer

import (
	"encoding/json"
	"fmt"
	"go-demo-unit-test/domain/entity"

	"github.com/gofiber/fiber/v2"
)

type (
	Handle interface {
		Get(c *fiber.Ctx) error
		Create(c *fiber.Ctx) error
	}
	handle struct {
		service Service
	}
)

func NewHandler(service Service) handle {
	return handle{service}
}

func (h handle) Get(c *fiber.Ctx) error {
	res, _ := h.service.Find()
	return c.JSON(res)
}

func (h handle) Create(c *fiber.Ctx) error {
	var req entity.Customer
	json.Unmarshal(c.Body(), &req)

	fmt.Printf("Create \n")
	res, _ := h.service.Create(req)

	fmt.Printf("end create %+v \n", res)

	return c.JSON(res)
}
