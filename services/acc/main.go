// main.go
package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
	telemetry "github.com/muhammadali07/system-event-trace/build/opentelemetry"
	"github.com/muhammadali07/system-event-trace/services/acc/api"
	"github.com/muhammadali07/system-event-trace/services/acc/pkg/utils"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	tracer            trace.Tracer
	propagator        propagation.TextMapPropagator
	statusOptions     = []int{http.StatusOK, http.StatusBadRequest}
	telemetryEndpoint = "localhost:4318"
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

	if telemetryEndpoint == "" {
		telemetryEndpoint = "http://localhost:4318"
	}

	ctx := context.Background()
	propagator = telemetry.NewTelemetryPropagators()
	tp := telemetry.NewHTTPTelemetryProvider(telemetryEndpoint, "account-service", ctx)
	tracer = tp.Tracer("main")

	server := fiber.New()
	logger := logrus.New()
	validator := validator.New()

	server.Use(otelfiber.Middleware(
		otelfiber.WithTracerProvider(tp),
		otelfiber.WithPropagators(propagator),
	))

	api.InitServer(server, db, logger, validator, tracer, true)

	logrus.Info(fmt.Sprintf(" 📢 Service: %v started successfully 🚀 running on -> %v", cfg.Service, API_ADDRESS))
	server.Listen(API_ADDRESS)
}
