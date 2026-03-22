package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/sudo-Brinthan/cyberlog-analyser/go-service/config"
	"github.com/sudo-Brinthan/cyberlog-analyser/go-service/handler"
)

func StartConsumer(cfg *config.Config) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{cfg.KafkaBroker},
		Topic:   cfg.Topic,
		GroupID: "cyberlog-group",
	})

	log.Println("📡 Kafka Consumer Started...")

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("❌ Error reading message:", err)
			continue
		}

		handler.HandleLog(string(msg.Value), cfg)
	}
}