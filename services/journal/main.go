// journal_service.go
package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/segmentio/kafka-go"
)

func main() {
	app := fiber.New()

	config := kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		GroupID: "journal-consumer-group",
		Topic:   "account_topic",
	}

	reader := kafka.NewReader(config)
	defer reader.Close()

	go func() {
		for {
			message, err := reader.ReadMessage(context.Background())
			if err != nil {
				log.Println("Error reading message:", err)
				continue
			}
			log.Printf("Received message: %s\n", message.Value)

			// Handle message received from Kafka, e.g., save to database
		}
	}()

	app.Listen(":4000")
}
