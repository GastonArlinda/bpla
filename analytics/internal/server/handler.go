package server

import (
	"fmt"
	"net/http"
)

func (s *Server) MetricsHandler(w http.ResponseWriter, r *http.Request) {
	metrics := s.storage.Read()

	w.Header().Set("Content-Type", "text/plain; version=0.0.4; charset=utf-8")

	fmt.Fprintf(w, "# HELP drone_total_flights Total number of flights\n")
	fmt.Fprintf(w, "# TYPE drone_total_flights counter\n")
	fmt.Fprintf(w, "drone_total_flights %d\n", metrics.TotalFlights)

	fmt.Fprintf(w, "# HELP drone_avg_distance_meters Average flight distance\n")
	fmt.Fprintf(w, "# TYPE drone_avg_distance_meters gauge\n")
	fmt.Fprintf(w, "drone_avg_distance_meters %.2f\n", metrics.AvgDistanceMeters)

	fmt.Fprintf(w, "# HELP drone_max_distance_meters Max flight distance\n")
	fmt.Fprintf(w, "# TYPE drone_max_distance_meters gauge\n")
	fmt.Fprintf(w, "drone_max_distance_meters %.2f\n", metrics.MaxDistanceMeters)

	fmt.Fprintf(w, "# HELP drone_max_flight_duration_seconds Max flight duration\n")
	fmt.Fprintf(w, "# TYPE drone_max_flight_duration_seconds gauge\n")
	fmt.Fprintf(w, "drone_max_flight_duration_seconds %d\n", metrics.MaxFlightDurationSec)

	fmt.Fprintf(w, "# HELP drone_flights_last_30sec Number of flights in last 30 seconds\n")
	fmt.Fprintf(w, "# TYPE drone_flights_last_30sec gauge\n")
	fmt.Fprintf(w, "drone_flights_last_30sec %d\n", metrics.FlightsLast30Sec)

	fmt.Fprintf(w, "# HELP drone_max_speed_mps Max speed\n")
	fmt.Fprintf(w, "# TYPE drone_max_speed_mps gauge\n")
	fmt.Fprintf(w, "drone_max_speed_mps %.2f\n", metrics.MaxSpeedMps)

	fmt.Fprintf(w, "# HELP drone_avg_battery_drain_percent Average battery drain\n")
	fmt.Fprintf(w, "# TYPE drone_avg_battery_drain_percent gauge\n")
	fmt.Fprintf(w, "drone_avg_battery_drain_percent %.2f\n", metrics.AvgBatteryDrainPercent)

	fmt.Fprintf(w, "# HELP drone_total_distance_meters Total distance of all flights\n")
	fmt.Fprintf(w, "# TYPE drone_total_distance_meters counter\n")
	fmt.Fprintf(w, "drone_total_distance_meters %.2f\n", metrics.TotalDistanceMeters)
}
