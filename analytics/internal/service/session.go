package service

import (
	"analytics/internal/storage"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Session interface {
	Create(ch <-chan SessionModel)
	Metrics(ctx context.Context)
}

type DroneSession struct {
	storagePg  *pgxpool.Pool
	storageMet storage.Storage
	tick       *time.Ticker
}

func NewSession(storagePg *pgxpool.Pool, storageMet storage.Storage) Session {
	return &DroneSession{
		storagePg:  storagePg,
		storageMet: storageMet,
		tick:       time.NewTicker(30 * time.Second),
	}
}

func (ds *DroneSession) Create(ch <-chan SessionModel) {
	for data := range ch {
		fmt.Printf("Creating session: %s\n", data.SessionID)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		sql := `INSERT INTO session (
				session_id, 
				timestamp, 
				latitude, 
				longitude, 
				altitude, 
				speed, 
				roll, 
				battery
			) VALUES (
				$1, $2, $3, $4, $5, $6, $7, $8
			)`

		_, _ = ds.storagePg.Exec(ctx, sql,
			data.SessionID,
			data.Timestamp,
			data.Latitude,
			data.Longitude,
			data.Altitude,
			data.Speed,
			data.Roll,
			data.Battery,
		)
	}
}

func (ds *DroneSession) Metrics(ctx context.Context) {
	for {
		ctxStorage, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		model := storage.Metrics{}

		err := ds.storagePg.QueryRow(ctxStorage,
			"SELECT * FROM get_flight_statistics()",
		).Scan(
			&model.TotalFlights,
			&model.AvgDistanceMeters,
			&model.MaxDistanceMeters,
			&model.MaxFlightDurationSec,
			&model.FlightsLast30Sec,
			&model.MaxSpeedMps,
			&model.AvgBatteryDrainPercent,
			&model.TotalDistanceMeters,
		)
		if err != nil {
			fmt.Printf("cannot get metrics: %s", err)
		}

		fmt.Println(model)

		ds.storageMet.Write(model)

		select {
		case <-ctx.Done():
			return
		case <-ds.tick.C:
		}
	}
}
