// main.go
package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/muhammadali07/service-grap-go-api/services/acc/api"
	"github.com/muhammadali07/service-grap-go-api/services/acc/pkg/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// type Account struct {
// 	ID       uint   `json:"id" gorm:"primaryKey"`
// 	Username string `json:"username"`
// 	Email    string `json:"email"`
// }

func main() {

	cfg := utils.Config{}

	fmt.Println(cfg)

	API_HOST := cfg.AppHost
	API_PORT := "3000"
	API_ADDRESS := fmt.Sprintf("%v:%v", API_HOST, API_PORT)
	// DB_HOST := cfg.DatabaseHost
	// DB_PORT := cfg.DatabasePort
	// DB_USERNAME := cfg.DatabaseUser
	// DB_PASSWORD := cfg.DatabasePassword
	// DB_DATABASE := cfg.Database

	// dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta", DB_HOST, DB_PORT, DB_USERNAME, DB_PASSWORD, DB_DATABASE)
	dsn := "host=localhost user=devkafkauser password=devkafkapassword dbname=devkafka port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		panic(err)
	}
	server := fiber.New()
	logger := logrus.New()
	validator := validator.New()

	api.InitServer(server, db, logger, validator)

	server.Listen(API_ADDRESS)

	// app := fiber.New()
	// app.Use(logger.New())

	// dsn := "host=localhost user=devkafkauser password=devkafkapassword dbname=devkafka port=5432 sslmode=disable"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic("failed to connect database")
	// }

	// db.AutoMigrate(&Account{})

	// // Create account handler
	// app.Post("/accounts", func(c *fiber.Ctx) error {
	// 	var account Account
	// 	if err := c.BodyParser(&account); err != nil {
	// 		return err
	// 	}

	// 	// Save account to database
	// 	if err := db.Create(&account).Error; err != nil {
	// 		return err
	// 	}

	// 	// Send message to Kafka after creating the account
	// 	err := sendMessageToKafka(account)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	return c.JSON(account)
	// })

	// app.Listen(":3000")
}

// func sendMessageToKafka(account Account) error {
// 	// Set up Kafka writer
// 	topic := "account_topic"
// 	w := kafka.NewWriter(kafka.WriterConfig{
// 		Brokers:  []string{"localhost:9092"},
// 		Topic:    topic,
// 		Balancer: &kafka.LeastBytes{},
// 	})
// 	defer w.Close()

// 	// Create message payload
// 	payload, err := json.Marshal(account)
// 	if err != nil {
// 		return err
// 	}

// 	// Write message to Kafka
// 	err = w.WriteMessages(context.Background(), kafka.Message{
// 		Value: payload,
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	log.Println("Message sent to Kafka:", string(payload))
// 	return nil
// }
