package eventstream

import (
	"log"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	writer *kafka.Writer
}

func NewKafkaProducer(brokers []string, topic string) *KafkaProducer {
	return &KafkaProducer{
		writer: &kafka.Writer{
			Addr:  kafka.TCP(brokers...),
			Topic: topic,
		},
	}
}

func (kp *KafkaProducer) ProduceMessage(key, value []byte) error {
	return kp.writer.WriteMessages(kafka.Message{
		Key:   key,
		Value: value,
	})
}

func (kp *KafkaProducer) Close() {
	err := kp.writer.Close()
	if err != nil {
		log.Printf("Error closing Kafka producer: %v", err)
	}
}
