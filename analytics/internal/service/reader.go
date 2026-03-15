package service

import (
	"context"
	"errors"
	"fmt"
	"math/rand/v2"
	"time"
)

type Fetcher interface{
	Fetch(ctx context.Context, ch chan<- SessionModel)
}

type KafkaFetcher struct{}

func NewFetcher() Fetcher {
	return &KafkaFetcher{}
}

func (kf *KafkaFetcher) Fetch(ctx context.Context, ch chan<- SessionModel) {
	for {
		ctxKafka, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		model, err := generateMockData(ctxKafka)
		if err != nil {
			fmt.Println("err")
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
	time.Sleep(2*time.Second)

	select{
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
