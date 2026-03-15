#include <iostream>
#include <chrono>
#include <thread>
#include <random>
#include <iomanip>
// #include <fstream>
#include "TelemetryGenerator.h"
#include "JsonSerializer.h"
#include "KafkaProducer.h"

std::string generate_uuid() {
    std::random_device rd;
    std::mt19937 gen(rd());
    std::uniform_int_distribution<> dis(0, 15);

    std::stringstream ss;
    ss << std::hex << std::setfill('0');

    // 8 символов
    for (int i = 0; i < 8; ++i) ss << dis(gen);
    ss << "-";

    // 4 символа
    for (int i = 0; i < 4; ++i) ss << dis(gen);
    ss << "-";

    // 4 символа (версия 4: 0100xxxx)
    ss << "4";  // версия 4
    for (int i = 0; i < 3; ++i) ss << dis(gen);
    ss << "-";

    // 4 символа (вариант: 10xx xxxx)
    ss << (dis(gen) % 4 + 8);  // 8,9,a,b
    for (int i = 0; i < 3; ++i) ss << dis(gen);
    ss << "-";

    // 12 символов
    for (int i = 0; i < 12; ++i) ss << dis(gen);

    return ss.str();
}

int main() {
    std::string brokers = "kafka:9092";
    std::string topic = "bpla-telemetry";

    KafkaProducer kafka(brokers, topic);
    if (!kafka.send("", "")) {
        std::cerr << "Kafka не инициализировалась, выходим\n";
        return 1;
    }

    TelemetryGenerator gen;
    std::string session = generate_uuid();

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