package service

import "time"

type SessionModel struct {
	SessionID string
	Timestamp time.Time
	Latitude  float64
	Longitude float64
	Altitude  float64
	Speed     float64
	Roll      float64
	Battery   float64
}
