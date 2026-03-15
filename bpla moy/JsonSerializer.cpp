#include "JsonSerializer.h"
#include <iomanip>
#include <sstream>

std::string iso8601(const std::chrono::system_clock::time_point& tp) {
    auto tt = std::chrono::system_clock::to_time_t(tp);
    std::tm tm = *std::gmtime(&tt);
    auto ms = std::chrono::duration_cast<std::chrono::milliseconds>(tp.time_since_epoch()) % 1000;

    std::ostringstream oss;
    oss << std::put_time(&tm, "%FT%T") << '.'
        << std::setfill('0') << std::setw(3) << ms.count() << 'Z';
    return oss.str();
}

std::string toJson(const std::string& session_uuid, const TelemetryPoint& p) {
    std::ostringstream oss;
    oss << "{"
        << "\"id\":\""         << session_uuid << "\","               // кавычки вокруг строки
        << "\"timestamp\":\""  << iso8601(p.timestamp) << "\","       // кавычки вокруг строки
        << "\"latitude\":"     << std::fixed << std::setprecision(6) << p.latitude << ","
        << "\"longitude\":"    << std::fixed << std::setprecision(6) << p.longitude << ","
        << "\"altitude\":"     << std::setprecision(2) << p.altitude << ","   // исправлена опечатка
        << "\"speed\":"        << p.speed << ","
        << "\"roll\":"         << p.roll << ","
        << "\"battery\":"      << p.battery
        << "}";
    return oss.str();
}