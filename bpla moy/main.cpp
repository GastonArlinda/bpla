#include <iostream>
#include <chrono>
#include <thread>
// #include <fstream>
#include "TelemetryGenerator.h"
#include "JsonSerializer.h"
#include "KafkaProducer.h"

std::string make_session_uuid() {
    return "sess-" + std::to_string(std::chrono::steady_clock::now().time_since_epoch().count());
}

int main() {
    std::string brokers = "localhost:9092";
    std::string topic = "bpla-telemetry";

    KafkaProducer kafka(brokers, topic);
    if (!kafka.send("", "")) {
        std::cerr << "Kafka не инициализировалась, выходим\n";
        return 1;
    }

    TelemetryGenerator gen;
    std::string session = make_session_uuid();

    std::cout << "Запуск генерации и отправки в Kafka\n";
    std::cout << "Сессия: " << session << "\n";
    std::cout << "Брокер: " << brokers << " | Топик: " << topic << "\n\n";

    const int POINTS = 120;
    const int INTERVAL_MS = 1000;

    for (int i = 0; i < POINTS; ++i) {
        TelemetryPoint p = gen.nextPoint();
        std::string json = toJson(session, p);

        bool ok = kafka.send(session, json);

        if (ok) {
            std::cout << "[" << i+1 << "/" << POINTS << "] Отправлено: "
                      << json.substr(0, 80) << "...\n";
        } else {
            std::cout << "Ошибка отправки!\n";
        }

        kafka.poll(0);

        // std::this_thread::sleep_for(std::chrono::milliseconds(INTERVAL_MS));
    }

    kafka.flush(15000);
    std::cout << "\nЗавершено. Все сообщения отправлены (или в очереди). \n";
    
    return 0;
}