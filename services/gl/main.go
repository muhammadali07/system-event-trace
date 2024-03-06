package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"github.com/muhammadali07/service-grap-go-api/services/gl/api"
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

	API_HOST := viper.GetString("API_HOST")
	API_PORT := viper.GetInt("API_PORT")
	API_ADDRESS := fmt.Sprintf("%s:%d", API_HOST, API_PORT)
	DB_HOST := viper.GetString("DB_HOST")
	DB_PORT := viper.GetInt("DB_PORT")
	DB_USERNAME := viper.GetString("DB_USERNAME")
	DB_PASSWORD := viper.GetString("DB_PASSWORD")
	DB_DATABASE := viper.GetString("DB_DATABASE")

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
