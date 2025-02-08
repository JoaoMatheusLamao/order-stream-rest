package product

import (
	"orderstreamrest/internal/config"
	"orderstreamrest/internal/models"
)

// create is a function that creates a product
func create(prodIn models.Product, cfg *config.Config) models.GenericResponse {

	response := models.GenericResponse{
		Status:     models.Success,
		Message:    "Product created",
		StatusCode: 201,
		Data:       nil,
	}

	err := prodIn.Validate()
	if err != nil {
		response.Status = models.Error
		response.Message = "Invalid product: " + err.Error()
		response.StatusCode = 400
		return response
	}

	existingProduct, err := cfg.Mongo.GetProduct(1, 1, prodIn.SKU)
	if err != nil {
		response.Status = models.Error
		response.Message = "Internal server error: " + err.Error()
		response.StatusCode = 500
		return response
	}

	if existingProduct != nil && len(existingProduct) > 0 {
		response.Status = models.Error
		response.Message = "Product already exists"
		response.StatusCode = 409
		return response
	}

	err = cfg.Mongo.InsertProduct(prodIn)
	if err != nil {
		response.Status = models.Error
		response.Message = "Internal server error: " + err.Error()
		response.StatusCode = 500
		response.Data = nil
		return response
	}

	return response
}
