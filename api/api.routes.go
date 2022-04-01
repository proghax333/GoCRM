package api

import (
	"GoCRM/api/v1"

	"github.com/gofiber/fiber/v2"
)

func APIRouter() *fiber.App {
	router := fiber.New()

	router.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "API is live",
		})
	})

	router.Mount("/v1", api.V1Router())

	return router
}
