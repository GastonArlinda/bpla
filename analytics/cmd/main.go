package cmd

import (
	"analytics/internal/config"
	"analytics/internal/service"
	"analytics/pkg"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const path = "../.env"

func main() {
	cfg := config.MustLoad(path)

	ctxPg, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pg, err := pkg.SetupPostgres(ctxPg, cfg.Storage.URL)
	if err != nil {
		return
	}

	fetcher := service.NewFetcher()
	session := service.NewSession(pg)

	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan service.SessionModel, 1000)

	fetcher.Fetch(ctx, ch)
	session.Create(ch)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop

	pg.Close()
	cancel()
}
