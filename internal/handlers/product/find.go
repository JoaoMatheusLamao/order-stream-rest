package product

import (
	"shortify/internal/config"
	"shortify/internal/models"
)

// find is a function that finds a product
func find(sku string, cfg *config.Config) (models.Product, error) {

	// Find the product in the database
	product, err := cfg.Mongo.GetProduct(sku)
	if err != nil {
		return product, err
	}

	return product, nil
}
