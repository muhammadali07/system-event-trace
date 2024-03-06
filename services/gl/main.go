package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/muhammadali07/service-grap-go-api/services/gl/api"
	"github.com/muhammadali07/service-grap-go-api/services/gl/pkg/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	cfg := utils.Config{}

	API_HOST := cfg.AppHost
	API_PORT := cfg.AppPort
	API_ADDRESS := fmt.Sprintf("%s:%d", API_HOST, API_PORT)
	DB_HOST := cfg.DatabaseHost
	DB_PORT := cfg.DatabasePort
	DB_USERNAME := cfg.DatabaseUser
	DB_PASSWORD := cfg.DatabasePassword
	DB_DATABASE := cfg.Database

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta", DB_HOST, DB_PORT, DB_USERNAME, DB_PASSWORD, DB_DATABASE)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		panic(err)
	}
	server := fiber.New()
	logger := logrus.New()
	validator := validator.New()

	api.InitServer(server, db, logger, validator)

	server.Listen(API_ADDRESS)
}
