package main

import (
	"log"

	"github.com/sudo-Brinthan/cyberlog-analyser/go-service/client"
	"github.com/sudo-Brinthan/cyberlog-analyser/go-service/config"
	"github.com/sudo-Brinthan/cyberlog-analyser/go-service/kafka"
)

func main() {
	cfg := config.LoadConfig()

	client.InitElastic(cfg.ElasticURL)

	log.Println("🚀 Go Log Ingestion Service Started")

	kafka.StartConsumer(cfg)
}
