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

	conn, err := kafka.Dial("tcp", broker)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Kafka broker: %w", err)
	}
	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		return nil, fmt.Errorf("failed to read partitions: %w", err)
	}

	topicExists := false
	for _, p := range partitions {
		if p.Topic == topic {
			topicExists = true
			break
		}
	}

	if !topicExists {
		err = conn.CreateTopics(kafka.TopicConfig{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 1,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create topic: %w", err)
		}
	}

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
