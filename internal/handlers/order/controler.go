package order

import (
	"orderstreamrest/internal/config"
	"orderstreamrest/internal/models"

	"github.com/gin-gonic/gin"
)

// CreateOrder is a function that creates an order
func CreateOrder(cfg *config.Config) gin.HandlerFunc {

	var request models.Order

	return func(c *gin.Context) {

		err := c.BindJSON(&request)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid request",
			})
			return
		}

		response := produce(request, cfg)
		c.JSON(response.StatusCode, response)

	}
}
