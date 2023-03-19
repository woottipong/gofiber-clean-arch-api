package response

import "github.com/gofiber/fiber/v2"

type Response interface {
	Send(ctx *fiber.Ctx, data interface{})
	SendError(ctx *fiber.Ctx, statusCode int, message string)
}

type response struct{}

func NewResponse() Response {
	return &response{}
}

func (*response) Send(ctx *fiber.Ctx, data interface{}) {
	ctx.JSON(data)
}

func (*response) SendError(ctx *fiber.Ctx, statusCode int, message string) {
	ctx.Status(statusCode).JSON(fiber.Map{"error": message})
}
