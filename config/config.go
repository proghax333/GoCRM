package config

import (
	"GoCRM/models"
	"GoCRM/utils"

	"github.com/joho/godotenv"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConfigureEnv() error {
	return godotenv.Load(".env")
}

func ConfigureDatabase(connection_uri string, database_name string) error {
	err := mgm.SetDefaultConfig(nil, database_name, options.Client().ApplyURI("mongodb://root:12345@localhost:27017"))
	return err
}

func ConfigureCoreServices() error {
	// Configure environment
	if err := ConfigureEnv(); err != nil {
		return err
	}

	// Configure database
	connection_uri := utils.GetEnv("DATABASE_URI")
	database_name := utils.GetEnv("DATABASE_NAME")
	if err := ConfigureDatabase(connection_uri, database_name); err != nil {
		return err
	}

	// Initialize all models
	models.Users = mgm.Coll(&models.User{})
	models.Customers = mgm.Coll(&models.Customer{})

	return nil
}
