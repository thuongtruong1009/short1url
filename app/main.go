package main

import (
	"fmt"
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	"github.com/thuongtruong1009/short1url/routes"
)

func setupRoutes(app *fiber.App) {

	app.Get("/", routes.HomePage)
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api", routes.ShortenURL)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file", err)
	}
	app := fiber.New(fiber.Config{
		Views: html.New("./templates", ".html"),
	})

	app.Use(logger.New())

	setupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}