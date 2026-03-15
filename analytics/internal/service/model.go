package service

import "time"

type SessionModel struct {
	SessionID string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Altitude  float64   `json:"altitude"`
	Speed     float64   `json:"speed"`
	Roll      float64   `json:"roll"`
	Battery   float64   `json:"battery"`
}
