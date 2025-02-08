package product

import (
	"orderstream/internal/config"
	"orderstream/internal/models"
)

// delete is a function that deletes a product
func delete(sku string, cfg *config.Config) models.GenericResponse {

	response := models.GenericResponse{
		StatusCode: 200,
		Status:     models.Success,
		Message:    "Product deleted successfully",
	}

	// Delete the product from the database
	deleted, err := cfg.Mongo.DeleteProduct(sku)
	if err != nil {
		response.StatusCode = 400
		response.Status = models.Error
		response.Message = err.Error()
		return response
	}

	if !deleted {
		response.StatusCode = 404
		response.Status = models.Error
		response.Message = "Product not found"
		return response
	}

	return response
}
