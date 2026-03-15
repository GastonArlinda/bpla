package storage

import "time"

type Metrics struct {
    TotalFlights           int       
    AvgDistanceMeters      float64   
    MaxDistanceMeters      float64   
    MaxFlightDurationSec   int       
    FlightsLast30Sec       int       
    MaxSpeedMps            float64   
    AvgBatteryDrainPercent float64   
    TotalDistanceMeters    float64   
    CalculatedAt           time.Time 
}