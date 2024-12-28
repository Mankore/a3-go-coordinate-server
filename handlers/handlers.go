package handlers

import (
	"a3-go-coordinate-server/models"
	"a3-go-coordinate-server/parser"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func HelloWorldHandler(c fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

// /coords
func CoordinatesHandler(c fiber.Ctx) error {
	x, err1 := strconv.Atoi(c.Params("x"))
	y, err2 := strconv.Atoi(c.Params("y"))

	if err1 != nil || err2 != nil {
		return c.SendString("Invalid params")
	}

	z := parser.GetZ(models.Coordinate{X: x, Y: y}, parser.CoordinatesMap)

	return c.JSON(fiber.Map{
		"x": x,
		"y": y,
		"z": z,
	})
}
