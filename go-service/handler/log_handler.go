package handler

import (
	"github.com/sudo-Brinthan/cyberlog-analyser/go-service/config"
	"github.com/sudo-Brinthan/cyberlog-analyser/go-service/service"
)

func HandleLog(log string, cfg *config.Config) {
	service.ProcessLog(log, cfg)
}