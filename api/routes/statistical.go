package routes

import (
	"fmt"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/thuongtruong1009/short1url/database"
)

func AllShortedURLsOfUser(c *fiber.Ctx) error {
	ip := c.Query("ip")

	r := database.CreateClient(0)
	defer r.Close()

	result, err := r.SMembers(database.Ctx, "index:" + ip).Result()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot retrieve client"})
	}

	var updatedResult []string

	// Iterate over the set members
	for _, member := range result {
		value, err := r.Get(database.Ctx, member).Result()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot retrieve URLs of client"})
		}

		updatedResult = append(updatedResult, os.Getenv("DOMAIN") + "/" + member)

		fmt.Printf("Key: %s, Value: %s\n", member, value)
	}

	return c.Status(fiber.StatusOK).JSON(updatedResult)
}
	