package routes

import (
	"GO_PROJECT/controller"
	"GO_PROJECT/service"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App, userService *service.UserService) {
	userController := controller.NewUserController(userService)
	api := app.Group("/api")
	{
		api.Post("/users", userController.CreateUser)
		api.Get("/users/:id", userController.GetUser)
		api.Get("/users", userController.GetAllUser)
		api.Put("/users/:id", userController.UpdateUser)
		api.Delete("/users/:id", userController.DeleteUser)
	}
}
