package config

import "os"

type Config struct {
	KafkaBroker string
	Topic       string
	PythonAPI   string
	ElasticURL  string
}

func LoadConfig() *Config {
	return &Config{
		KafkaBroker: getEnv("KAFKA_BROKER", "localhost:9092"),
		Topic:       getEnv("KAFKA_TOPIC", "nlp-text-pipeline"),
		PythonAPI:   getEnv("PYTHON_API", "http://localhost:8000/api/v1/analyze"),
		ElasticURL:  getEnv("ELASTIC_URL", "http://localhost:9200"),
	}
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
