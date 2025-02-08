package product

import (
	"shortify/internal/config"
	"shortify/internal/models"
)

// update is a function that updates a product
func update(sku string, request models.Product, cfg *config.Config) error {

	request.SKU = sku

	// Validate the product
	err := request.Validate()
	if err != nil {
		return err
	}

	// Update the product in the database
	err = cfg.Mongo.UpdateProduct(sku, request)
	if err != nil {
		return err
	}

	return nil
}
