// main.go
package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/segmentio/kafka-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Account struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	dsn := "host=localhost user=devkafkauser password=devkafkapassword dbname=devkafka port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Account{})

	// Create account handler
	app.Post("/accounts", func(c *fiber.Ctx) error {
		var account Account
		if err := c.BodyParser(&account); err != nil {
			return err
		}

		// Save account to database
		if err := db.Create(&account).Error; err != nil {
			return err
		}

		// Send message to Kafka after creating the account
		err := sendMessageToKafka(account)
		if err != nil {
			return err
		}

		return c.JSON(account)
	})

	app.Listen(":3000")
}

func sendMessageToKafka(account Account) error {
	// Set up Kafka writer
	topic := "account_topic"
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
	defer w.Close()

	// Create message payload
	payload, err := json.Marshal(account)
	if err != nil {
		return err
	}

	// Write message to Kafka
	err = w.WriteMessages(context.Background(), kafka.Message{
		Value: payload,
	})
	if err != nil {
		return err
	}

	log.Println("Message sent to Kafka:", string(payload))
	return nil
}
