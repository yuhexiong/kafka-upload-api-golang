package kafka

import (
	"kafka-upload-api-golang/pkg/config"

	kafkago "github.com/segmentio/kafka-go"
)

func NewKafkaWriter() *kafkago.Writer {
	return &kafkago.Writer{
		Addr:         kafkago.TCP(config.KafkaBroker),
		Balancer:     &kafkago.LeastBytes{},
		RequiredAcks: kafkago.RequireOne,
	}
}
