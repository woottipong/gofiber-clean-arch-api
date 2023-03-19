package handler

import (
	"github.com/gofiber/fiber/v2"
)

type BookHandler struct {
	// bookUseCase usecase.BookUseCase
}

// route fiber.Router,
func NewBookHandler(route fiber.Router) {
	handler := &BookHandler{
		// bookUseCase: bookUseCase,
	}

	// Declare routing endpoints
	user := route.Group("/books")
	user.Get("/", handler.GetAll)
}

func (h *BookHandler) GetAll(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": "this is a book.",
	})
}
