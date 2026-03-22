package config

import "os"

type Config struct {
	KafkaBroker string
	Topic       string
	PythonAPI   string
}

func LoadConfig() *Config {
	return &Config{
		KafkaBroker: getEnv("KAFKA_BROKER", "localhost:9092"),
		Topic:       getEnv("KAFKA_TOPIC", "logs"),
		PythonAPI:   getEnv("PYTHON_API", "http://localhost:8000/analyze"),
	}
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}