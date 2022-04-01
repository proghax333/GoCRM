package api

import (
	"GoCRM/api/v1/auth"
	"GoCRM/api/v1/customers"

	"github.com/gofiber/fiber/v2"
)

func V1Router() *fiber.App {
	router := fiber.New()

	router.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "V1 API is live",
		})
	})

	router.Mount("/auth", auth.AuthV1Router())
	router.Mount("/customers", customers.CustomersV1Router())

	return router
}
