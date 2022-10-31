package product

import (
	"encoding/json"
	"go-demo-unit-test/domain/entity"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type (
	Handler interface {
		Get(c *fiber.Ctx) error
		Create(c *fiber.Ctx) error
		Update(c *fiber.Ctx) error
		Delete(c *fiber.Ctx) error
	}
	handler struct {
		service Service
	}
)

func NewHandler(service Service) handler {
	return handler{service}
}

func (h handler) Get(c *fiber.Ctx) error {
	res, _ := h.service.Find()
	return c.JSON(res)
}

func (h handler) Create(c *fiber.Ctx) error {
	var req entity.Product
	json.Unmarshal(c.Body(), &req)
	res, _ := h.service.Create(req)

	return c.JSON(res)
}

func (h handler) Update(c *fiber.Ctx) error {
	var req entity.Product
	json.Unmarshal(c.Body(), &req)
	res, _ := h.service.Update(req)

	return c.JSON(res)
}

func (h handler) Delete(c *fiber.Ctx) error {
	strId := c.Params("id")
	id, _ := strconv.Atoi(strId)

	res, _ := h.service.Delete(uint(id))

	return c.JSON(res)
}
