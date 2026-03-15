package main

import (
	"analytics/internal/config"
	"analytics/internal/service"
	"analytics/internal/storage"
	"analytics/pkg"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const path = "../.env"

func main() {
	fmt.Println("Starting service")

	cfg := config.MustLoad(path)

	ctxPg, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pg, err := pkg.SetupPostgres(ctxPg, cfg.Storage.URL)
	if err != nil {
		fmt.Println(err)
		return
	}
	met := storage.NewMetricsStorage()
	kafka := pkg.SetupKafka(cfg.Kafka.Brokers, cfg.Kafka.Topic, "storage-telemetry")

	fetcher := service.NewFetcher(kafka)
	session := service.NewSession(pg, met)

	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan service.SessionModel, 1000)

	for range 5 { go fetcher.Fetch(ctx, ch) }
	for range 10 { go session.Create(ch) }
	go session.Metrics(ctx)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop

	pg.Close()
	kafka.Close()
	cancel()
}
