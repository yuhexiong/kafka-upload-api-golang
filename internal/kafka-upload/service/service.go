package service

import (
	"context"
	"encoding/json"

	"kafka-upload-api-golang/pkg/kafka"

	kafkago "github.com/segmentio/kafka-go"
)

type KafkaUploadService struct{}

func (s *KafkaUploadService) Upload(c context.Context, topic string, data map[string]interface{}) error {
	dataByte, err := json.Marshal(data)
	if err != nil {
		return err
	}
	message := kafkago.Message{
		Topic: topic,
		Value: dataByte,
	}

	writer := kafka.NewKafkaWriter()
	if err := writer.WriteMessages(c, message); err != nil {
		return err
	}

	return nil
}
