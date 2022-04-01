package main

import (
	"GoCRM/api"
	"GoCRM/config"
	"GoCRM/utils"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	if err := config.ConfigureCoreServices(); err != nil {
		log.Fatalf("Error occured while configuring services: %v", err)
	} else {
		log.Printf("Started all services successfully!")
	}

	// Initialize web server
	app := fiber.New()

	// Configure logger
	app.Use(logger.New())

	// Initialize all the routes
	app.Mount("/api", api.APIRouter())

	app_port, err := strconv.Atoi(utils.GetEnv("APP_PORT"))
	if err != nil {
		app_port = 5000
	}
	log.Printf("Starting the server on port %v...\n", app_port)

	app.Listen(fmt.Sprintf(":%v", app_port))
}
