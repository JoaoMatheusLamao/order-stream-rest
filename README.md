# Order Stream

## Description

Order Stream is a microservice-based application designed to handle product orders. It includes a CRUD for products, an API for creating orders that sends messages to Kafka, and a consumer that reads messages from Kafka and stores them in MongoDB.

## Features

- **Product CRUD**: Create, Read, Update, and Delete products.
- **Order Producer**: API to create orders and send them to Kafka.
- **Order Consumer**: Reads orders from Kafka and stores them in MongoDB.

## Project Structure

- `internal/models`: Contains data structure definitions such as `Order`, `Item`, and `Customer`.
- `internal/repositories/mongo`: Contains logic for interacting with MongoDB.
- `internal/stream`: Contains logic for interacting with Kafka, including the producer and consumer.

## Prerequisites

- Docker
- Docker Compose

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/your-username/orderstreamrest.git
    cd orderstreamrest
    ```

2. Start the services using Docker Compose:

    ```sh
    docker-compose up -d
    ```

3. Access the application at `http://localhost:8080`.

## Usage

### Product CRUD

- **Create Product**: `POST /products`
  - Request Body:

    ```json
    {
      "name": "Product Name",
      "description": "Product Description",
      "price": 100.0
    }
    ```

- **List Products**: `GET /products`

- **Update Product**: `PUT /products/{id}`
  - Request Body:

    ```json
    {
      "name": "Updated Product Name",
      "description": "Updated Product Description",
      "price": 150.0
    }
    ```

- **Delete Product**: `DELETE /products/{id}`

### Order Producer

- **Create Order**: `POST /orders`
  - Request Body:

    ```json
    {
      "order_id": "12345",
      "internal_order_id": "54321",
      "items": [
        {
          "item_id": "item1",
          "quantity": 2,
          "price": 10.5
        },
        {
          "item_id": "item2",
          "quantity": 1,
          "price": 20.0
        }
      ],
      "total": 41.0,
      "customer": {
        "customer_id": "cust123",
        "name": "John Doe",
        "email": "john.doe@example.com"
      }
    }
    ```

### Order Consumer

The consumer is automatically started and reads messages from Kafka, sending them to MongoDB.

## Environment Variables

The following environment variables can be configured:

- `KAFKA_BROKER`: Kafka broker address (default: `kafka:9092`)
- `KAFKA_TOPIC`: Kafka topic for orders (default: `orders`)
- `KAFKA_GROUP_ID`: Kafka consumer group ID (default: `order-consumer-group`)
- `MONGO_URI`: MongoDB connection URI (default: `mongodb://mongo:27017`)
- `MONGO_DB`: MongoDB database name (default: `orderstream`)

## Technologies Used

- Go
- Kafka
- MongoDB
- Docker
- Docker Compose

## Running Tests

To run the tests, use the following command:

```sh
go test ./...
```

## Contributing

1. Fork the project.
2. Create a branch for your feature (`git checkout -b feature/new-feature`).
3. Commit your changes (`git commit -am 'Add new feature'`).
4. Push to the branch (`git push origin feature/new-feature`).
5. Open a Pull Request.
