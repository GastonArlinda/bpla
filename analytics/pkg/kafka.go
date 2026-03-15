package pkg

import (
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader
}

type MessageHandler func([]byte) error

func SetupKafka(brokers, topic, groupID string) (*kafka.Reader) {
	fmt.Println([]string{brokers})

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"kafka:9092"},
		Topic:       topic,
		GroupID:     groupID,
		MinBytes:    10e3,
		MaxBytes:    10e6,
		MaxWait:     1 * time.Second,
		StartOffset: kafka.FirstOffset,
	})

	return reader
}
