#include "KafkaProducer.h"
#include <iostream>

KafkaProducer::KafkaProducer(const std::string& brokers, const std::string& topic_name)
    : topic_name_(topic_name)
{
    std::string errstr;

    RdKafka::Conf* conf = RdKafka::Conf::create(RdKafka::Conf::CONF_GLOBAL);
    if (!conf) {
        std::cerr << "Не удалось создать Conf\n";
        return;
    }

    conf->set("bootstrap.servers", brokers, errstr);
    if (!errstr.empty()) std::cerr << "bootstrap.servers error: " << errstr << "\n";

    conf->set("acks", "1", errstr);  // или "all" если нужна максимальная надёжность

    producer_ = RdKafka::Producer::create(conf, errstr);
    if (!producer_) {
        std::cerr << "Ошибка создания Producer: " << errstr << "\n";
        delete conf;
        return;
    }

    delete conf;

    // Создаём Topic handle
    topic_ = RdKafka::Topic::create(producer_, topic_name_.c_str(), nullptr, errstr);
    if (!topic_) {
        std::cerr << "Ошибка создания Topic: " << errstr << "\n";
    }
}

KafkaProducer::~KafkaProducer() {
    if (topic_) {
        delete topic_;
        topic_ = nullptr;
    }
    if (producer_) {
        flush(15000);
        delete producer_;
        producer_ = nullptr;
    }
}

bool KafkaProducer::send(const std::string& key, const std::string& value) {
    if (!producer_ || !topic_) {
        std::cerr << "Producer или Topic не инициализированы\n";
        return false;
    }

    RdKafka::ErrorCode resp = producer_->produce(
        topic_,                                 // ← RdKafka::Topic*
        RdKafka::Topic::PARTITION_UA,           // авто-выбор партиции
        RdKafka::Producer::RK_MSG_COPY,
        const_cast<char*>(value.c_str()),       // payload
        value.size(),
        key.empty() ? nullptr : key.c_str(),    // key
        key.size(),
        nullptr                                 // msg_opaque / delivery callback
    );

    if (resp != RdKafka::ERR_NO_ERROR) {
        std::cerr << "produce failed: " << RdKafka::err2str(resp) << "\n";
        return false;
    }

    return true;
}

void KafkaProducer::poll(int timeout_ms) {
    if (producer_) producer_->poll(timeout_ms);
}

void KafkaProducer::flush(int timeout_ms) {
    if (producer_) {
        RdKafka::ErrorCode err = producer_->flush(timeout_ms);
        if (err != RdKafka::ERR_NO_ERROR) {
            std::cerr << "flush failed: " << RdKafka::err2str(err) << "\n";
        }
    }
}