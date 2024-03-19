package app

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/muhammadali07/service-grap-go-api/services/acc/models"
	"github.com/segmentio/kafka-go"
)

func (a *AccountApp) SendMessageToKafka(params models.Account) (response models.RespSendingKafka, err error) {
	// Set up Kafka writer
	topic := "account_topic"
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
	defer w.Close()

	// Create message payload
	payload, err := json.Marshal(params)
	if err != nil {
		err = fmt.Errorf(err.Error())
		return
	}

	// Write message to Kafka
	err = w.WriteMessages(context.Background(), kafka.Message{
		Value: payload,
	})
	if err != nil {
		err = fmt.Errorf(err.Error())
		return
	}

	log.Println("Message sent to Kafka:", string(payload))
	return
}
