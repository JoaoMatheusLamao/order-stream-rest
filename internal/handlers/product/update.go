package product

import (
	"shortify/internal/config"
	"shortify/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

// update is a function that updates a product
func update(sku string, request models.Product, cfg *config.Config) models.GenericResponse {

	request.SKU = sku
	response := models.GenericResponse{
		Status:     models.Success,
		Message:    "Product updated",
		StatusCode: 200,
		Data:       nil,
	}

	err := request.Validate()
	if err != nil {
		response.Status = models.Error
		response.Message = "Invalid product: " + err.Error()
		response.StatusCode = 400
		return response
	}

	// Update the product in the database
	err = cfg.Mongo.UpdateProduct(sku, request)
	if err == mongo.ErrNoDocuments {
		response.Status = models.Error
		response.Message = "Product not found"
		response.StatusCode = 404
		response.Data = nil
		return response
	} else if err != nil {
		response.Status = models.Error
		response.Message = "Internal server error: " + err.Error()
		response.StatusCode = 500
		response.Data = nil
		return response
	}

	return response
}
