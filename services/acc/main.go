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

func main() {
	cfg, err := utils.InitConfig()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"config": cfg,
		}).Warn(err.Error())
	}

	API_ADDRESS := fmt.Sprintf("%v:%v", cfg.AppHost, cfg.AppPort)
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

	logrus.Info(fmt.Sprintf(" 📢 Server :%v started successfully 🚀 running on : %v", cfg.Service, API_ADDRESS))
	server.Listen(API_ADDRESS)
}
