package utils

import "log"

func Info(msg string) {
	log.Println("ℹ️", msg)
}

func Error(msg string) {
	log.Println("❌", msg)
}