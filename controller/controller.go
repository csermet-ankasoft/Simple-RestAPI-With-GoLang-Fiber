package controller

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

var degisken = ""

func Getapi(c *fiber.Ctx) error {
	var m = make(map[string]string)
	m["TEST"] = string(degisken)
	return c.Status(http.StatusOK).JSON(m)
}

func Postapi(c *fiber.Ctx) error {
	degisken = string(c.Body())
	return c.Status(http.StatusOK).JSON(c.Body())
}
