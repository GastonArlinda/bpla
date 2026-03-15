#ifndef TELEMETRY_GENERATOR_H
#define TELEMETRY_GENERATOR_H

#include "TelemetryPoint.h"
#include <vector>
#include <random>

class TelemetryGenerator {
private:
    std::mt19937 rng;
    double current_altitude = 180.0;
    double current_batery = 100.0;

    const double BASE_LAT = 55.831;
    const double BASE_LON = 37.553;

public:
    TelemetryGenerator();

    TelemetryPoint nextPoint();

    std::vector<TelemetryPoint> generateBatch(size_t count);
};

#endif