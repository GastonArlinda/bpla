#include"TelemetryGenerator.h"
#include <chrono>


TelemetryGenerator::TelemetryGenerator() : rng(std::random_device{}()) {}

TelemetryPoint TelemetryGenerator::nextPoint() {
    std::uniform_real_distribution<double> d_lat(-0.0008, 0.0008);
    std::uniform_real_distribution<double> d_lon(-0.0008, 0.0008);
    std::uniform_real_distribution<double> d_alt(-4.0, +7.0);
    std::uniform_real_distribution<double> d_speed(8.0, 24.0);
    std::uniform_real_distribution<double> d_roll(-12.0, 12);
    std::uniform_real_distribution<double> d_batt(-0.35, -0.05);

    TelemetryPoint p;
    p.timestamp = std::chrono::system_clock::now();

    p.latitude = BASE_LAT + d_lat(rng);
    p.longitude = BASE_LON + d_lon(rng);

    current_altitude += d_alt(rng);
    if (current_altitude < 40.0) current_altitude = 40.0;
    p.altitude = current_altitude;

    p.speed = d_speed(rng);
    p.roll = d_roll(rng);

    current_batery += d_batt(rng);
    if (current_batery < 0.0) current_batery = 0.0;
    p.battery = current_batery;

    return p;
}

std::vector<TelemetryPoint> TelemetryGenerator::generateBatch(size_t count) {
    std::vector<TelemetryPoint> batch;
    batch.reserve(count);
    for (size_t i = 0; i < count; i++) {
        batch.push_back(nextPoint());
    }
    return batch;
}