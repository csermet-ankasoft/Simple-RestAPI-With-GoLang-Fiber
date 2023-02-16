package main

import (
	"API/routes"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	log.Printf("Created By CNR")
	log.Printf("Fiber Service Started")
	app := fiber.New()

	//routes
	routes.VMRoute(app) //add this

	err := app.Listen(":200")
	if err != nil {
		log.Printf("Cannot listen port 200. Error: %v", err.Error())
		return
	}

}
