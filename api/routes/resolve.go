package routes

import(
	"github.com/vgnshiyer/tinyurl/database"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func ResolveURL(c *fiber.Ctx) error {
	url := c.Params("url")

	r := database.CreateClient(0)
	defer r.Close()

	value, err := r.Get(database.Ctx, url).Result()
	if err == redis.nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":"tinyurl not found in the database"
		})
	} else if err != nil {
		return c.Status(fiber.StatusInternalError).JSON(fiber.Map {
			"error": "DB connection error"
		})
	}

	rInr := database.CreateClient(1)
	defer rInr.Close()

	_ = rInr.Incr(database.Ctx, "counter")

	return c.Redirect(value, 301)
}