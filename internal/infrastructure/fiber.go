package infrastructure

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func NewFiberApp() (*fiber.App, error) {
	// Creates a new Fiber instance.
	app := fiber.New(fiber.Config{
		AppName: "Register API version 1.0.0",
	})

	// Use global middleware.
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	return app, nil
}
