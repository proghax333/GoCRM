package customers

import (
	"github.com/gofiber/fiber/v2"
)

func CustomersV1Router() *fiber.App {
	router := fiber.New()

	router.
		Get("/", GetAllCustomers).
		Get("/:id", GetCustomer).
		Post("/", CreateCustomer).
		Put("/:id", UpdateCustomer).
		Delete("/:id", DeleteCustomer)

	return router
}
