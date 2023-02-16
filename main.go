package main

import (
	"API/routes"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	log.Printf("Service Started")
	app := fiber.New()

	//routes
	routes.VMRoute(app) //add this

	err := app.Listen(":4000")
	if err != nil {
		log.Printf("Cannot listen port 4000. Error: %v", err.Error())
		return
	}

}
