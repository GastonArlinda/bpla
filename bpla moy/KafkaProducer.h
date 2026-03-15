#ifndef KAFKA_PRODUCER_H
#define KAFKA_PRODUCER_H

#include <librdkafka/rdkafkacpp.h>
#include <string>
#include <memory>

class KafkaProducer {
public:
    KafkaProducer(const std::string& brokers, const std::string& topic);
    ~KafkaProducer();

    bool send(const std::string& key, const std::string& value);

    // Вызывать периодически
    void poll(int timeout_ms = 0);

    // Дождаться отправки всего
    void flush(int timeout_ms = 10000);

private:
    RdKafka::Producer* producer_ = nullptr;
    RdKafka::Topic* topic_ = nullptr;
    std::string topic_name_;
};

#endif