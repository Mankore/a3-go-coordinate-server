package main

import (
	"a3-go-coordinate-server/handlers"
	"a3-go-coordinate-server/parser"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.HelloWorldHandler)
	app.Get("/coords/:map/:x.:y", handlers.CoordinatesHandler)
}

func main() {
	defaultMap := "altis"
	parser.ReadCoordinatesFromFile(defaultMap)

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
