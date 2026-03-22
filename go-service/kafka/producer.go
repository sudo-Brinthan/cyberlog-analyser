package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func ProduceLog(broker, topic, message string) {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker},
		Topic:   topic,
	})

	err := writer.WriteMessages(context.Background(),
		kafka.Message{Value: []byte(message)},
	)

	if err != nil {
		log.Println("❌ Error producing message:", err)
	} else {
		log.Println("✅ Message sent to Kafka")
	}
}