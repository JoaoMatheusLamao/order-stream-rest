package order

import (
	"encoding/json"
	"orderstreamrest/internal/config"
	"orderstreamrest/internal/models"
	"orderstreamrest/internal/utils"
)

// produce is a function that produces an order
func produce(order models.Order, cfg *config.Config) models.GenericResponse {

	order.InternalOrderID = utils.GenerateUniqueID()

	orderBy, err := json.Marshal(order)
	if err != nil {
		return models.GenericResponse{
			StatusCode: 500,
			Status:     models.Error,
			Message:    "Error marshalling order",
		}
	}

	err = cfg.Kafka.WriteMessage(orderBy)
	if err != nil {
		return models.GenericResponse{
			StatusCode: 500,
			Status:     models.Error,
			Message:    "Error producing order",
		}
	}

	return models.GenericResponse{
		StatusCode: 200,
		Status:     models.Success,
		Message:    "Order produced",
		Data:       map[string]string{"internal_order_id": order.InternalOrderID},
	}

}
