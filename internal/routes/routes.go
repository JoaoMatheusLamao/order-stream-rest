package routes

import (
	"orderstream/internal/config"
	"orderstream/internal/handlers/healthcheck"
	"orderstream/internal/handlers/product"

	"github.com/gin-gonic/gin"
)

// InitiateRoutes is a function that initializes the routes for the application
func InitiateRoutes(engine *gin.Engine, cfg *config.Config) {
	healthGroup := engine.Group("/healthcheck")
	healthGroup.GET("/", healthcheck.HealthCheck)

	productGroup := engine.Group("/product")
	productGroup.GET("", product.GetAllProducts(cfg))
	productGroup.GET("/:sku", product.GetProduct(cfg))
	productGroup.POST("", product.CreateProduct(cfg))
	productGroup.PUT("/:sku", product.UpdateProduct(cfg))
	productGroup.DELETE("/:sku", product.DeleteProduct(cfg))

}
