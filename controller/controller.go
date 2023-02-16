package controller

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

var variable = ""

func Root(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON("WELCOME TO THE Simple-RestAPI-With-GoLang-Fiber Application")
}

func Getapi(c *fiber.Ctx) error {
	var mapping = make(map[string]string)
	mapping["external"] = string(variable)
	if variable == "" {
		mapping["external"] = "Please Firstly Save Variable With Post Rest API request."
		return c.Status(http.StatusNotFound).JSON(mapping)
	} else {
		mapping["external"] = string(variable)
		return c.Status(http.StatusOK).JSON(mapping)
	}
}

func Postapi(c *fiber.Ctx) error {
	if string(c.Body()) == "" {
		return c.Status(http.StatusBadRequest).JSON("Body cannot empty.")
	} else {
		variable = string(c.Body())
		return c.Status(http.StatusOK).JSON(c.Body())
	}
}
