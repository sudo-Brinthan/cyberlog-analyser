package service

import (
	"log"

	"github.com/sudo-Brinthan/cyberlog-analyser/go-service/client"
	"github.com/sudo-Brinthan/cyberlog-analyser/go-service/config"
)

func ProcessLog(logMsg string, cfg *config.Config) {
	log.Println("📥 Received Log:", logMsg)

	// 1. Send to Python AI for Analysis
	result, err := client.SendToPython(logMsg, cfg.PythonAPI)
	if err != nil {
		log.Println("❌ Python API error:", err)
		return
	}

	log.Println("🧠 AI Response Received")

	// 2. Save the result to Elasticsearch
	client.SaveLog(result)
}
