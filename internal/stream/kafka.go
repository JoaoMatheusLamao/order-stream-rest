package stream

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

// KafkaInternal is a struct that contains a Kafka producer
type KafkaInternal struct {
	writer *kafka.Writer
}

// NewKafkaInternal is a function that returns a new KafkaInternal struct
func NewKafkaInternal() (*KafkaInternal, error) {

	broker := "localhost:9092"
	topic := "orders"

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{broker},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	return &KafkaInternal{
		writer: w,
	}, nil
}

// WriteMessage is a function that writes a message to a Kafka topic
func (k *KafkaInternal) WriteMessage(message []byte) error {
	log.Println(fmt.Sprintf("Writing message to Kafka: %s", string(message)))

	msg := kafka.Message{
		Partition: 0,
		Value:     message,
		Time:      time.Now(),
		Key:       []byte(fmt.Sprintf("%d", time.Now().Unix())),
	}

	return k.writer.WriteMessages(context.Background(), msg)
}
