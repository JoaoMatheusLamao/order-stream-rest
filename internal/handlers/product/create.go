package product

import (
	"shortify/internal/config"
	"shortify/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
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

	err = cfg.Mongo.InsertProduct(prodIn)
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
