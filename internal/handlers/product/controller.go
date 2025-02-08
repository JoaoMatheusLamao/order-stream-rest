package product

import (
	"shortify/internal/config"
	"shortify/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetProduct is a function that gets a product
func GetProduct(cfg *config.Config) gin.HandlerFunc {

	return func(c *gin.Context) {

		sku := c.Param("sku")

		response := find(cfg, 1, 1, sku)
		c.JSON(response.StatusCode, response)

	}
}

// GetAllProducts is a function that gets all products
func GetAllProducts(cfg *config.Config) gin.HandlerFunc {

	return func(c *gin.Context) {

		pageStr := c.DefaultQuery("page", "1")
		limitStr := c.DefaultQuery("limit", "50")

		page, err := strconv.ParseInt(pageStr, 10, 64)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid page parameter",
			})
			return
		}

		limit, err := strconv.ParseInt(limitStr, 10, 64)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid limit parameter",
			})
			return
		}

		if page <= 0 {
			page = 1
		}
		response := find(cfg, page, limit)
		c.JSON(response.StatusCode, response)

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

		response := create(request, cfg)
		c.JSON(response.StatusCode, response)
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

		response := update(sku, request, cfg)
		c.JSON(response.StatusCode, response)
	}
}

// DeleteProduct is a function that deletes a product
func DeleteProduct(cfg *config.Config) gin.HandlerFunc {

	return func(c *gin.Context) {

		sku := c.Param("sku")

		response := delete(sku, cfg)
		c.JSON(response.StatusCode, response)

	}
}
