package handlers

import (
	"a3-go-coordinate-server/models"
	"a3-go-coordinate-server/parser"
	"fmt"
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
	mapName := c.Params("map")
	fmt.Println("Map:", mapName, "x:", x, "y:", y)

	if err1 != nil || err2 != nil || mapName == "" {
		return c.SendString("Invalid params")
	}

	if _, exists := parser.MapsMap[mapName]; !exists {
		fmt.Println("Map not found in memory, attempt reading from file:", mapName)
		parser.ReadCoordinatesFromFile(mapName)
	}

	z := parser.GetZ(models.Coordinate{X: x, Y: y}, parser.MapsMap[mapName])

	return c.JSON(fiber.Map{
		"x": x,
		"y": y,
		"z": z,
	})
}
