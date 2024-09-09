package main

import (
	"kafka-upload-api-golang/internal/route"
	"kafka-upload-api-golang/pkg/config"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	config.KafkaBroker = os.Getenv("KAFKA_BROKER")
	if config.KafkaBroker == "" {
		slog.Error("should provide KAFKA_BROKER")
	}

	r := route.SetupRouter()
	r.Run(":8080")
}
