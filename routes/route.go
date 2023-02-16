package routes

import (
	controller "API/controller"

	"github.com/gofiber/fiber/v2"
)

func VMRoute(app *fiber.App) {
	app.Get("/api", controller.Getapi)
	app.Post("/api", controller.Postapi)
}
