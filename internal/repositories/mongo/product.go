package mongo

import (
	"context"
	"log"
	"shortify/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetProduct is a function that gets a product with pagination
func (m *MongoInternal) GetProduct(page int64, limit int64, sku ...string) ([]models.Product, error) {

	collection := m.client.Database("stream_orders").Collection("products")

	var filter bson.D
	if len(sku) > 0 {
		filter = bson.D{{Key: "sku", Value: bson.D{{Key: "$in", Value: sku}}}}
	} else {
		filter = bson.D{}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * limit)
	findOptions.SetLimit(limit)

	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		log.Println("Error finding products in MongoDB database: " + err.Error())
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []models.Product
	if err = cursor.All(ctx, &products); err != nil {
		log.Println("Error decoding products from MongoDB cursor: " + err.Error())
		return nil, err
	}

	return products, nil
}

// InsertProduct is a function that upserts a product
func (m *MongoInternal) InsertProduct(prod models.Product) error {

	collection := m.client.Database("stream_orders").Collection("products")

	// Add insertion date
	prod.InsertionDate = time.Now()

	_, err := collection.InsertOne(context.Background(), prod)
	if err != nil {
		log.Println("Error inserting product into MongoDB database: " + err.Error())
	}
	return err
}

// UpdateProduct is a function that updates a product
func (m *MongoInternal) UpdateProduct(sku string, prod models.Product) error {

	collection := m.client.Database("stream_orders").Collection("products")

	filter := bson.D{{Key: "sku", Value: sku}}

	// Get the current product for history
	var currentProd models.Product
	err := collection.FindOne(context.Background(), filter).Decode(&currentProd)
	if err != nil {
		log.Println("Error getting current product from MongoDB database: " + err.Error())
		return err
	}

	// Convert current product to JSON
	currentProdJSON, err := bson.MarshalExtJSON(currentProd, false, false)
	if err != nil {
		log.Println("Error marshalling current product to JSON: " + err.Error())
		return err
	}

	// Create update history entry
	updateHistory := bson.D{
		{Key: "date", Value: time.Now()},
		{Key: "productOld", Value: string(currentProdJSON)},
	}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "name", Value: prod.Name},
			{Key: "stock", Value: prod.Stock},
			{Key: "price", Value: prod.Price},
		}},
		{Key: "$push", Value: bson.D{
			{Key: "updateHistory", Value: updateHistory},
		}},
	}

	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println("Error updating product in MongoDB database: " + err.Error())
	}
	return err
}

// DeleteProduct is a function that deletes a product
func (m *MongoInternal) DeleteProduct(sku string) (bool, error) {

	collection := m.client.Database("stream_orders").Collection("products")

	filter := bson.D{{Key: "sku", Value: sku}}

	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Println("Error deleting product from MongoDB database: " + err.Error())
		return false, err
	}

	if result.DeletedCount == 0 {
		log.Println("No product found with SKU: ", sku)
		return false, nil
	}

	return true, nil
}
