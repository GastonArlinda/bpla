package config

import (
	"log"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env     string `env:"LEVEL"`
	Srv     Server
	Storage Storage
	Kafka   Kafka
}

type Server struct {
	Type        string        `env:"SERVER_TYPE"`
	Addr        string        `env:"SERVER_ADDR"`
	Timeout     time.Duration `env:"SERVER_TIMEOUT_SECONDS"`
	IdleTimeout time.Duration `env:"SERVER_IDLE_TIMEOUT_SECONDS"`
	MaxConn     string        `env:"SERVER_MAX_CONN"`
}

type Storage struct {
	Type string `env:"DB_TYPE"`
	URL  string `env:"DB_URL"`
}

type Kafka struct {
	Brokers string `env:"KAFKA_BROKERS" env-default:"kafka:9092"`
	Topic   string `env:"KAFKA_TOPIC" env-default:"drone-sessions"`
	GroupID string `env:"KAFKA_GROUP_ID" env-default:"analytics-group"`
}

func MustLoad(path string) *Config {
	var cfg Config

	_ = cleanenv.ReadConfig(path, &cfg)

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("failed to read environment variables: %v", err)
	}

	return &cfg
}