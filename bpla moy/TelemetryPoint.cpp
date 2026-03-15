#include "TelemetryPoint.h"
#include <iomanip>
#include <sstream>

std::string TelemetryPoint::toString() const {
    auto tt = std::chrono::system_clock::to_time_t(timestamp);
    std::tm tm = *std::gmtime(&tt);
    auto ms = std::chrono::duration_cast<std::chrono::milliseconds>(timestamp.time_since_epoch()) % 1000;

    std::ostringstream oss;
    oss << "ts=" << std::put_time(&tm, "%Y-%m-%dT%H:%M:%S") << '.'
        << std::setfill('0') << std::setw(3) << ms.count() << "Z "
        << "lat=" << std::fixed << std::setprecision(6) << latitude << " "
        << "lon= " << longitude << " "
        << "alt=" << std::setprecision(1) << altitute << " "
        << "spd=" << speed << " "
        << "roll=" << roll << " "
        << "batt=" << battery;
    return oss.str();
}