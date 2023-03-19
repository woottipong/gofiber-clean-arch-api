package router

import (
	"golang-clean-arch-api/internal/handler"
	"golang-clean-arch-api/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(route fiber.Router, u usecase.UserUseCase) {
	h := handler.NewUserHandler(u)

	route.Post("/user", h.CreateUser)
	route.Get("/user/:id", h.GetUserByID)
	route.Put("/user/:id", h.UpdateUser)
	route.Delete("/user/:id", h.DeleteUser)
}
