package main

import (
	"log"

	"github.com/sudo-Brinthan/cyberlog-analyser/go-service/config"
	"github.com/sudo-Brinthan/cyberlog-analyser/go-service/kafka"
)

func main() {
	cfg := config.LoadConfig()

	log.Println("🚀 Go Log Ingestion Service Started")

	kafka.StartConsumer(cfg)
}