package auth

import (
	"github.com/gofiber/fiber/v2"
)

func AuthV1Router() *fiber.App {
	router := fiber.New()

	router.
		Post("/login", Login).
		Post("/register", Register)

	return router
}
