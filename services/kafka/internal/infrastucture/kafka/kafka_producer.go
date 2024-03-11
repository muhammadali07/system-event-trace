package kafka

import (
    "github.com/IBM/sarama"
)

// KafkaProducer represents the Kafka producer
type KafkaProducer struct {
    producer sarama.SyncProducer
}

// NewKafkaProducer creates a new Kafka producer
func NewKafkaProducer(brokers []string) (*KafkaProducer, error) {
    config := sarama.NewConfig()
    config.Producer.RequiredAcks = sarama.WaitForLocal       // Only wait for the leader to acknowledge the message
    config.Producer.Compression = sarama.CompressionSnappy   // Compress messages
    config.Producer.Flush.Frequency = 500                    // Flush batches every 500ms

    producer, err := sarama.NewSyncProducer(brokers, config)
    if err != nil {
        return nil, err
    }
    return &KafkaProducer{producer: producer}, nil
}

// ProduceMessage produces a message to Kafka
func (p *KafkaProducer) ProduceMessage(topic string, message []byte) error {
    // Create a new message to send to Kafka
    msg := &sarama.ProducerMessage{
        Topic: topic,
        Value: sarama.ByteEncoder(message),
    }

    // Send the message to Kafka
    _, _, err := p.producer.SendMessage(msg)
    return err
}

// Close closes the Kafka producer
func (p *KafkaProducer) Close() error {
    return p.producer.Close()
}
