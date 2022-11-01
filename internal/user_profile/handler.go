package user_profile

import "github.com/gofiber/fiber/v2"

type handler struct {
	service Service
}

func NewHandler(service Service) handler {
	return handler{service}
}

func (h *handler) GetByEmail(c *fiber.Ctx) error {
	email := c.Query("email")
	res, err := h.service.GetByEmail(email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	return c.JSON(res)
}
