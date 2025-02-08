package product

import (
	"shortify/internal/config"
	"shortify/internal/models"
)

// create is a function that creates a product
func create(prodIn models.Product, cfg *config.Config) error {

	err := prodIn.Validate()
	if err != nil {
		return err
	}

	return cfg.Mongo.InsertProduct(prodIn)
}
