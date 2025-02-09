package stream

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

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
