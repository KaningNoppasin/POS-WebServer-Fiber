package response

import (
	"github.com/KaningNoppasin/Web-Server-Fiber/internal/core/entity"
	"github.com/gofiber/fiber/v2"
)

type EntityType interface {
	*entity.Product | []entity.Product |
	*entity.Stock | []entity.Stock |
	*entity.Customer | []entity.Customer |
	*entity.Bill | []entity.Bill
}

func SendSuccessResponse[T EntityType](c *fiber.Ctx, data T) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   data,
	})
}

func SendErrorResponse(c *fiber.Ctx, statusCode int, err error) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"status":  "error",
		"message": err.Error(),
	})
}
