package books

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kasamens/go-fiber-api/pkg/common/models"
)

func (h handler) DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	h.DB.Delete(&book)

	return c.SendStatus(fiber.StatusOK)
}
