package product

import (
	"orderstreamrest/internal/config"
	"orderstreamrest/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

// find is a function that finds a product
func find(cfg *config.Config, page, limit int64, sku ...string) models.GenericResponse {

	response := models.GenericResponse{
		Status:     models.Success,
		Message:    "Product found",
		StatusCode: 200,
		Data:       nil,
	}

	// Find the product in the database
	product, err := cfg.Mongo.GetProduct(page, limit, sku...)
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

	if len(product) == 0 {
		response.Status = models.Error
		response.Message = "Product not found"
		response.StatusCode = 404
		response.Data = nil
		return response
	} else {
		response.Data = product
	}

	return response
}
