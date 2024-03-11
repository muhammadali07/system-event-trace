package usecase

import (
    "context"
    "encoding/json"
    "github.com/muhammadali07/service-grap-go-api/services/kafka/internal/app/entity"
    "github.com/IBM/sarama"
)

type KafkaConsumerUsecase struct {
    // Define any dependencies here
}

func (uc *KafkaConsumerUsecase) ProcessMessage(message *sarama.ConsumerMessage) error {
    // Decode the message data into entity.Message
    var msg entity.Message
    if err := json.Unmarshal(message.Value, &msg); err != nil {
        return err
    }

    // Your code to process the message
    // Example: Log the received message
    // Replace it with your actual business logic
    log.Printf("Received message: %s\n", string(message.Value))

    return nil
}

func (uc *KafkaConsumerUsecase) StartConsuming(ctx context.Context, consumer sarama.Consumer, topic string) error {
    // Start consuming messages from Kafka
    partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
    if err != nil {
        return err
    }
    defer partitionConsumer.Close()

    for {
        select {
        case <-ctx.Done():
            // Context canceled, stop consuming
            return nil
        case err := <-partitionConsumer.Errors():
            // Handle errors
            return err
        case msg := <-partitionConsumer.Messages():
            // Process the received message
            if err := uc.ProcessMessage(msg); err != nil {
                // Handle processing error
                return err
            }
        }
    }
}
