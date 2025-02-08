package product

import (
	"shortify/internal/config"

	"go.mongodb.org/mongo-driver/mongo"
)

// delete is a function that deletes a product
func delete(sku string, cfg *config.Config) error {

	// Delete the product from the database
	deleted, err := cfg.Mongo.DeleteProduct(sku)
	if err != nil {
		return err
	}

	if !deleted {
		return mongo.ErrNoDocuments
	}

	return nil
}
