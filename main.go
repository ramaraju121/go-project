package main

import (
	dbconfiguration "GO_PROJECT/dbConfiguration"
	"GO_PROJECT/repositories"
	"GO_PROJECT/routes"
	"GO_PROJECT/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	db := dbconfiguration.ConnectDB()
	userRepo := repositories.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	routes.SetUpRoutes(app, userService)
	app.Listen(":8080")
}
