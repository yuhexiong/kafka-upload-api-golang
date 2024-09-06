package service

import "fmt"

type KafkaUploadService struct{}

func (s *KafkaUploadService) Upload(topic string, data map[string]interface{}) interface{} {
	fmt.Println("topic", topic)
	fmt.Println("data", data)
	return "File uploaded successfully"
}
