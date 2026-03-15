#ifndef JSON_SERIALIZER_H
#define JSON_SERIALIZER_H

#include "TelemetryPoint.h"
#include <string>

std::string toJson(const std::string& session_uuid, const TelemetryPoint& point);

#endif