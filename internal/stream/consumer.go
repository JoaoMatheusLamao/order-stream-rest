package stream

import (
	"context"
	"encoding/json"
	"log"
	"orderstreamrest/internal/models"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/segmentio/kafka-go"
)

// Consumer is a function that reads messages from a Kafka topic
func (k *KafkaInternal) Consumer() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Capturando sinais do SO para shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Criando o channel para mensagens do Kafka
	messageChan := make(chan kafka.Message)
	go toProcessConsumer(ctx, k.reader, messageChan)

	for {
		select {
		case msg := <-messageChan:
			log.Printf("Order received: %s\n", string(msg.Value))
			go sendToMongo(msg.Value, k)

		case sig := <-sigChan:
			log.Printf("\nReceived signal %v, shutting down consumer...\n", sig)
			cancel()
			time.Sleep(3 * time.Second)
			return
		}
	}
}

// toProcessConsumer is a function that reads messages from a Kafka topic
func toProcessConsumer(ctx context.Context, r *kafka.Reader, messageChan chan<- kafka.Message) {
	defer close(messageChan)

	for {
		select {
		case <-ctx.Done():
			log.Println("Shutting down consumer...")
			r.Close()
			return

		default:
			msg, err := r.ReadMessage(ctx)
			if err != nil {
				log.Printf("Error reading message: %v", err)
				continue
			}
			messageChan <- msg
		}
	}
}

// sendToMongo is a function that inserts an order into MongoDB
func sendToMongo(orderIn []byte, k *KafkaInternal) {
	order := models.Order{}
	err := json.Unmarshal(orderIn, &order)
	if err != nil {
		log.Println("Error deserializing order: ", err)
		return
	}

	err = k.mgo.InsertOrder(order)
	if err != nil {
		log.Println("Error inserting order into MongoDB: ", err)
		return
	}

	log.Println("Order successfully inserted into MongoDB!")
}
