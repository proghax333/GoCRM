package customers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllCustomers(c *fiber.Ctx) error {
	result := fiber.Map{
		"message": "Get all customers",
		"data":    []interface{}{},
	}

	return c.JSON(result)
}

func GetCustomer(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		id = 0
	}

	result := fiber.Map{
		"message": "Get customer",
		"data":    id,
	}

	return c.JSON(result)
}

func CreateCustomer(c *fiber.Ctx) error {
	result := fiber.Map{
		"message": "Create customer",
		"data":    nil,
	}

	return c.JSON(result)
}

func UpdateCustomer(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		id = 0
	}

	result := fiber.Map{
		"message": "Update customer",
		"data":    id,
	}

	return c.JSON(result)
}

func DeleteCustomer(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		id = 0
	}

	result := fiber.Map{
		"message": "Delete customer",
		"data":    id,
	}

	return c.JSON(result)
}
