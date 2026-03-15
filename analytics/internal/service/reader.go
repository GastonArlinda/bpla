package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand/v2"
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
		ctxKafka, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		msg, err := kf.kafka.ReadMessage(ctxKafka)
		if err != nil {
			fmt.Printf("Error reading message: %v\n", err)
			continue
		}

		model := SessionModel{}

		if err := json.Unmarshal(msg.Value, &model); err != nil {
			fmt.Printf("Error parse json %s\n", model.SessionID)
			continue
		}

		fmt.Printf("Fetching data %s\n", model.SessionID)

		ch <- model

		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}

func generateMockData(ctx context.Context) (SessionModel, error) {
	time.Sleep(2 * time.Second)

	select {
	case <-ctx.Done():
		return SessionModel{}, errors.New("timeout")
	default:
		return SessionModel{
			SessionID: "550e8400-e29b-41d4-a716-446655440000",
			Timestamp: time.Now(),
			Latitude:  55.7558 + rand.Float64()*0.01,
			Longitude: 37.6176 + rand.Float64()*0.01,
			Altitude:  100.0 + rand.Float64()*50,
			Speed:     10.0 + rand.Float64()*15,
			Roll:      rand.Float64()*30 - 15,
			Battery:   75.0 + rand.Float64()*20,
		}, nil
	}
}
