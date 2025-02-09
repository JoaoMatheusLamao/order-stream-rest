package stream

import (
	"fmt"
	"orderstreamrest/internal/repositories/mongo"

	"github.com/segmentio/kafka-go"
)

// KafkaInternal is a struct that contains a Kafka producer
type KafkaInternal struct {
	writer *kafka.Writer
	reader *kafka.Reader
	mgo    *mongo.MongoInternal
}

// NewKafkaInternal is a function that returns a new KafkaInternal struct
func NewKafkaInternal(mgo *mongo.MongoInternal) (*KafkaInternal, error) {
	broker := "kafka:9092"
	topic := "orders"
	groupID := "order-consumer-group"

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
		Brokers:      []string{broker},
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		Async:        true,
		RequiredAcks: int(kafka.RequireOne),
	})

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{broker},
		Topic:       topic,
		Partition:   0,
		MinBytes:    10e3,
		MaxBytes:    10e6,
		StartOffset: kafka.FirstOffset,
		GroupID:     groupID,
	})

	return &KafkaInternal{
		writer: w,
		reader: r,
		mgo:    mgo,
	}, nil
}
