package kafka

import (
    "github.com/IBM/sarama"
)

// KafkaConsumer represents the Kafka consumer
type KafkaConsumer struct {
    consumer sarama.Consumer
}

// NewKafkaConsumer creates a new Kafka consumer
func NewKafkaConsumer(brokers []string) (*KafkaConsumer, error) {
    config := sarama.NewConfig()
    consumer, err := sarama.NewConsumer(brokers, config)
    if err != nil {
        return nil, err
    }
    return &KafkaConsumer{consumer: consumer}, nil
}

// ConsumeMessages starts consuming messages from Kafka
func (c *KafkaConsumer) ConsumeMessages(topic string) (<-chan *sarama.ConsumerMessage, <-chan error, error) {
    partitionConsumer, err := c.consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
    if err != nil {
        return nil, nil, err
    }
    messages := partitionConsumer.Messages()
    errors := partitionConsumer.Errors()
    return messages, errors, nil
}
