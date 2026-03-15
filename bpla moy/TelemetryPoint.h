#ifndef TELEMETRY_POINT_H
#define TELEMETRY_POINT_H

#include <chrono>
#include <string>

struct TelemetryPoint {
    std::chrono::system_clock::time_point timestamp;
    double latitude = 0.0;
    double longitude = 0.0;
    double altitute = 0.0;
    double speed = 0.0;
    double roll = 0.0;
    double battery = 0.0;

    std::string toString() const;
};

#endif