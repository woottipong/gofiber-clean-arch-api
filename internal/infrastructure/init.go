package infrastructure

import (
	"fmt"
	"golang-clean-arch-api/config"
	"golang-clean-arch-api/internal/handler"
	"golang-clean-arch-api/internal/repository"
	"golang-clean-arch-api/internal/usecase"
	"log"
)

func Initial(cfg *config.Config) {
	// Initialize database connection
	db, err := NewMySQLDatabase(cfg)
	if err != nil {
		log.Fatalf("failed connect to database: %v", err)
	}
	defer db.Close()

	// Initialize fiber app
	app, err := NewFiberApp()
	if err != nil {
		log.Fatalf("failed initializing app: %v", err)
	}

	// define Group /v1
	v1 := app.Group("/v1")

	// Create dependencies
	userRepository := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepository)

	// Initialize HTTP handlers and router
	handler.NewUserHandler(userUseCase, v1)

	// Start server
	err = app.Listen(fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port))
	if err != nil {
		log.Fatalf("Error starting server: %s", err.Error())
	}
}
