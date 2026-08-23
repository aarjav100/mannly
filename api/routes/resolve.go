package routes

import (
	"url-shortener/database"

	"github.com/gofiber/fiber/v3"
	"github.com/redis/go-redis/v9"
)

func ResolveURL(c fiber.Ctx) error {
	url := c.Params("url")

	// Create a Database Client for local db 0
	rdb := database.NewClient(0)

	defer rdb.Close()

	val, err := rdb.Get(database.Ctx, url).Result()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	} else if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "URL not found",
		})
	}

	// Create a Database Client
	rInr := database.NewClient(1)
	defer rInr.Close()

	// Increment the number of times the URL has been resolved
	_ = rInr.Incr(database.Ctx, "counter")
	c.Redirect().Status(fiber.StatusMovedPermanently).To(val)
	return nil
}
