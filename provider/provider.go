package provider

import (
	"rest-api/cmd/server"
	"rest-api/internals/handler"
	"rest-api/internals/repository"
	"rest-api/internals/routes"
	"rest-api/internals/services"

	"gorm.io/gorm"
)

func NewProvider(db *gorm.DB, server server.GinServer) {
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	routes.RegisterUserRoutes(server, userHandler)
}
