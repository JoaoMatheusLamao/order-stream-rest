package mongo

import (
	"context"
	"orderstreamrest/internal/models"
)

// InsertOrder is a function that inserts an order into the database
func (m *MongoInternal) InsertOrder(order models.Order) error {

	ctx := context.Background()
	// Insert order into database
	_, err := m.client.Database("stream_orders").Collection("orders").InsertOne(ctx, order)
	if err != nil {
		return err
	}
	return nil
}
