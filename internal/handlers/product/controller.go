package product

import (
	"shortify/internal/config"
	"shortify/internal/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetProduct is a function that gets a product
func GetProduct(cfg *config.Config) gin.HandlerFunc {

	return func(c *gin.Context) {

		sku := c.Param("sku")

		product, err := find(sku, cfg)
		if err == mongo.ErrNoDocuments {
			c.JSON(404, gin.H{
				"error": "Product not found",
			})
			return
		} else if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, product)
	}
}

// CreateProduct is a function that creates a product
func CreateProduct(cfg *config.Config) gin.HandlerFunc {

	var request models.Product

	return func(c *gin.Context) {

		err := c.BindJSON(&request)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = create(request, cfg)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(201, gin.H{
			"message": "Product created successfully",
		})
	}

}

// UpdateProduct is a function that updates a product
func UpdateProduct(cfg *config.Config) gin.HandlerFunc {

	var request models.Product

	return func(c *gin.Context) {

		sku := c.Param("sku")

		err := c.BindJSON(&request)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = update(sku, request, cfg)
		if err == mongo.ErrNoDocuments {
			c.JSON(404, gin.H{
				"error": "Product not found",
			})
			return
		} else if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "Product updated successfully",
		})
	}
}

// DeleteProduct is a function that deletes a product
func DeleteProduct(cfg *config.Config) gin.HandlerFunc {

	return func(c *gin.Context) {

		sku := c.Param("sku")

		err := delete(sku, cfg)
		if err == mongo.ErrNoDocuments {
			c.JSON(404, gin.H{
				"error": "Product not found",
			})
			return
		} else if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "Product deleted successfully",
		})
	}
}
