package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

type Fetcher interface {
	Fetch(ctx context.Context, ch chan<- SessionModel)
}

type KafkaFetcher struct {
	kafka *kafka.Reader
}

func NewFetcher(kafka *kafka.Reader) (Fetcher) {
	return &KafkaFetcher{
		kafka: kafka,
	}
}

func (kf *KafkaFetcher) Fetch(ctx context.Context, ch chan<- SessionModel) {
	for {
		ctxKafka, cancel := context.WithTimeout(context.Background(), 30*time.Second)

		msg, err := kf.kafka.ReadMessage(ctxKafka)
		if err != nil {
			fmt.Printf("Error reading message: %v\n", err)
			continue
		}
		cancel()

		model := SessionModel{}

		if err := json.Unmarshal(msg.Value, &model); err != nil {
			fmt.Printf("Error parse json\n")
			continue
		}

		fmt.Printf("Fetching data %+v\n", model)

		ch <- model

		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}
