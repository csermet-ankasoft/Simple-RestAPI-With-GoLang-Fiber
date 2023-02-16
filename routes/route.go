package routes

import (
	controller "API/controller"

	"github.com/gofiber/fiber/v2"
)

func VMRoute(app *fiber.App) {
	app.Get("/", controller.Root)
	app.Get("/external", controller.Getapi)
	app.Post("/internal", controller.Postapi)
}
