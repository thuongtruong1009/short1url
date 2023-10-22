package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/thuongtruong1009/short1url/routes"
)

func setupRoutes(app *fiber.App) {
	app.Post("/", routes.ShortenURL)
	app.Post("", routes.ShortenURL)
	app.Get("/:url", routes.ResolveURL)
	app.Post("/all", routes.AllShortedURLsOfUser)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file", err)
	}
	
	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${latency} ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Asia/Ho_Chi_Minh",
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		// AllowMethods: "GET,POST",
		// AllowOrigins: "http://localhost:3000, http://localhost:3001, https://short1url.vercel.app",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
		}, ","),
    	AllowHeaders:  "Origin, Content-Type, Accept",
	}))

	// proxy
	// app.Use(corsMiddleware)

	setupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}

// func corsMiddleware(c *fiber.Ctx) error {
// 	c.Response().Header.Set("Access-Control-Allow-Origin", "*")
// 	c.Response().Header.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
// 	c.Response().Header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

// 	if c.Method() == "OPTIONS" {
// 		return c.SendStatus(http.StatusOK)
// 	}

// 	return c.Next()
// }
