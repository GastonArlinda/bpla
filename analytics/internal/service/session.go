package service

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Session interface{
	Create(ch <-chan SessionModel)
}

type DroneSession struct {
	storage *pgxpool.Pool
}

func NewSession(storage *pgxpool.Pool) Session {
	return &DroneSession{
		storage: storage,
	}
}

func (ds *DroneSession) Create(ch <-chan SessionModel) {
	for data := range ch {
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

		_, _ = ds.storage.Exec(ctx, sql,
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
