package routes

import (
	"orderstreamrest/internal/config"
	"orderstreamrest/internal/handlers/healthcheck"
	"orderstreamrest/internal/handlers/order"
	"orderstreamrest/internal/handlers/product"

	"github.com/gin-gonic/gin"
	// "google.golang.org/protobuf/internal/order"
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

	orderGroup := engine.Group("/order")
	orderGroup.POST("", order.SendToKafka(cfg))

}
